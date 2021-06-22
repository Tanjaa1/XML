package model

import (
	"github.com/google/uuid"
)

type Complaint struct {
	ID                     	uuid.UUID          `json:"id"`
	Offence                	Offence            `json:"offence" gorm:"not null"`
	LinkId                  int                `json:"inappropriateContent" gorm:"not null"`
}

type Offence int
const (
	RACISM Offence = iota
	SEXUAL_CONTENT
	HATE_SPEECH
	SPAM
	SUICIDE
	DRUG_USE
	OTHER
)