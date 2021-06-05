package Model

import "github.com/google/uuid"

type Agent struct {
	ID                        uuid.UUID          `json:"id"`
	UserId                    int                `json:"userId" gorm:"not null"`
	IsRegistrationConfirmed   bool               `json:"isRegistrationConfirmed" gorm:"not null"`
	CampaignIdList            []int              `json:"campaigns"`

}
