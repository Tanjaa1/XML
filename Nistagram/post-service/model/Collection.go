package model

import "gorm.io/gorm"

type Collection struct {
	//ID                        uuid.UUID          `json:"id"`
	gorm.Model
	Name                      string			 `json:"name" gorm:"not null"`
	UserId                    uint                `json:"userId"`
	Posts                     []PostIdList       `json:"posts" gorm:"many2many:collection_postIdList;"`
}

type PostIdList struct {
	gorm.Model
	PostId                    uint			    `json:"postId"`
}

type Highlight struct {
	gorm.Model
	Name                      string			 `json:"name" gorm:"not null"`
	UserId                    uint                `json:"userId"`
	//Posts                     []PostIdList       `json:"posts" gorm:"many2many:highlights_postIdList;"`
}

type HighlightStory struct {
	gorm.Model
	PostId                    uint			    `json:"postId"`
	CollectionName            string	        `json:"collectionName"`
}
