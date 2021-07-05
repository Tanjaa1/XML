package dto

type PostDto struct {
	Images                   []ImageDTO 		 `json:"images"`
	Comments                 []CommentDTO        `json:"comments"`
	UserId                    int                `json:"userId"`
	Description               string             `json:"description"`
	//Likes                     []int              `json:"likes"`
	//Dislikes                  []int              `json:"dislikes"`
	TagsLink                  []LinkDTO          `json:"tagsLink"`
	HashTags                  []HashtagDTO       `json:"hashTags"`
	Location                  LocationDTO        `json:"location"`
	//SeenBy                    []int			     `json:"seenBy" gorm:"not null"`
	CloseFriends              bool               `json:"closeFriends"`
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
	Name                      string             `json:"name"`
	LinkType                  string             `json:"linkType"`
}

type HashtagDTO struct {
	Name                      string             `json:"name"`
}

type LocationDTO struct {
	Place                     string             `json:"place"`
	City                      string             `json:"city"`
	Country                   string             `json:"country"`
}