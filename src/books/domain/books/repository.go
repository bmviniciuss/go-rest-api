package books

import (
	"errors"
)

type Repository interface {
	GetById(ID int) (*Book, error)
	GetAll() ([]*Book, error)
	Create(book *Book) (*Book, error)
}

var ErrBookNotFound = errors.New("books not found")
