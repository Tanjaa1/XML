package service

import (
	"post-service/dto"
	"post-service/model"
	"post-service/repository"
)

type PostService struct {
	Repo *repository.PostRepository
}

func (service *PostService) CreatePost(dtoo *dto.PostDto, imagess []dto.ImageDTO) error {

	var images []model.Image
	for _, item := range imagess {
		images = append(images, model.Image{Filename: item.Filename, Filepath: item.Filepath})
	}

	location := model.Location{Place: dtoo.Location.Place, City: dtoo.Location.City, Country: dtoo.Location.Country}
	var hashtags []model.Hashtag
	for _, item := range dtoo.HashTags {
		hashtags = append(hashtags, model.Hashtag{Name: item.Name})
	}

	var links []model.Link
	for _, item := range dtoo.TagsLink {
		links = append(links, model.Link{Name: item.Name, LinkType: model.ConvertLinkType(item.LinkType)})
	}

	var comments []model.Comment
	for _, item := range dtoo.Comments {
		comments = append(comments, model.Comment{Content: item.Content, AuthorIdLink: item.AuthorIdLink})
	}

	post := model.Post{Images: images,Comments: comments,UserId: dtoo.UserId,Description: dtoo.Description,
		TagsLink: links,HashTags: hashtags,Location: location, CloseFriends: false}

	service.Repo.CreatePost(&post)
	return nil
}
