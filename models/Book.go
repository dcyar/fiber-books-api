package models

import (
	"time"
)

type Book struct {
	ID        uint      `gorm:"primaryKey" json:"id" form:"id"`
	AuthorID  int       `json:"author_id" form:"author_id"`
	Author    Author    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"author"`
	Title     string    `json:"title" form:"title"`
	Year      int       `json:"year" form:"year"`
	Cover     string    `json:"cover,omitempty" form:"cover,omitempty"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
