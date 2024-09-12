package Models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Id         uint      `gorm:"primaryKey" json:"id"`
	Isbn       string    `json:"isbn"`
	Title      string    `json:"title"`
	Director   *Director `gorm:"foreignKey:DirectorID" json:"director" `
	DirectorID uint      `json:"directorID"`
}

func NewMove(isbn string, title string, director *Director) *Movie {
	return &Movie{
		Isbn:     isbn,
		Title:    title,
		Director: director,
	}
}
