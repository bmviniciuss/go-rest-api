package database

import (
	"github.com/bmvinicius/go-rest-api/models"
	"gorm.io/gorm"
	"log"
)

func RunMigrations(db *gorm.DB) {
	err := db.AutoMigrate(&models.Book{})
	if err != nil {
		log.Fatal(err)
	}
}
