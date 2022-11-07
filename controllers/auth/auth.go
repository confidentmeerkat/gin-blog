package auth

import (
	"errors"
	"fmt"
	"gin-blog/models"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var uniqueValidator validator.Func = func(fl validator.FieldLevel) bool {
	field := fl.Field().String()
	match := fl.Param()

	db, _ := models.Database()
	var user models.User

	result := db.Where(fmt.Sprintf("%s = ?", match), field).Find(&user)

	if result.RowsAffected > 0 {
		return false
	}

	return true
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateJWT(username string) (string, error) {
	secretKey := os.Getenv("SECRET")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	mySecret := []byte(secretKey)

	tokenString, err := token.SignedString(mySecret)

	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}
	return tokenString, nil
}

type RegisterInput struct {
	Username string `json:"username" binding:"required,unique=username"`
	Email    string `json:"email" binding:"required,unique=email"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var registerInput RegisterInput

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("unique", uniqueValidator)
	}

	if err := c.ShouldBindJSON(&registerInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, _ := HashPassword(registerInput.Password)

	newUser := models.User{Username: registerInput.Username, Email: registerInput.Email, Password: hashedPassword}

	db, err := models.Database()

	if err != nil {
		log.Fatal(err.Error())
	}

	if err := db.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newUser)

}

func Login(c *gin.Context) {
	var loginInput LoginInput

	if err := c.ShouldBindJSON(&loginInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User

	db, err := models.Database()

	if err != nil {
		log.Fatal(err.Error())
	}

	err = db.Where("username = ?", loginInput.Username).Find(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not a registered user"})
		return
	}

	match := CheckPasswordHash(loginInput.Password, user.Password)

	if !match {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect password"})
		return
	}

	token, _ := GenerateJWT(user.Username)

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Currentuser(c *gin.Context) {
	username := c.GetString("username")

	db, err := models.Database()
	if err != nil {
		log.Fatal(err.Error())
	}

	var user models.User
	err = db.Where("username = ?", username).Find(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNetworkAuthenticationRequired, "Incorrect user")
	}

	c.JSON(http.StatusOK, user)

}
