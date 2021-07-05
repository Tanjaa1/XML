package model

import "gorm.io/gorm"

type Link struct {
	//CODE                      uuid.UUID          `json:"id"`
	gorm.Model
	Name                      string             `json:"name" gorm:"not null"`
	LinkType                  LinkType           `json:"linkType" gorm:"not null"`
}

type LinkType int
const (
	USER_LINK LinkType = iota
	POST_LINK
	STORY_LINK
	CAMPAIGN_LINK
	HASHTAG_LINK
)

func ConvertLinkType(linkTypeString string)(linkType LinkType) {
	if linkTypeString == "USER_LINK" {
		return USER_LINK
	}else if linkTypeString == "POST_LINK"{
		return POST_LINK
	}else if linkTypeString == "STORY_LINK"{
		return STORY_LINK
	}else if linkTypeString == "CAMPAIGN_LINK"{
		return CAMPAIGN_LINK
	}else{
		return HASHTAG_LINK
	}
}