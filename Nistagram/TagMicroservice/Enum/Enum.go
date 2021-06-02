package Enum

type LinkType int
const (
	USER LinkType = iota
	POST
	STORY
	CAMPAIGN
	HASHTAG
)
