package model

import "github.com/google/uuid"

type MessageRequest struct {
	ID                        uuid.UUID          `json:"id"`
	RequestId                 int                `json:"requestId" gorm:"not null"`
	MessageId                 int                `json:"messageId" gorm:"not null"`
	IsRejected                bool               `json:"isRejected" gorm:"not null"`
}
