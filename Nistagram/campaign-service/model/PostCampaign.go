package model

import "github.com/google/uuid"

type PostCampaign struct {
	ID       		  uuid.UUID                `json:"id"`
	Campaign          Campaign													// da li da bude int ???
	PostId            int                      `json:"content" gorm:"not null"`
	NumberOfViews     int                      `json:"numberOfViews" gorm:"not null"`
	LinkId            int                      `json:"link" gorm:"not null"`
}
