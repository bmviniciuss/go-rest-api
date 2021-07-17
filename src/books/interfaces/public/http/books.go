package http

import (
	"errors"
	"github.com/bmvinicius/go-rest-api/src/books/domain/books"
	"github.com/bmvinicius/go-rest-api/src/common/http"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Books struct {
	s *books.Service
}

func NewBooks(s *books.Service) *Books {
	return &Books{s: s}
}

func (b *Books) ApplyRoutes(rg *gin.RouterGroup) {
	booksRouterGroup := rg.Group("books")
	booksRouterGroup.GET("/", b.getBooks)
	booksRouterGroup.GET("/:id", b.getBook)
	booksRouterGroup.POST("/", b.createBook)
}

func (b *Books) getBook(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		http.BadRequestErrorResponse(c, "Id must be a integer")
		return
	}

	var book *books.Book
	book, err = b.s.GetBookById(id)

	if err != nil {
		if errors.Is(err, books.ErrBookNotFound) {
			http.NotFoundErrorResponse(c, err.Error())
			return
		}
		http.ServerErrorResponse(c, err.Error())
		return
	}

	http.OkResponse(c, book)
}

func (b *Books) getBooks(c *gin.Context) {
	booksList, err := b.s.GetBooks()
	if err != nil {
		http.ServerErrorResponse(c, err.Error())
		return
	}
	http.OkResponse(c, booksList)
}

func (b *Books) createBook(c *gin.Context) {
	var book books.Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		http.BadRequestErrorResponse(c, "Cannot bind JSON: "+err.Error())
		return
	}
	_, err = b.s.AddBook(&book)
	if err != nil {
		http.ServerErrorResponse(c, err.Error())
		return
	}
	http.CreatedResponse(c, book)
}
