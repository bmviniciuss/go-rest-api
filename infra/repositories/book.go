package repositories

import (
	"errors"
	"github.com/bmvinicius/go-rest-api/models"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) GetById(id int) (*models.Book, error) {
	var book models.Book
	err := r.db.First(&book, "id = ?", id).Error
	if err != nil {
		// No book was found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &models.Book{}, models.ErrProductNotFound
		}
		return &models.Book{}, err
	}

	return &book, nil
}

func (r *BookRepository) GetAll() ([]*models.Book, error) {
	var books []*models.Book
	err := r.db.Find(&books).Error

	if err != nil {
		return nil, err
	}

	return books, nil
}

func (r *BookRepository) Create(book *models.Book) (*models.Book, error) {
	err := r.db.Create(&book).Error
	if err != nil {
		return &models.Book{}, err
	}

	return book, nil
}
