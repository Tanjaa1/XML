package model

import (
	"github.com/google/uuid"
)

type Statistic struct {
	ID       		  uuid.UUID                 `json:"id"`
	InfluencerId      int                       `json:"influencerId" gorm:"not null"`
	NumberOfAccesses  int                       `json:"NumberOfAccesses"`
}
