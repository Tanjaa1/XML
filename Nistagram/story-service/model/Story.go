package model

import (
	"github.com/google/uuid"
	"time"
)

type Story struct {
	ID                        uuid.UUID          `json:"id"`
	SeenBy                    []int			     `json:"seenBy" gorm:"not null"`
	Content                   string             `json:"content" gorm:"not null"`
	CloseFriends              bool               `json:"closeFriends" gorm:"not null"`
	TagsIdLink                []int              `json:"tags"`
	LinkId                    int                `json:"link"`
	UserId                    int                `json:"userId" gorm:"not null"`
	HashTagsIdList            []int              `json:"hashTags"`
	CreateDate                time.Time          `json:"createDate" gorm:"not null"`
	LocationId                int			     `json:"locationId"`
	HighlightIdList           []int              `json:"highlightIdList"`
}
