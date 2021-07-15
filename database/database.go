package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func ConnectToDB() *gorm.DB {
	connectionStr := "host=db port=5432 user=postgres password=root sslmode=disable dbname=gobook"

	db, err := gorm.Open(postgres.Open(connectionStr), &gorm.Config{})
	if err != nil {
		log.Fatal("error: ", err)
	}

	return db
}
