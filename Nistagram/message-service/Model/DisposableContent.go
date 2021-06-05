package Model

import (
	"github.com/google/uuid"
)

type DisposableContent struct {
	ID                        uuid.UUID          `json:"id"`
	IsSeen                    bool               `json:"isSeen" gorm:"not null"`
	Content                   string             `json:"content" gorm:"not null"`
}
