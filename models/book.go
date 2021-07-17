package models

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

var ErrProductNotFound = errors.New("product not found")

type BookRepositoryReader interface {
	GetById(ID uint) (*Book, error)
	GetAll() ([]*Book, error)
}

type BookRepositoryWriter interface {
	Create(book *Book) (*Book, error)
}

type BookRepository interface {
	BookRepositoryReader
	BookRepositoryWriter
}

// TODO: Separate gorm model from domain model

type Book struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	MediumPrice float64        `json:"mediumPrice"`
	Author      string         `json:"author"`
	ImageURL    string         `json:"imageURL"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
