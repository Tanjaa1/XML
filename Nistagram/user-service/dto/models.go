package dto

type RequestRegisteredUser struct {
	Account                  RequestAccount        `json:"account"`
	Description              string                `json:"description"`
	Website                  string                `json:"website"`
	IsVerified               bool                  `json:"isVerified"`
	IsPrivate                bool                  `json:"isPrivate"`
	AcceptingMessage         bool                  `json:"acceptingMessage"`
	AcceptingTag             bool                  `json:"acceptingTag"`
	UserType                 string                `json:"userType"`
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

type RequestAccount struct {
	Name           string    `json:"name" gorm:"not null"`
	Surname        string    `json:"surname" gorm:"not null"`
	DateOfBirth    string    `json:"dateOfBirth" gorm:"not null"`
	Email          string    `json:"email" gorm:"unique;not null"`
	Username       string    `json:"username" gorm:"unique;not null"`
	Password       string    `json:"password" gorm:"not null"`
	Gender         string    `json:"gender" gorm:"not null"`
	PhoneNumber    string    `json:"phoneNumber" gorm:"not null"`
}