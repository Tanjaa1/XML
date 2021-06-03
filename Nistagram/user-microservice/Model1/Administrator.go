package Model1

import "github.com/google/uuid"

type Administrator struct {
	ID            uuid.UUID        `json:"id"`
	AccountId     int              `json:"accountId" gorm:"not null"`
}
