package model

import "gorm.io/gorm"

type RegisteredUser struct {
	//ID                       uuid.UUID             `json:"id"`
	gorm.Model
	Account                  Account               `json:"account" gorm:"foreignKey:ID"`
	//Account                  Account               `json:"account"`
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
	RelatedUsers             []RelatedUser         `json:"relatedUsers" gorm:"many2many:registered_related_users;"`
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

func ConvertUserType(userTypeString string)(userType UserType) {
	if userTypeString == "INFLUENCER" {
		return INFLUENCER
	}else if userTypeString == "SPORTS"{
		return SPORTS
	}else if userTypeString == "NEWSMEDIA"{
		return NEWSMEDIA
	}else if userTypeString == "BUSINESS"{
		return BUSINESS
	}else if userTypeString == "BRAND"{
		return BRAND
	}else if userTypeString == "ORGANIZATION"{
		return ORGANIZATION
	}else if userTypeString == "ARTIST"{
		return ARTIST
	}else if userTypeString == "EDUCATION"{
		return EDUCATION
	}else {
		return NONE
	}
}

func ConvertUserTypeToString(userType UserType)(userTypeString string) {
	if userType == INFLUENCER {
		return "INFLUENCER"
	}else if userType == SPORTS{
		return "SPORTS"
	}else if userType == NEWSMEDIA{
		return "NEWSMEDIA"
	}else if userType == BUSINESS{
		return "BUSINESS"
	}else if userType == BRAND{
		return "BRAND"
	}else if userType == ORGANIZATION{
		return "ORGANIZATION"
	}else if userType == ARTIST{
		return "ARTIST"
	}else if userType == EDUCATION{
		return "EDUCATION"
	}else {
		return "NONE"
	}
}