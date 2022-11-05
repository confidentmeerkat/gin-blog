package auth

import (
	"fmt"
	"gin-blog/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
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

type RegisterInput struct {
	Username string `json:"username" binding:"required,unique=username"`
	Email    string `json:"email" binding:"required,unique=email"`
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
