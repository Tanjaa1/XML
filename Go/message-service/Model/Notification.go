package Model

import (
	"github.com/google/uuid"
)

type Notification struct {
	ID                        uuid.UUID          	 `json:"id"`
	ContentLinkId             int			     	 `json:"content" gorm:"not null"`
	NotificationType          NotificationType  `json:"notificationType" gorm:"not null"`
}

type NotificationType int
const (
	FOLLOW_REQUEST NotificationType = iota
	TAG
	LIKE
	DISLIKE
	COMMENT
)

