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

func (repo *PostRepository) RemoveFromCollection(collection *model.Collection) error {
	err := repo.Database.Where("ID = ?", collection.ID).Preload("Posts").Delete(&model.Collection{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *PostRepository) GetCollectionByName(name string) (*model.Collection, error) {
	collection := &model.Collection{}
	repo.Database.Model(&collection)
	err := repo.Database.Preload("Posts").First(&collection,"name = ?" , name).Error
	fmt.Println("Ispis err")
	fmt.Println(err)
	if err != nil{
		return nil,err
	}
	return collection, nil
}

func (repo *PostRepository) GetPostById(id uint) (*model.Post, error) {
	fmt.Println(id)
	post := &model.Post{}
	repo.Database.Model(&post)
	//repo.Database.First(&collection,"name = ?" , name)
	err := repo.Database.Preload("Images").Preload("Comments").Preload("TagsLink").Preload("HashTagsIdList").First(&post, "ID = ?", id).Error
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

func (repo *PostRepository) GetCollectionsByUserId(id uint) ([]model.Collection, error) {
	fmt.Println("Iispis id korisnika")
	fmt.Println(id)
	var collections []model.Collection
	repo.Database.Model(&collections)
	//repo.Database.First(&collection,"name = ?" , name)
	err := repo.Database.Preload("Posts").Find(&collections, "user_id = ?", id).Error
	if err != nil{
		return nil,err
	}
	fmt.Println("ispis broj kolekcija")
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

func (repo *PostRepository) GetLikeByUserIdAndPostId(userId uint,postId uint) (*model.Like, error) {
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

func (repo *PostRepository) GetLikeByPostId(postId uint, likeType model.LikeType) ([]model.Like, error) {
	fmt.Println("Ispis id -----------------")
	fmt.Println(postId)
	var like []model.Like
	repo.Database.Model(&like)
	err := repo.Database.Find(&like,"post_id = ? and like_type = ?" , postId,likeType).Error
	fmt.Println("Ispis err")
	fmt.Println(err)
	fmt.Println("Ispis like")
	fmt.Println(like)
	if err != nil{
		return nil,err
	}
	return like, nil
}

func (repo *PostRepository) GetLikeByUserId(userId uint) ([]model.Like, error) {
	fmt.Println("Ispis id -----------------")
	fmt.Println(userId)
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

func (repo *PostRepository) GetStoriesByUserId(id uint) ([]model.Post, error) {
	fmt.Println("Ispis model.PostType(1)")
	fmt.Println(model.PostType(1))
	//post := &model.Post{}
	var stories []model.Post
	repo.Database.Model(&stories)
	//repo.Database.First(&collection,"name = ?" , name)
	err := repo.Database.Preload("Images").Preload("Comments").Preload("TagsLink").Preload("HashTagsIdList").Find(&stories, "user_id = ? and post_type = ?", id,model.PostType(1)).Error
	if err != nil{
		return nil,err
	}
	return stories, nil
}

func (repo *PostRepository) GetLocationById(id uint) (*model.Location, error) {
	location := &model.Location{}
	repo.Database.Model(&location)
	err := repo.Database.First(&location,"ID = ?" , id).Error
	fmt.Println("Ispis err")
	fmt.Println(err)
	if err != nil{
		return nil,err
	}
	return location, nil
}

func (repo *PostRepository) GetHashtagById(id uint) (*model.Hashtag, error) {
	hashtag := &model.Hashtag{}
	repo.Database.Model(&hashtag)
	err := repo.Database.First(&hashtag,"ID = ?" , id).Error
	fmt.Println("Ispis err")
	fmt.Println(err)
	if err != nil{
		return nil,err
	}
	return hashtag, nil
}

func (repo *PostRepository) GetHashtagByName(name string) (*model.Hashtag, error) {
	hashtag := &model.Hashtag{}
	repo.Database.Model(&hashtag)
	err := repo.Database.First(&hashtag,"name = ?" , name).Error
	fmt.Println("Ispis err")
	fmt.Println(err)
	if err != nil{
		return nil,err
	}
	return hashtag, nil
}

func (repo *PostRepository) GetLinkById(id uint) (*model.Link, error) {
	link := &model.Link{}
	repo.Database.Model(&link)
	err := repo.Database.First(&link,"ID = ?" , id).Error
	fmt.Println("Ispis err")
	fmt.Println(err)
	if err != nil{
		return nil,err
	}
	return link, nil
}

func (repo *PostRepository) CreateHashtag(hashtag *model.Hashtag) error {
	result := repo.Database.Create(hashtag)
	if result.RowsAffected == 0 {
		return fmt.Errorf("Like not created")
	}else{
		fmt.Println("Like created")
		return  nil
	}
}

func (repo *PostRepository) CreateLink(link *model.Link) error {
	result := repo.Database.Create(link)
	if result.RowsAffected == 0 {
		return fmt.Errorf("Like not created")
	}else{
		fmt.Println("Like created")
		return  nil
	}
}

func (repo *PostRepository) HashtagSearch(name string)  ([] model.Hashtag, error) {
	var listResult []model.Hashtag
	result:=repo.Database.Table("hashtags").Find(&listResult,"name like ?",name)
	return listResult,result.Error
}

func (repo *PostRepository) GetPostsByUserId(id uint) ([]model.Post, error) {
	var posts []model.Post
	repo.Database.Model(&posts)
	//repo.Database.First(&collection,"name = ?" , name)
	err := repo.Database.Preload("Images").Preload("Comments").Preload("TagsLink").Preload("HashTagsIdList").Find(&posts, "user_id = ? and post_type = ?", id,model.PostType(0)).Error
	if err != nil{
		return nil,err
	}
	fmt.Println(len(posts))
	return posts, nil
}