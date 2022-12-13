package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password" gorm:"unique"`
}

type UserPublic struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}
