package Model

import (
	"github.com/google/uuid"
)

type StoryCampaign struct {
	ID       		  uuid.UUID                 `json:"id"`
	Campaign          Campaign                                                   // da li da bude int ???
	StoryId           int                       `json:"content" gorm:"not null"`
}
