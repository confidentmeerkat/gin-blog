package main

import (
	"gin-blog/controllers/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"test": "This is only for test"})
}

func main() {
	router := gin.Default()

	router.GET("/test", Test)
	router.GET("/users", user.Index)

	router.Run("localhost:8080")
}
