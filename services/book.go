package services

import (
	"github.com/bmvinicius/go-rest-api/infra/repositories"
	"github.com/bmvinicius/go-rest-api/models"
)

type BookService struct {
	bookRepo *repositories.BookRepository
}

func NewBookService(br *repositories.BookRepository) *BookService {
	return &BookService{bookRepo: br}
}

func (s *BookService) GetBookById(id int) (*models.Book, error) {
	book, err := s.bookRepo.GetById(id)
	if err != nil {
		return &models.Book{}, err
	}
	return book, nil
}

func (s *BookService) GetBooks() ([]*models.Book, error) {
	books, err := s.bookRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (s *BookService) AddBook(book *models.Book) (*models.Book, error) {
	// TODO: Check if book already exists
	_, err := s.bookRepo.Create(book)
	if err != nil {
		return &models.Book{}, err
	}
	return book, nil
}
