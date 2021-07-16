package routes

import (
	"github.com/bmvinicius/go-rest-api/controllers"
	"github.com/bmvinicius/go-rest-api/infra/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ConfigRoutes(router *gin.Engine, db *gorm.DB) {
	main := router.Group("api/v1")
	applyBooksRoutes(main, db)
}

func applyBooksRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	booksRepository := repositories.NewBookRepository(db)
	booksController := controllers.NewBookController(booksRepository)

	books := rg.Group("books")
	books.GET("/", booksController.GetBooks)
	books.GET("/:id", booksController.GetBook)
	books.POST("/", booksController.CreateBook)
}
