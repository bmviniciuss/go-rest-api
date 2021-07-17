package handlers

import (
	"errors"
	httpHelpers "github.com/bmvinicius/go-rest-api/infra/http"
	"github.com/bmvinicius/go-rest-api/services"
	"strconv"

	"github.com/bmvinicius/go-rest-api/models"
	"github.com/gin-gonic/gin"
)

const routerPath = "books"

type BookHandler struct {
	s *services.BookService
}

func NewBookHandler(s *services.BookService) *BookHandler {
	return &BookHandler{s}
}

func (b *BookHandler) getBook(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		httpHelpers.BadRequestErrorResponse(c, "Id must be a integer")
		return
	}

	var book *models.Book
	book, err = b.s.GetBookById(id)

	if err != nil {
		if errors.Is(err, models.ErrProductNotFound) {
			httpHelpers.NotFoundErrorResponse(c, err.Error())
			return
		}
		httpHelpers.ServerErrorResponse(c, err.Error())
		return
	}

	httpHelpers.OkResponse(c, book)
}

func (b *BookHandler) getBooks(c *gin.Context) {
	books, err := b.s.GetBooks()
	if err != nil {
		httpHelpers.ServerErrorResponse(c, err.Error())
		return
	}
	httpHelpers.OkResponse(c, books)
}

func (b *BookHandler) createBook(c *gin.Context) {
	var book models.Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		httpHelpers.BadRequestErrorResponse(c, "Cannot bind JSON: "+err.Error())
		return
	}
	_, err = b.s.AddBook(&book)
	if err != nil {
		httpHelpers.ServerErrorResponse(c, err.Error())
		return
	}
	httpHelpers.CreatedResponse(c, book)
}

func (b *BookHandler) ApplyRoutes(rg *gin.RouterGroup) {
	booksRouterGroup := rg.Group(routerPath)
	booksRouterGroup.GET("/", b.getBooks)
	booksRouterGroup.GET("/:id", b.getBook)
	booksRouterGroup.POST("/", b.createBook)
}
