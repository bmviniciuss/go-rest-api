package handlers

import (
	"errors"
	http_helpers "github.com/bmvinicius/go-rest-api/infra/http"
	"github.com/bmvinicius/go-rest-api/services"
	"net/http"
	"strconv"

	"github.com/bmvinicius/go-rest-api/infra/repositories"
	"github.com/bmvinicius/go-rest-api/models"
	"github.com/gin-gonic/gin"
)

const routerPath = "books"

type BookHandler struct {
	br *repositories.BookRepository // TODO: remove repository
	s  *services.BookService
}

func NewBookHandler(br *repositories.BookRepository, s *services.BookService) *BookHandler {
	return &BookHandler{br, s}
}

func (b *BookHandler) getBook(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		http_helpers.BadRequestErrorResponse(c, "Id must be a integer")
		return
	}

	var book *models.Book
	book, err = b.s.GetBookById(id)

	if err != nil {
		if errors.Is(err, models.ErrProductNotFound) {
			http_helpers.NotFoundErrorResponse(c, err.Error())
			return
		}
		http_helpers.ServerErrorResponse(c, err.Error())
		return
	}

	http_helpers.OkResponse(c, book)
}

func (b *BookHandler) createBook(c *gin.Context) {
	var book models.Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Cannot bind JSON: " + err.Error(),
		})
		return
	}

	_, err = b.br.Create(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Cannot create Book: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": book,
	})
}

func (b *BookHandler) getBooks(c *gin.Context) {
	books, err := b.br.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error: " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"data": books,
	})
}

func (b *BookHandler) ApplyRoutes(rg *gin.RouterGroup) {
	booksRouterGroup := rg.Group(routerPath)
	booksRouterGroup.GET("/", b.getBooks)
	booksRouterGroup.GET("/:id", b.getBook)
	booksRouterGroup.POST("/", b.createBook)
}
