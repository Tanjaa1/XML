package Model

import (
	"Nistagram/TagMicroservice/Enum"
	"github.com/google/uuid"
)

type Link struct {
	CODE                      uuid.UUID          `json:"id"`
	Name                      string             `json:"name" gorm:"not null"`
	LinkType                  Enum.LinkType      `json:"linkType" gorm:"not null"`
}
