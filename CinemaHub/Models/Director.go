package Models

import "gorm.io/gorm"

type Director struct {
	gorm.Model
	ID        uint   `gorm:"primary_key" json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func NewDirector(firstName, lastName string) *Director {
	return &Director{
		FirstName: firstName,
		LastName:  lastName,
	}
}
