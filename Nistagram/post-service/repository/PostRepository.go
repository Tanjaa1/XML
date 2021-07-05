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

func (repo *PostRepository) CreateCollection(collection *model.Collection) error {
	result := repo.Database.Create(collection)
	if result.RowsAffected == 0 {
		return fmt.Errorf("Post not created")
	}else{
		fmt.Println("Post created")
		return  nil
	}
}

func (repo *PostRepository) AddIntoCollection(collection *model.Collection) error {
		err := repo.Database.Save(collection).Error
		if err != nil {
			return err
		}
		return nil
}

func (repo *PostRepository) GetCollectionByName(name string) (*model.Collection, error) {
	collection := &model.Collection{}
	repo.Database.Model(&collection)
	repo.Database.First(&collection,"name = ?" , name)
	return collection, nil
}

func (repo *PostRepository) GetPostById(id int) (*model.Post, error) {
	post := &model.Post{}
	repo.Database.Model(&post)
	//repo.Database.First(&collection,"name = ?" , name)
	err := repo.Database.Preload("Images").Preload("Comments").Preload("TagsLink").Preload("HashTags").Preload("Location").First(&post, "ID = ?", id).Error
	fmt.Println("Images: ")
	fmt.Println( len(post.Images))
	fmt.Println("Comments:")
	fmt.Println(len(post.Comments))
	fmt.Println("tAgsLink:")
	fmt.Println(len(post.TagsLink))
	fmt.Println("Hashtags:")
	fmt.Println(len(post.HashTags))
	fmt.Println("LoCATION")
	fmt.Println(post.Location.Place)
	if err != nil{
		return nil,err
	}
	return post, nil
}

func (repo *PostRepository) AddComment(post *model.Post) error {
	err := repo.Database.Save(post).Error
	if err != nil {
		return err
	}
	return nil
}