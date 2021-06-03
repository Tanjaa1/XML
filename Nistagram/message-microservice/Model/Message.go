package Model

import (
	"github.com/google/uuid"
	"time"
)

type Message struct {
	ID                        uuid.UUID          `json:"id"`
	SenderId                  int                `json:"senderId" gorm:"not null"`
	ReceiverId                int                `json:"receiverId" gorm:"not null"`
	Date                      time.Time          `json:"date" gorm:"not null"`
	Content                   string			 `json:"content" gorm:"not null"`
	DisposableContent         DisposableContent
	LinkIdList                []int              `json:"links"`
}
