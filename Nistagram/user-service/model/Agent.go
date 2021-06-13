package model

import "gorm.io/gorm"

type Agent struct {
	//ID                        uuid.UUID          `json:"id"`
	gorm.Model
	UserId                    int                `json:"userId" gorm:"not null"`
	IsRegistrationConfirmed   bool               `json:"isRegistrationConfirmed" gorm:"not null"`
	CampaignIdList            []int              `json:"campaigns"`

}
