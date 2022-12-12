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

	docs "gin-blog/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"test": "This is only for test"})
}

// @title		Gin-Blog API
// @version 	1.0
// @description	This is a backend for simple blog site built with Gin-Gonic
// @BasePath 	/api

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	apiRouter := router.Group("/api")

	authorized := apiRouter.Group("/")

	// apiRouter.GET("/test", Test)
	apiRouter.GET("/users", user.Index)
	apiRouter.POST("/register", auth.Register)
	apiRouter.POST("/login", auth.Login)

	authorized.Use(middlewares.AuthenticationMiddleware())
	authorized.GET("/test", Test)
	authorized.GET("/me", auth.Currentuser)
	authorized.POST("/posts", post.Create)

	posts := apiRouter.Group("/posts")
	posts.GET("/", post.Index)

	docs.SwaggerInfo.BasePath = "/api"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run("localhost:8080")
}
