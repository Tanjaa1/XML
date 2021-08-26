package dto

type PostDto struct {
	Images                   []ImageDTO 		 `json:"images"`
	Comments                 []CommentDTO        `json:"comments"`
	UserId                    string                `json:"userId"`
	Description               string             `json:"description"`
	//Likes                     []int              `json:"likes"`
	//Dislikes                  []int              `json:"dislikes"`
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
}

type CommentDTO struct {
	Content                   string			 `json:"content"`
	AuthorIdLink              int 				 `json:"author"`
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
	UserId                    int                `json:"userId"`
	Posts                     []int              `json:"posts"`
}

type CollectionDTOO struct {
	Name                      string			 `json:"name"`
	UserId                    int                `json:"userId"`
	Posts                     []PostDto          `json:"posts"`
}

type LikeDTO struct {
	PostId            int                 `json:"postId"`
	UserId            int                 `json:"userId"`
	Username          string              `json:"username"`
	LikeType          string              `json:"linkType"`

}