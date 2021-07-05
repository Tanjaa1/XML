package model

import "gorm.io/gorm"

type Administrator struct {
	//ID            uuid.UUID        `json:"id"`
	gorm.Model
	AccountId     int              `json:"accountId" gorm:"not null"`
}
