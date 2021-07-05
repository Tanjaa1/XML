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

type PostHashtag struct {
	Post uint `json:"post_id"`
	Hashtag uint `json:"hashtag_id"`
}

type PostComment struct {
	Post uint `json:"post_id"`
	Comment uint `json:"comment_id"`
}
type PostImages struct {
	Post uint `json:"post_id"`
	Image uint `json:"image_id"`
}
type PostTags struct {
	Post uint `json:"post_id"`
	Tag uint `json:"tag_id"`
}