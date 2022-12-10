package models

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Database() (*gorm.DB, error) {
	dbURL := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}

	if err = db.AutoMigrate(&User{}, &Post{}); err != nil {
		log.Println(err)
	}

	if err != nil {
		log.Fatal(err.Error())
	}

	return db, err
}
