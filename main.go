package main

import (
	"github.com/bmvinicius/go-rest-api/database"
	"github.com/bmvinicius/go-rest-api/server"
)

func main() {
	db := database.ConnectToDB()
	database.RunMigrations(db)
	s := server.NewServer("3000", db)
	s.Run()
}
