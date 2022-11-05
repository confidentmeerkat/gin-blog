package main

import (
	"gin-blog/controllers/auth"
	"gin-blog/controllers/user"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"test": "This is only for test"})
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	router.GET("/test", Test)
	router.GET("/users", user.Index)
	router.POST("/register", auth.Register)
	router.POST("/login", auth.Login)

	router.Run("localhost:8080")
}
