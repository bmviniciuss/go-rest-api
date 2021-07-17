package router

import (
	handlers "github.com/bmvinicius/go-rest-api/handlers"
	"github.com/bmvinicius/go-rest-api/infra/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ConfigRoutes(router *gin.Engine, db *gorm.DB) {
	baseRouter := router.Group("api/v1")
	// /books
	booksRepository := repositories.NewBookRepository(db)
	handlers.NewBookHandler(booksRepository).ApplyRoutes(baseRouter)
}
