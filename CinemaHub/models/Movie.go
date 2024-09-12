package models

import (
	"time"
)

// Movie model
type Movie struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `gorm:"index" json:"deleted_at,omitempty"`
	Isbn       string     `json:"isbn"`
	Title      string     `json:"title"`
	Director   *Director  `gorm:"foreignKey:DirectorID" json:"director"`
	DirectorID uint       `json:"directorID"`
}

// NewMovie creates a new movie instance
func NewMovie(isbn string, title string, director *Director) *Movie {
	return &Movie{
		Isbn:     isbn,
		Title:    title,
		Director: director,
	}
}
