package model

import "gorm.io/gorm"

type Location struct {
	//ID                        uuid.UUID          `json:"id"`
	gorm.Model
	Place                     string             `json:"place" gorm:"not null"`
	City                      string             `json:"city" gorm:"not null"`
	Country                   string             `json:"country" gorm:"not null"`
}
