package model

import "gorm.io/gorm"

type RelatedUser struct {
	//ID                     uuid.UUID             `json:"id"`
	gorm.Model
	RegisteredUserId       int					 `json:"registeredUserId"`
	Username               string                `json:"username"`
	IsCloseFriend          bool                  `json:"isCloseFriend"`
	IsBlocked              bool                  `json:"isBlocked"`
	IsMuted                bool                  `json:"isMuted"`
	Follower               bool                  `json:"follower"`
	Following              bool                  `json:"following"`
	//MessageIdList          []int                 `json:"messages"`

}
