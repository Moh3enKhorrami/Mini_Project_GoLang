package models

import (
	"time"
)

// Director model
type Director struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
}

// NewDirector creates a new director instance
func NewDirector(firstName, lastName string) *Director {
	return &Director{
		FirstName: firstName,
		LastName:  lastName,
	}
}
