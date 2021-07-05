package service

import (
	"fmt"
	"post-service/dto"
	"post-service/model"
	"post-service/repository"
	"strconv"
	"strings"
)

type PostService struct {
	Repo *repository.PostRepository
}
func (service *PostService) SearchHashtag(name string) ([] model.Hashtag, error) {
	exists ,err:= service.Repo.TagSearch("%"+name+"%")
	fmt.Print("u repozitorijumu")
	return exists,err
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

func (service *PostService) SearchLocation(name string) ([] model.Location, error) {
	exists ,err:= service.Repo.LocationSearch("%"+name+"%")
	fmt.Print("u repozitorijumu")
	return exists,err
}

func (service *PostService) GetPostByRegisterUser(idR string) ([] model.Post, error) {
	exists ,err:= service.Repo.GetPostByRegisterUser(idR)
	fmt.Print("u repozitorijumu")
	return exists,err
}

func (service *PostService) GetPostByLocation(idR string) ([] model.Post, error) {
	exists ,err:= service.Repo.GetPostByLocation()
	var result [] model.Post

	for item := range exists{
		if strings.Compare(strconv.FormatUint(uint64(exists[item].Location.ID), 10),idR)==0{
			result=append(result, exists[item])
		}
	}
	return result,err
}

func (service *PostService) GetPostByHashtag(idR string) ([] model.Post, error) {
	exists ,err:= service.Repo.GetPostByHashtag(idR)
	//var result [] model.Post
	//
	//for item := range exists{
	//	for item2 := range exists[item].HashTags{
	//		fmt.Println(exists[item].HashTags[item2].ID)
	//		if strings.Compare(strconv.FormatUint(uint64(exists[item].HashTags[item2].ID), 10),idR)==0{
	//			result=append(result, exists[item])
	//			continue
	//		}
	//	}
	//}
	return exists,err
}