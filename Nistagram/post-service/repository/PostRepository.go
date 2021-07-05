package repository

import (
	"fmt"
	"gorm.io/gorm"
	"post-service/dto"
	"post-service/model"
)

type PostRepository struct {
	Database *gorm.DB
}


func (repo *PostRepository) CreatePost(post *model.Post) error {
	result := repo.Database.Create(post)
	if result.RowsAffected == 0 {
		return fmt.Errorf("Post not created")
	}else{
		fmt.Println("Post created")
		return  nil
	}
}

func (repo *PostRepository) TagSearch(name string)  ([] model.Hashtag, error) {
	var listResult []model.Hashtag
	result:=repo.Database.Table("hashtags").Find(&listResult,"name like ?",name)
	return listResult,result.Error
}

func (repo *PostRepository) LocationSearch(name string)  ([] model.Location, error) {
	var listResult []model.Location
	result:=repo.Database.Table("locations").Find(&listResult,"place like ? or city like ? or country like ?",name,name,name)
	return listResult,result.Error
}

func (repo *PostRepository) GetPostByRegisterUser(idR string)  ([] model.Post, error) {
	var listResult []model.Post
	result:=repo.Database.Table("posts").Find(&listResult,"user_id like ?",idR)
	return listResult,result.Error
}

func (repo *PostRepository) GetPostByLocation()  ([] model.Post, error) {
	var listResult []model.Post
	result := repo.Database.Exec("Select * from posts").Find(&listResult);
	if len(listResult)==0{
		return listResult,nil
	}
	return listResult,result.Error
}

func (repo *PostRepository) GetPostByHashtag(idH string)  ([] model.Post, error) {
	var listResult []model.Post
	var Pom model.Post
	var res []dto.PostHashtag
	result := repo.Database.Table("hash_tags_posts").Find(&res,"hashtag_id like ?",idH)
	for itt := range res{
		repo.Database.Table("posts").Find(&Pom,"id like ?",res[itt].Post)
		listResult=append(listResult, Pom)
	}
	if len(listResult)==0{
		return listResult,nil
	}
	return listResult,result.Error
}