package model

import "gorm.io/gorm"

type Collection struct {
	//ID                        uuid.UUID          `json:"id"`
	gorm.Model
	Name                      string			 `json:"name" gorm:"not null"`
	UserId                    int                `json:"userId"`
	Posts                     []PostIdList       `json:"posts" gorm:"many2many:collection_postIdList;"`
}

type PostIdList struct {
	gorm.Model
	PostId                    int			    `json:"postId"`
}