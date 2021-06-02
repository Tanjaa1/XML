package Model

import (
	"github.com/google/uuid"
)

type Request struct {
	ID                        uuid.UUID          `json:"id"`
	SenderId                  int                `json:"senderId" gorm:"not null"`
	ReceiverId                int                `json:"receiverId" gorm:"not null"`
}
