package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title    string `json:"title" gorm:"required"`
	Content  string `json:"content"`
	AuthorID int
	Author   User
}
