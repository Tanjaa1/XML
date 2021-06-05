ModelModelpackage Model

import (
	"github.com/google/uuid"
	"time"
)

type Campaign struct {
	ID       		      uuid.UUID                 `json:"id"`
	TargetGroupIdList     []int                     `json:"targetGroup" gorm:"not null"`
	StartTime             time.Time                 `json:"startTime" gorm:"not null"`
	EndTime               time.Time                 `json:"endTime" gorm:"not null"`
	CollaboratorIdList    []int                     `json:"collaborator" gorm:"not null"`
	Repetition            int                       `json:"repetition"`
	Statistic             []Statistic               `json:"statistic"`                    // da li da ostane ili da bude int????
	ModifiedId            int                       `json:"modifiedId"`
}