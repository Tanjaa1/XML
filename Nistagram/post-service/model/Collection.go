package model

import "gorm.io/gorm"

type Collection struct {
	//ID                        uuid.UUID          `json:"id"`
	gorm.Model
	Name                      string			 `json:"content" gorm:"not null"`
	UserId                    int                `json:"userId" gorm:"not null"`
	Posts                     []Post             `json:"posts" gorm:"many2many:collection_posts;"`
}
