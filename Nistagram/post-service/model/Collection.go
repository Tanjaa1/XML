package model

import "github.com/google/uuid"

type Collection struct {
	ID                        uuid.UUID          `json:"id"`
	Name                      string			 `json:"content" gorm:"not null"`
	UserId                    int                `json:"userId" gorm:"not null"`
}
