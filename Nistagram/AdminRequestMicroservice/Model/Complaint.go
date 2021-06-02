package Model

import (
	"Nistagram/AdminRequestMicroservice/Enum"
	"github.com/google/uuid"
)

type Complaint struct {
	ID                     	uuid.UUID          `json:"id"`
	Offence                	Enum.Offence       `json:"offence" gorm:"not null"`
	LinkId                  int                `json:"inappropriateContent" gorm:"not null"`
}
