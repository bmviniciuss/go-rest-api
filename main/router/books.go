package router

import (
	"github.com/bmvinicius/go-rest-api/src/books/domain/books"
	repository "github.com/bmvinicius/go-rest-api/src/books/infra/books"
	"github.com/bmvinicius/go-rest-api/src/books/interfaces/public/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateBooksRouter(rg *gin.RouterGroup, db *gorm.DB) {
	r := repository.NewBookGormRepository(db)
	s := books.NewService(&r)
	h := http.NewBooks(s)
	h.ApplyRoutes(rg)
}
