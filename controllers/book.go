package controllers

import (
	"github.com/bmvinicius/go-rest-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type BookController struct {
	db *gorm.DB // TODO: remove db and use service to read/write
	// TODO: add Service
}

func NewBookController(db *gorm.DB) *BookController {
	return &BookController{db}
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

	var book models.Book
	err = b.db.First(&book, id).Error

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

	err = b.db.Create(&book).Error
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
	var books []*models.Book
	err := b.db.Find(&books).Error
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
