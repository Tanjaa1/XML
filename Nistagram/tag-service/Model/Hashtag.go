package Model

import (
	"github.com/google/uuid"
)

type Hashtag struct {
	ID                        uuid.UUID          `json:"id"`
	Name                      string             `json:"name" gorm:"not null"`
}
