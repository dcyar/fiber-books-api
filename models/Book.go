package models

import (
	"gorm.io/gorm"
	"time"
)

type Book struct {
	ID        uint           `json:"id"`
	Title     string         `json:"title"`
	Author    string         `json:"author"`
	Year      int            `json:"year" validate:"required,number"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"json:"deleted_at"`
}
