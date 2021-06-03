package Model1

import (
	"github.com/google/uuid"
)

type RegisteredUser struct {
	ID                       uuid.UUID             `json:"id"`
	AccountId                int                   `json:"accountId" gorm:"not null"`
	Description              string                `json:"description" gorm:"not null"`
	Website                  string                `json:"website" gorm:"not null"`
	IsVerified               bool                  `json:"isVerified" gorm:"not null"`
	IsPrivate                bool                  `json:"isPrivate" gorm:"not null"`
	AcceptingMessage         bool                  `json:"acceptingMessage" gorm:"not null"`
	AcceptingTag             bool                  `json:"acceptingTag" gorm:"not null"`
	UserType                 UserType         `json:"userType" gorm:"not null"`
	FollowingRequestIdList   []int                 `json:"followingRequestIdList"`
	RelatedUsers             []RelatedUser         `json:"relatedUsers"`         //da li da bude tip int ???????
	CollectionsIdList        []int                 `json:"collections"`           //da li da ostane ili ne ?????
	CooperationRequestIdList []int                 `json:"cooperationRequestIdList"`
	MessageRequestIdList     []int                 `json:"messageRequestIdList"`
	HighlightsIdList         []int                 `json:"highlightsIdList"`     //da li da ostane ili ne?????
}

type UserType int
const (
	INFLUENCER UserType = iota
	SPORTS
	NEWSMEDIA
	BUSINESS
	BRAND
	ORGANIZATION
	ARTIST
	EDUCATION
)