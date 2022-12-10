package main

import (
	"gin-blog/controllers/auth"
	"gin-blog/controllers/post"
	"gin-blog/controllers/user"
	"gin-blog/middlewares"
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

	authorized := router.Group("/")

	// router.GET("/test", Test)
	router.GET("/users", user.Index)
	router.POST("/register", auth.Register)
	router.POST("/login", auth.Login)

	authorized.Use(middlewares.AuthenticationMiddleware())
	authorized.GET("/test", Test)
	authorized.GET("/me", auth.Currentuser)
	authorized.POST("/posts", post.Create)

	posts := router.Group("/posts")
	posts.GET("/", post.Index)

	router.Run("localhost:8080")
}
