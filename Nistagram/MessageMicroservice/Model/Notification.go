package Model

import (
	"Nistagram/MessageMicroservice/Enum"
	"github.com/google/uuid"
)

type Notification struct {
	ID                        uuid.UUID          	 `json:"id"`
	ContentLinkId             int			     	 `json:"content" gorm:"not null"`
	NotificationType          Enum.NotificationType  `json:"notificationType" gorm:"not null"`
}
