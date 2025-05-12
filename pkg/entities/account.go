package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	ID      uuid.UUID `json:"id" gorm:"primaryKey"`
	Email   string    `json:"username" gorm:"unique;not null;index"`
	Profile string    `json:"profile"`

	// Add other fields as necessary

}
