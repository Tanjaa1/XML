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
	//Likes                     []Link             `json:"likes" gorm:"many2many:likes_posts;"`
	//Dislikes                  []Link             `json:"dislikes" gorm:"many2many:dislikes_posts;"`
	TagsLink                  []Link             `json:"tagsLink" gorm:"many2many:tagsLink_posts;"`
	HashTags                  []Hashtag          `json:"hashTags" gorm:"many2many:hashTags_posts;"`
	Location                  Location           `json:"location" gorm:"foreignKey:ID"`
	//SeenBy                    []int			     `json:"seenBy" gorm:"not null"`
	CloseFriends              bool               `json:"closeFriends"`
	PostType                  PostType			 `json:"postType"`
}

type Like struct {
	gorm.Model
	PostId            int                `json:"postId"`
	UserId            int                `json:"userId"`
	Username          string              `json:"username"`
	LikeType          LikeType           `json:"likeType"`

}

type LikeType int
const (
	LIKE LikeType = iota
	DISLIKE
)

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


func ConvertPostType(postTypeString string)(postType PostType) {
	if postTypeString == "POST" {
		return POST
	}else {
		return STORY
	}
}

func ConvertPostTypeToString(postType PostType)(postTypeString string) {
	if postType == POST {
		return "POST"
	}else {
		return "STORY"
	}
}

func ConvertLikeType(likeTypeString string)(likeType LikeType) {
	if likeTypeString == "LIKE" {
		return LIKE
	}else {
		return DISLIKE
	}
}

func ConvertLikeTypeToString(likeType LikeType)(likeTypeString string) {
	if likeType == LIKE {
		return "LIKE"
	}else {
		return "DISLIKE"
	}
}