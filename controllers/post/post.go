package post

import (
	"fmt"
	"gin-blog/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var posts []models.Post

	db, err := models.Database()

	if err != nil {
		log.Fatal(err.Error())
	}

	db.Find(&posts)

	c.JSON(http.StatusOK, posts)
}

type NewPost struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"Content" binding:"required"`
}

func Create(c *gin.Context) {
	username := c.GetString("username")
	fmt.Println("username :", username)
	var newPostInput NewPost

	if err := c.ShouldBindJSON(&newPostInput); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	db, err := models.Database()
	if err != nil {
		log.Fatal(err.Error())
	}

	var user models.User
	db.Where("username = ?", username).Find(&user)
	fmt.Println("user :", user)

	newPost := models.Post{Title: newPostInput.Title, Content: newPostInput.Content, AuthorID: int(user.ID)}

	db.Create(&newPost)

	c.JSON(http.StatusOK, newPost)
}
