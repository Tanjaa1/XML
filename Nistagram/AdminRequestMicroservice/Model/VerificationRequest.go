package Model

import "github.com/google/uuid"

type VerificationRequest struct {
	ID                        uuid.UUID          `json:"id"`
	RegisteredUserId          int                `json:"senderId" gorm:"not null"`
	DocumentPicture           string             `json:"documentPicture" gorm:"not null"`
}
