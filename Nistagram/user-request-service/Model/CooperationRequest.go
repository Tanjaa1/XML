package Model

import "github.com/google/uuid"

type CooperationRequest struct {
	ID                        uuid.UUID          `json:"id"`
	RequestId                 int                `json:"requestId" gorm:"not null"`
	CampaignId                int                `json:"campaignId" gorm:"not null"`
}
