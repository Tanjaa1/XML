package Model

import (
	"github.com/google/uuid"
)

type Highlights struct {
	ID                        uuid.UUID          `json:"id"`
	Name                      string             `json:"name" gorm:"not null"`
	UserId                    int                `json:"userId" gorm:"not null"`
}
