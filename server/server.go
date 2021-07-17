package server

import (
	"github.com/bmvinicius/go-rest-api/main/router"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	port   string
	server *gin.Engine
	db     *gorm.DB
}

func NewServer(port string, db *gorm.DB) *Server {
	return &Server{
		port:   port,
		server: gin.Default(),
		db:     db,
	}
}

func (s *Server) Run() {
	router.ConfigRoutes(s.server, s.db)
	log.Fatal(s.server.Run(":" + s.port))
}
