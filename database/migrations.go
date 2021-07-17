package database

import (
	"github.com/bmvinicius/go-rest-api/src/books/domain/books"
	"gorm.io/gorm"
	"log"
)

func RunMigrations(db *gorm.DB) {
	err := db.AutoMigrate(&books.Book{})
	if err != nil {
		log.Fatal(err)
	}
}
