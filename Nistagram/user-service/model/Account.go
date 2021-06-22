package model

import (
	"gorm.io/gorm"
	"time"

	//"gorm.io/gorm"
)

type Account struct {
	//ID             uuid.UUID `json:"id"`
	gorm.Model
	Name           string    `json:"name" gorm:"not null"`
	Surname        string    `json:"surname" gorm:"not null"`
	DateOfBirth    time.Time `json:"dateOfBirth" gorm:"not null"`
	Email          string    `json:"email" gorm:"unique;not null"`
	Username       string    `json:"username" gorm:"unique;not null"`
	Password       string    `json:"password" gorm:"not null"`
	Gender         Gender    `json:"gender" gorm:"not null"`
	PhoneNumber    string    `json:"phoneNumber" gorm:"not null"`
	//RegisteredUserId uint
}

type Gender int
const (
	MALE Gender = iota
	FEMALE
)

func ConvertGender(genderString string)(gender Gender) {
	if genderString == "MALE" {
		return MALE
	}else{
		return FEMALE
	}
}

func ConvertGenderToString(gender Gender)(genderString string) {
	if gender == MALE {
		return "MALE"
	}else{
		return "FEMALE"
	}
}
//func (account *Account) BeforeCreate(scope *gorm.DB) error {
//	account.ID = uuid.New()
//	return nil
//}
