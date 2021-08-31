package dto

type RequestDTO struct {
	Id         		                uint               `json:"id" gorm:"not null"`
	FollowingId                     uint				 `json:"fol1"`
	FollowerId						uint                ` json:"fol2"`
}

