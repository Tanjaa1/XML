package Model

import (
	"github.com/google/uuid"
)

type Location struct {
	ID                        uuid.UUID          `json:"id"`
	Place                     string             `json:"place" gorm:"not null"`
	City                      string             `json:"city" gorm:"not null"`
	Country                   string             `json:"country" gorm:"not null"`
}
