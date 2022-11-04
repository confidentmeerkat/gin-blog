package models

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Database() (*gorm.DB, error) {
	dbURL := "postgres://postgres:cm1108@Jupiter@localhost:5432/gin-blog"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err = db.AutoMigrate(&User{}); err != nil {
		log.Println(err)
	}

	if err != nil {
		log.Fatal(err.Error())
	}

	return db, err
}
