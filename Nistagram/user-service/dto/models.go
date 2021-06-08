package dto

type RequestRegisteredUser struct {
	Account                  string
	Description              string                `json:"description" gorm:"not null"`
	Website                  string                `json:"website" gorm:"not null"`
	IsVerified               bool                  `json:"isVerified" gorm:"not null"`
	IsPrivate                bool                  `json:"isPrivate" gorm:"not null"`
	AcceptingMessage         bool                  `json:"acceptingMessage" gorm:"not null"`
	AcceptingTag             bool                  `json:"acceptingTag" gorm:"not null"`
	UserType                 string                `json:"userType" gorm:"not null"`
	//FollowingRequestIdList   []int                 `json:"followingRequestIdList"`
	//RelatedUsers             []string              `json:"relatedUsers"`
	//RelatedUsers             []int                 `json:"relatedUsers"`
	//CollectionsIdList        []int                 `json:"collections"`
	//CooperationRequestIdList []int                 `json:"cooperationRequestIdList"`
	//MessageRequestIdList     []int                 `json:"messageRequestIdList"`
	//HighlightsIdList         []int                 `json:"highlightsIdList"`
}

type ResponseId struct {
	Id int `json:"id"`
}
