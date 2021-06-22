package model

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Images                   []Image			 `json:"images" gorm:"many2many:images_posts;"`
	Comments                  []Comment          `json:"comments" gorm:"many2many:comments_posts;"`
	UserId                    int                `json:"userId"`
	Description               string             `json:"description"`
	//Likes                     []int              `json:"likes"`
	//Dislikes                  []int              `json:"dislikes"`
	TagsLink                  []Link             `json:"tagsLink" gorm:"many2many:tagsLink_posts;"`
	HashTags                  []Hashtag          `json:"hashTags" gorm:"many2many:hashTags_posts;"`
	Location                  Location           `json:"location" gorm:"foreignKey:ID"`
	//SeenBy                    []int			     `json:"seenBy" gorm:"not null"`
	CloseFriends              bool               `json:"closeFriends"`
}

type PostType int
const (
	POST PostType = iota
	STORY
)

type Image struct {
	gorm.Model
	Filename        string    `json:"filename"`
	Filepath        string    `json:"filepath"`
}


func (post *Post) AddHashtag(item Hashtag){
	post.HashTags = append(post.HashTags, item)
}
