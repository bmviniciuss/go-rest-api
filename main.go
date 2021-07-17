package main

import (
	database "github.com/bmvinicius/go-rest-api/database"
	"github.com/bmvinicius/go-rest-api/main/server"
)

func main() {
	db := database.ConnectToDB()
	database.RunMigrations(db)
	// TODO: Read from env
	s := server.NewServer("3000", db)
	s.Run()
}
