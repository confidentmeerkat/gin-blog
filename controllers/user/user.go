package user

import (
	"gin-blog/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var users []models.User

	db, err := models.Database()
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := db.Find(&users).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}
