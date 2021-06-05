package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"

	//"gorm.io/gorm"
)

type Account struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name" gorm:"not null"`
	Surname        string    `json:"surname" gorm:"not null"`
	DateOfBirth    time.Time `json:"dateOfBirth" gorm:"not null"`
	Email          string    `json:"email" gorm:"unique;not null"`
	Username       string    `json:"username" gorm:"unique;not null"`
	Password       string    `json:"password" gorm:"not null"`
	Gender         Gender    `json:"gender" gorm:"not null"`
	PhoneNumber    string    `json:"phoneNumber" gorm:"not null"`
}

type Gender int
const (
	MALE Gender = iota
	FEMALE
)

func (account *Account) BeforeCreate(scope *gorm.DB) error {
	account.ID = uuid.New()
	return nil
}
