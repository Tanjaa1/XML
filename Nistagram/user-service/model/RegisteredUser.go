package model

import "github.com/google/uuid"

type RegisteredUser struct {
	ID                       uuid.UUID             `json:"id"`
	Account                  Account               `json:"account" gorm:"foreignKey:ID;"`
	Description              string                `json:"description"`
	Website                  string                `json:"website"`
	IsVerified               bool                  `json:"isVerified"`
	IsPrivate                bool                  `json:"isPrivate"`
	AcceptingMessage         bool                  `json:"acceptingMessage"`
	AcceptingTag             bool                  `json:"acceptingTag"`
	UserType                 UserType              `json:"userType"`
	//FollowingRequestIdList   []int                 `json:"followingRequestIdList"`
	//RelatedUsers             []RelatedUser         `json:"relatedUsers" gorm:"polymorphic:Owner;"`
	//RelatedUsers             []RelatedUser         `json:"relatedUsers" gorm:"joinReferences:person_interests;"`
	//RelatedUsers             []int                 `json:"relatedUsers"`
	//CollectionsIdList        []int                 `json:"collections"`
	//CooperationRequestIdList []int                 `json:"cooperationRequestIdList"`
	//MessageRequestIdList     []int                 `json:"messageRequestIdList"`
	//HighlightsIdList         []int                 `json:"highlightsIdList"`
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