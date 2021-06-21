package model

import "gorm.io/gorm"

type Hashtag struct {
	//ID                        uuid.UUID          `json:"id"`
	gorm.Model
	Name                      string             `json:"name" gorm:"not null"`
}
