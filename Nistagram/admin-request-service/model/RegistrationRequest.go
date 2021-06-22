package model

import (
	"github.com/google/uuid"
)

type RegistrationRequest struct {
	ID                        uuid.UUID          `json:"id"`
	AgentId                   int                `json:"senderId" gorm:"not null"`
}
