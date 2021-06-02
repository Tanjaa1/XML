package Model

import (
	"github.com/google/uuid"
)

type Link struct {
	CODE                      uuid.UUID          `json:"id"`
	Name                      string             `json:"name" gorm:"not null"`
	LinkType                  LinkType      `json:"linkType" gorm:"not null"`
}

type LinkType int
const (
	USER LinkType = iota
	POST
	STORY
	CAMPAIGN
	HASHTAG
)