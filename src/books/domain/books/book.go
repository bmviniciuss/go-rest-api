package books

import (
	"gorm.io/gorm"
	"time"
)

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
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}
