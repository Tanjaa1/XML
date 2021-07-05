package service

import (
	"fmt"
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

func (service *PostService) CreateCollection(dtoo *dto.CollectionDTO) error {

	//var images []model.Image
	//for _, item := range imagess {
	//	images = append(images, model.Image{Filename: item.Filename, Filepath: item.Filepath})
	//}
	//
	//location := model.Location{Place: dtoo.Location.Place, City: dtoo.Location.City, Country: dtoo.Location.Country}
	//var hashtags []model.Hashtag
	//for _, item := range dtoo.HashTags {
	//	hashtags = append(hashtags, model.Hashtag{Name: item.Name})
	//}
	//
	//var links []model.Link
	//for _, item := range dtoo.TagsLink {
	//	links = append(links, model.Link{Name: item.Name, LinkType: model.ConvertLinkType(item.LinkType)})
	//}
	//
	//var comments []model.Comment
	//for _, item := range dtoo.Comments {
	//	comments = append(comments, model.Comment{Content: item.Content, AuthorIdLink: item.AuthorIdLink})
	//}
	//
	//post := model.Post{Images: images,Comments: comments,UserId: dtoo.UserId,Description: dtoo.Description,
	//	TagsLink: links,HashTags: hashtags,Location: location, CloseFriends: false}

	//var posts []model.Post
	//for _, item := range dtoo.Posts {
	//	var images []model.Image
	//	for _, item5 := range item.Images {
	//		images = append(images, model.Image{Filename: item5.Filename, Filepath: item5.Filepath})
	//	}
	//	location := model.Location{Place: item.Location.Place, City: item.Location.City, Country: item.Location.Country}
	//	var hashtags []model.Hashtag
	//	for _, item1 := range item.HashTags {
	//		hashtags = append(hashtags, model.Hashtag{Name: item1.Name})
	//	}
	//
	//	var links []model.Link
	//	for _, item2 := range item.TagsLink {
	//		links = append(links, model.Link{Name: item2.Name, LinkType: model.ConvertLinkType(item2.LinkType)})
	//	}
	//
	//	var comments []model.Comment
	//	for _, item3 := range item.Comments {
	//		comments = append(comments, model.Comment{Content: item3.Content, AuthorIdLink: item3.AuthorIdLink})
	//	}
	//	posts = append(posts, model.Post{Images: images,Comments: comments,UserId: dtoo.UserId,Description: item.Description,
	//		TagsLink: links,HashTags: hashtags,Location: location, CloseFriends: false})
	//}
	result,_ := service.Repo.GetCollectionByName(dtoo.Name)
	if result == nil {
		var posts []model.PostIdList
		for _, s := range dtoo.Posts {
			posts = append(posts, model.PostIdList{PostId: s})
		}
		//var posts []int
		//for _,s := range dtoo.Posts {
		//	posts = append(posts, s)
		//}

		collection := model.Collection{Name: dtoo.Name, UserId: dtoo.UserId, Posts: posts}

		service.Repo.CreateCollection(&collection)
		return nil
	}
	return fmt.Errorf("collection with this name already exists")
}

func (service *PostService) AddIntoCollection(postId int, collectionName string) error {

	result,_ := service.Repo.GetCollectionByName(collectionName)
	//var posts []model.PostIdList

	result.Posts = append(result.Posts, model.PostIdList{PostId: postId})
	//collection := model.Collection{Name: result.Name, UserId: result.UserId, Posts: result.}

		service.Repo.AddIntoCollection(result)
		return nil
}

func (service *PostService) AddComment(dtoo *dto.CommentDTO, postId int) error {

	result,_ := service.Repo.GetPostById(postId)
	//var posts []model.PostIdList

	result.Comments = append(result.Comments, model.Comment{Content: dtoo.Content, AuthorIdLink: dtoo.AuthorIdLink})
	//collection := model.Collection{Name: result.Name, UserId: result.UserId, Posts: result.}

	service.Repo.AddComment(result)
	return nil
}