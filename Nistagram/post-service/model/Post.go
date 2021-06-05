package model

import (
	"github.com/google/uuid"
	"time"
)

type Post struct {
	ID                        uuid.UUID          `json:"id"`
	Content                   []string			 `json:"content" gorm:"not null"`
	CreateDate                time.Time          `json:"createDate" gorm:"not null"`
	Comments                  []Comment          `json:"comments"`                    // da li da ostane ili ne?????
	Likes                     []int              `json:"likes"`
	Dislikes                  []int              `json:"dislikes"`
	UserId                    int                `json:"userId" gorm:"not null"`
	TagsIdLink                []int              `json:"tagsIdLink"`
	HashTagsId                []int              `json:"hashTagsId"`
	LocationId                int                `json:"locationId"`
	Description               string             `json:"description"`
	CollectionIdList          []int              `json:"collectionIdList"`

}
