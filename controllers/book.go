package controllers

import (
	"github.com/bmvinicius/go-rest-api/infra/repositories"
	"github.com/bmvinicius/go-rest-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BookController struct {
	br *repositories.BookRepository
	// TODO: add Service
}

func NewBookController(br *repositories.BookRepository) *BookController {
	return &BookController{br}
}

func (b *BookController) GetBook(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Id must be a integer",
		})
		return
	}

	var book *models.Book
	book, err = b.br.GetById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Book not found: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func (b *BookController) CreateBook(c *gin.Context) {
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

func (b *BookController) GetBooks(c *gin.Context) {
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
