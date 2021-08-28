package model

import (
	"gorm.io/gorm"
)

type Request struct {
	gorm.Model
	SenderId                  int                `json:"senderId" gorm:"not null"`
	ReceiverId                int                `json:"receiverId" gorm:"not null"`
}
