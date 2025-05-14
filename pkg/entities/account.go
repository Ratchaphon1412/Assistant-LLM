package entities

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	ID      uuid.UUID `json:"id" gorm:"primaryKey"`
	Email   string    `json:"email" gorm:"unique;not null;index"`
	Profile string    `json:"profile"`

	// Add other fields as necessary

}

func (a *Account) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.NewV4()
	return
}
