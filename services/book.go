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
