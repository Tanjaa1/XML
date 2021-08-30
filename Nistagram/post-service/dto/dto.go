package dto

type PostDto struct {
	Id                       uint                `json:"id"`
	Images                   []ImageDTO 		 `json:"images"`
	Comments                 []CommentDTO        `json:"comments"`
	UserId                    string             `json:"userId"`
	Description               string             `json:"description"`
	Likes                     []string           `json:"likes"`
	Dislikes                  []string           `json:"dislikes"`
	TagsLink                  []LinkDTO          `json:"tagsLink"`
	HashTags                  []HashtagDTO       `json:"hashTags"`
	Location                  LocationDTO        `json:"location"`
	//SeenBy                    []int			     `json:"seenBy" gorm:"not null"`
	CloseFriends              bool               `json:"closeFriends"`
	PostType                  string			 `json:"postType"`
}

type ImageDTO struct {
	Filename        string    `json:"filename"`
	Filepath        string    `json:"filepath"`
	Img             []byte    `json:"img"`
}

type CommentDTO struct {
	Content                   string			 `json:"content"`
	Username                  string 		     `json:"username"`
}

type LinkDTO struct {
	Id                        uint               `json:"id" gorm:"not null"`
	Name                      string             `json:"name"`
	LinkType                  string             `json:"linkType"`
}

type HashtagDTO struct {
	Id                        uint               `json:"id" gorm:"not null"`
	Name                      string             `json:"name"`
}

type LocationDTO struct {
	Id                        uint               `json:"id" gorm:"not null"`
	Place                     string             `json:"place"`
	City                      string             `json:"city"`
	Country                   string             `json:"country"`
}

type CollectionDTO struct {
	Name                      string			 `json:"name"`
	UserId                    uint                `json:"userId"`
	Posts                     []int              `json:"posts"`
}

type HighlightDTO struct {
	Name                      string			 `json:"name"`
	UserId                    uint                `json:"userId"`
	Posts                     []int              `json:"posts"`
}

type CollectionDTOO struct {
	Name                      string			 `json:"name"`
	UserId                    uint                `json:"userId"`
	Posts                     []PostDto          `json:"posts"`
}

type HighlightDTOO struct {
	Name                      string			 `json:"name"`
	UserId                    uint                `json:"userId"`
	Posts                     []PostDto          `json:"posts"`
}

type CollectionD struct {
	Name                      string			 `json:"name"`
	PostId                    uint                `json:"postId"`
}

type HighlightD struct {
	Name                      string			 `json:"name"`
	PostId                    uint                `json:"postId"`
}

type LikeDTO struct {
	PostId            uint                 `json:"postId"`
	UserId            uint                 `json:"userId"`
	Username          string              `json:"username"`
	LikeType          string              `json:"linkType"`

}