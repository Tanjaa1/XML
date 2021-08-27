package model

import "gorm.io/gorm"

type Comment struct {
	//ID                        uuid.UUID          `json:"id"`
	gorm.Model
	Content                   string			 `json:"content" gorm:"not null"`
	Username                  string 			 `json:"username" gorm:"not null"`
}