package books

import (
	"errors"
	"github.com/bmvinicius/go-rest-api/src/books/domain/books"
	"gorm.io/gorm"
)

type BookGormRepository struct {
	db *gorm.DB
}

func NewBookGormRepository(db *gorm.DB) books.Repository {
	return BookGormRepository{db: db}
}

func (r BookGormRepository) GetById(id int) (*books.Book, error) {
	var book books.Book
	err := r.db.First(&book, "id = ?", id).Error
	if err != nil {
		// No books was found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &books.Book{}, books.ErrBookNotFound
		}
		return &books.Book{}, err
	}

	return &book, nil
}

func (r BookGormRepository) GetAll() ([]*books.Book, error) {
	var b []*books.Book
	err := r.db.Find(&b).Error
	if err != nil {
		return []*books.Book{}, err
	}
	return b, nil
}

func (r BookGormRepository) Create(book *books.Book) (*books.Book, error) {
	err := r.db.Create(&book).Error
	if err != nil {
		return &books.Book{}, err
	}
	return book, nil
}
