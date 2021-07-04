package repository

import (
	"fmt"
	"gorm.io/gorm"
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
