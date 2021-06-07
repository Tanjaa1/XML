package model

import (
	"github.com/google/uuid"
)

type RegisteredUser struct {
	ID                       uuid.UUID             `json:"id" gorm:"type:text"`
	Account                  Account               `json:"account" gorm:"foreignKey:id;type:text"`
	Description              string                `json:"description" gorm:"not null;type:text"`
	Website                  string                `json:"website" gorm:"not null;type:text"`
	IsVerified               bool                  `json:"isVerified" gorm:"not null;type:text"`
	IsPrivate                bool                  `json:"isPrivate" gorm:"not null;type:text"`
	AcceptingMessage         bool                  `json:"acceptingMessage" gorm:"not null;type:text"`
	AcceptingTag             bool                  `json:"acceptingTag" gorm:"not null;type:text"`
	UserType                 UserType              `json:"userType" gorm:"not null;type:text"`
	FollowingRequestIdList   []int                 `json:"followingRequestIdList;type:text"`
	RelatedUsers             []RelatedUser         `json:"relatedUsers" gorm:"polymorphic:Owner;type:text"`
	CollectionsIdList        []int                 `json:"collections" gorm:"type:text"`
	CooperationRequestIdList []int                 `json:"cooperationRequestIdList" gorm:"type:text"`
	MessageRequestIdList     []int                 `json:"messageRequestIdList" gorm:"type:text"`
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