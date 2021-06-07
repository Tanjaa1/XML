package model

import (
	"github.com/google/uuid"
)

type RegisteredUser struct {
	ID                       uuid.UUID             `json:"id"`
	Account                  Account
	Description              string                `json:"description" gorm:"not null"`
	Website                  string                `json:"website" gorm:"not null"`
	IsVerified               bool                  `json:"isVerified" gorm:"not null"`
	IsPrivate                bool                  `json:"isPrivate" gorm:"not null"`
	AcceptingMessage         bool                  `json:"acceptingMessage" gorm:"not null"`
	AcceptingTag             bool                  `json:"acceptingTag" gorm:"not null"`
	UserType                 UserType              `json:"userType" gorm:"not null"`
	FollowingRequestIdList   []int                 `json:"followingRequestIdList"`
	RelatedUsers             []RelatedUser         `json:"relatedUsers"`
	CollectionsIdList        []int                 `json:"collections"`
	CooperationRequestIdList []int                 `json:"cooperationRequestIdList"`
	MessageRequestIdList     []int                 `json:"messageRequestIdList"`
	HighlightsIdList         []int                 `json:"highlightsIdList"`
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
	NONE
)