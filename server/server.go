package server

import (
	"github.com/bmvinicius/go-rest-api/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
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
	routes.ConfigRoutes(s.server, s.db)
	log.Fatal(s.server.Run(":" + s.port))
}
