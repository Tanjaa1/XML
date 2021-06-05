package model

import (
	"github.com/google/uuid"
)

type Comment struct {
	ID                        uuid.UUID          `json:"id"`
	Content                   string			 `json:"content" gorm:"not null"`
	AuthorIdLink              int 				 `json:"author" gorm:"not null"`
}