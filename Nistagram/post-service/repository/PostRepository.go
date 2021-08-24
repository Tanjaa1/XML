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
	fmt.Println("Usao u repo")
	fmt.Println(collection.Name)
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
	err := repo.Database.First(&collection,"name = ?" , name).Error
	fmt.Println("Ispis err")
	fmt.Println(err)
	if err != nil{
		return nil,err
	}
	return collection, nil
}

func (repo *PostRepository) GetPostById(id int) (*model.Post, error) {
	fmt.Println(id)
	post := &model.Post{}
	repo.Database.Model(&post)
	//repo.Database.First(&collection,"name = ?" , name)
	err := repo.Database.Preload("Images").Preload("Comments").Preload("TagsLink").Preload("HashTags").Preload("Location").First(&post, "ID = ?", id).Error
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

func (repo *PostRepository) GetCollectionsByUserId(id int) ([]model.Collection, error) {
	collections := []model.Collection{}
	repo.Database.Model(&collections)
	//repo.Database.First(&collection,"name = ?" , name)
	err := repo.Database.Preload("Posts").Find(&collections, "user_id = ?", id).Error
	if err != nil{
		return nil,err
	}
	fmt.Println(len(collections))
	return collections, nil
}


func (repo *PostRepository) CreateLike(like *model.Like) error {
	result := repo.Database.Create(like)
	if result.RowsAffected == 0 {
		return fmt.Errorf("Like not created")
	}else{
		fmt.Println("Like created")
		return  nil
	}
}

func (repo *PostRepository) GetLikeByUserIdAndPostId(userId int,postId int) (*model.Like, error) {
	like := &model.Like{}
	repo.Database.Model(&like)
	err := repo.Database.First(&like,"user_id = ? and post_id = ?" , userId,postId).Error
	fmt.Println("Ispis err")
	fmt.Println(err)
	fmt.Println("Ispis like")
	fmt.Println(like)
	if err != nil{
		return nil,err
	}
	return like, nil
}

func (repo *PostRepository) LocationSearch(name string)  ([] model.Location, error) {
	var listResult []model.Location
	result:=repo.Database.Table("locations").Find(&listResult,"place like ? or city like ? or country like ?",name,name,name)
	return listResult,result.Error
}

func (repo *PostRepository) UpdateLike(like *model.Like) error {
	err := repo.Database.Save(like).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *PostRepository) GetLikeByPostId(postId int) ([]model.Like, error) {
	var like []model.Like
	repo.Database.Model(&like)
	err := repo.Database.Find(&like,"post_id = ?" , postId).Error
	fmt.Println("Ispis err")
	fmt.Println(err)
	fmt.Println("Ispis like")
	fmt.Println(like)
	if err != nil{
		return nil,err
	}
	return like, nil
}

func (repo *PostRepository) GetLikeByUserId(userId int) ([]model.Like, error) {
	var like []model.Like
	repo.Database.Model(&like)
	err := repo.Database.Find(&like,"user_id = ?" , userId).Error
	fmt.Println("Ispis err")
	fmt.Println(err)
	fmt.Println("Ispis like")
	fmt.Println(like)
	if err != nil{
		return nil,err
	}
	return like, nil
}