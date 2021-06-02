package Enum

type Gender int
const (
	MALE Gender = iota
	FEMALE
)

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