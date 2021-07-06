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
	result := repo.Database.Preload("Location").Exec("Select * from posts").Find(&listResult);
	fmt.Println("repooo")
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
		repo.Database.Preload("Location").Find(&Pom,"id like ?",res[itt].Post)
		//repo.Database.First(&Pom, "id = ?", res[itt].Post)
		//listResult=append(listResult, Pom)
		fmt.Println(Pom.ID)
	}
	//if len(listResult)==0{
	//	return listResult,nil
	//}
	return listResult,result.Error
}

func (repo *PostRepository) GetCommentsByPostId(idH string)  ([] model.Comment, error) {
	var listResult []model.Comment
	var Pom model.Comment
	var res []dto.PostComment
	result := repo.Database.Table("comments_posts").Find(&res,"post_id like ?",idH)
	for itt := range res{
		//repo.Database.Table("comments").Find(&Pom,"id like ?",res[itt].Comment)
		repo.Database.Preload("Comment").Exec("Select * from comments where id like ?",res[itt].Comment).Find(&Pom);
		listResult=append(listResult, Pom)
	}
	if len(listResult)==0{
		return listResult,nil
	}
	return listResult,result.Error
}

func (repo *PostRepository) GetImagesByPostId(idH string)  ([] model.Image, error) {
	var listResult []model.Image
	var Pom model.Image
	var res []dto.PostImages
	result := repo.Database.Table("images_posts").Find(&res,"post_id like ?",idH)
	for itt := range res{
		//repo.Database.Table("images").Find(&Pom,"id like ?",res[itt].Image)
		repo.Database.Preload("Image").Exec("Select * from images where id like ?",res[itt].Image).Find(&Pom);

		listResult=append(listResult, Pom)
	}
	if len(listResult)==0{
		return listResult,nil
	}
	return listResult,result.Error
}

func (repo *PostRepository) GetTagsByPostId(idH string)  ([] model.Link, error) {
	var listResult []model.Link
	var Pom model.Link
	var res []dto.PostTags
	result := repo.Database.Table("tags_link_posts").Find(&res,"post_id like ?",idH)
	for itt := range res{
		//repo.Database.Table("comments").Find(&Pom,"id like ?",res[itt].Tag)
		repo.Database.Preload("Link").Exec("Select * from links where id like ?",res[itt].Tag).Find(&Pom);

		listResult=append(listResult, Pom)
	}
	if len(listResult)==0{
		return listResult,nil
	}
	return listResult,result.Error
}

func (repo *PostRepository) GetHashtagsByPostId(idH string)  ([] model.Hashtag, error) {
	var listResult []model.Hashtag
	var Pom model.Hashtag
	var res []dto.PostHashtag
	fmt.Print("repozitorijum")

	result := repo.Database.Table("hash_tags_posts").Find(&res,"post_id like ?",idH)
	for itt := range res{
		//repo.Database.Table("hashtags").Find(&Pom,"id like ?",res[itt].Hashtag)
		repo.Database.Preload("HashTag").Exec("Select * from hashtags where id like ?",res[itt].Hashtag).Find(&Pom);

		listResult=append(listResult, Pom)
	}
	if len(listResult)==0{
		return listResult,nil
	}
	return listResult,result.Error
}