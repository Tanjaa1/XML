package model

import (
	"github.com/google/uuid"
)

type RelatedUser struct {
	ID                     uuid.UUID             `json:"id"`
	RegisteredUserId       int					 `json:"registeredUserId" gorm:"not null"`
	Username               string                `json:"username" gorm:"not null"`
	IsCloseFriend          bool                  `json:"isCloseFriend" gorm:"not null"`
	IsBlocked              bool                  `json:"isBlocked" gorm:"not null"`
	IsMuted                bool                  `json:"isMuted" gorm:"not null"`
	Follower               bool                  `json:"follower" gorm:"not null"`
	Following              bool                  `json:"following" gorm:"not null"`
	MessageIdList          []int                 `json:"messages"`

}
