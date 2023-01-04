package models

import (
	"time"
)

type Book struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	AuthorID  int       `json:"author_id"`
	Author    Author    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"author"`
	Title     string    `json:"title"`
	Year      int       `json:"year" validate:"required,number"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
