package service

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"os"
	"post-service/dto"
	"post-service/model"
	"post-service/repository"
	"strconv"
	"time"
)

type PostService struct {
	Repo *repository.PostRepository
}

func (service *PostService) CreatePost(dtoo *dto.PostDto, imagess []dto.ImageDTO) error {

	var images []model.Image
	for _, item := range imagess {
		images = append(images, model.Image{Filename: item.Filename, Filepath: item.Filepath})
	}

	var hashtags []model.HashtagIdList
	for _, item := range dtoo.HashTags {
		h,_ := service.Repo.GetHashtagByName(item.Name)
		if h != nil {
			hashtags = append(hashtags, model.HashtagIdList{HashtagId: h.ID})
		}else{
			hashtag := model.Hashtag{Name: item.Name}
			service.Repo.CreateHashtag(&hashtag)
			hashtags = append(hashtags, model.HashtagIdList{HashtagId: hashtag.ID})
		}
	}

	var links []model.Link
	for _, item := range dtoo.TagsLink {
		links = append(links, model.Link{Name: item.Name, LinkType: model.ConvertLinkType(item.LinkType)})
	}

	var comments []model.Comment
	for _, item := range dtoo.Comments {
		comments = append(comments, model.Comment{Content: item.Content, Username: item.Username})
	}
	intVar, _ := strconv.Atoi(dtoo.UserId)
	intVa := uint(intVar)
	post := model.Post{Images: images,Comments: comments,UserId: intVa,Description: dtoo.Description,
		TagsLink: links,HashTagsIdList: hashtags,LocationId: dtoo.Location.Id, CloseFriends: false, PostType: model.ConvertPostType(dtoo.PostType)}

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
	fmt.Println("ispis name")
	fmt.Println(dtoo.Name)
	fmt.Println("ispis result")
	fmt.Println(result)
	if result == nil {
		var posts []model.PostIdList
		for _, s := range dtoo.Posts {
			posts = append(posts, model.PostIdList{PostId: uint(s)})
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

func (service *PostService) AddIntoCollection(postId uint, collectionName string) error {

	result,_ := service.Repo.GetCollectionByName(collectionName)
	//var posts []model.PostIdList

	result.Posts = append(result.Posts, model.PostIdList{PostId: postId})
	//collection := model.Collection{Name: result.Name, UserId: result.UserId, Posts: result.}

		service.Repo.AddIntoCollection(result)
		return nil
}

func (service *PostService) AddComment(dtoo *dto.CommentDTO, postId uint) error {

	result,_ := service.Repo.GetPostById(postId)
	//var posts []model.PostIdList

	result.Comments = append(result.Comments, model.Comment{Content: dtoo.Content, Username: dtoo.Username})
	//collection := model.Collection{Name: result.Name, UserId: result.UserId, Posts: result.}

	service.Repo.AddComment(result)
	return nil
}

func (service *PostService) GetCollectionsByUserId(userId uint) ([]dto.CollectionDTOO, error) {
	collections, err := service.Repo.GetCollectionsByUserId(userId)
	if err != nil {
		return []dto.CollectionDTOO{}, err
	}
	var result []dto.CollectionDTOO
	var postList []dto.PostDto
	for _,item := range collections{
		//var post &model.Post{}
		for _, item1 := range item.Posts {
			post,_ := service.Repo.GetPostById(item1.PostId)

			var images []dto.ImageDTO
			for _, item2 := range post.Images {
				images = append(images, dto.ImageDTO{Filename: item2.Filename,Img: service.ConvertImgToBytes(item2.Filepath)})
			}

			l, _ := service.Repo.GetLocationById(post.LocationId)
			location := dto.LocationDTO{Place: l.Place, City: l.City, Country: l.Country}
			var hashtags []dto.HashtagDTO
			for _, item3 := range post.HashTagsIdList {
				h,_ := service.Repo.GetHashtagById(item3.HashtagId)
				hashtags = append(hashtags, dto.HashtagDTO{Name: h.Name})
			}

			var links []dto.LinkDTO
			for _, item4 := range post.TagsLink {
				links = append(links, dto.LinkDTO{Name: item4.Name, LinkType: model.ConvertLinkTypeToString(item4.LinkType)})
			}

			var comments []dto.CommentDTO
			for _, item5 := range post.Comments {
				comments = append(comments, dto.CommentDTO{Content: item5.Content, Username: item5.Username})
			}
			stringVar := strconv.Itoa(int(post.UserId))
			postDTO := dto.PostDto{Id: post.ID,Images: images, Comments: comments, UserId: stringVar, Description: post.Description,
				TagsLink: links, HashTags: hashtags, Location: location, CloseFriends: false, PostType: model.ConvertPostTypeToString(post.PostType)}

			postList = append(postList, postDTO)

			}


		result =  append(result, dto.CollectionDTOO{Name: item.Name, UserId: item.UserId, Posts: postList})
	}

	return result, nil
}

func (service *PostService) CreateLike(dtoo *dto.LikeDTO) error {

	l, _ := service.Repo.GetLikeByUserIdAndPostId(dtoo.UserId,dtoo.PostId)
	if l != nil {
		if l.LikeType != model.ConvertLikeType(dtoo.LikeType){
			l.LikeType = model.ConvertLikeType(dtoo.LikeType)
			err1 := service.Repo.UpdateLike(l)
			return err1
		}else {
			return fmt.Errorf("Like already exists")
		}
	}else{
		like := model.Like{PostId: dtoo.PostId,UserId: dtoo.UserId,Username: dtoo.Username,
			LikeType: model.ConvertLikeType(dtoo.LikeType)}
		result := service.Repo.CreateLike(&like)
		return result
	}
}

func (service *PostService) GetLikeByPostId(postId uint, liketype model.LikeType) ([]dto.LikeDTO,error) {

	fmt.Println("Usao u service")
	likes, _ := service.Repo.GetLikeByPostId(postId, liketype)
	var likesDto []dto.LikeDTO
	for _, item := range likes {
		likesDto = append(likesDto, dto.LikeDTO{PostId: item.PostId, UserId: item.UserId, Username: item.Username,
			LikeType: model.ConvertLikeTypeToString(item.LikeType)})
	}

	return likesDto,nil
}

func (service *PostService) GetLikedPostsByUserId(userId uint) ([]dto.PostDto,error) {

	fmt.Println("Usao u service")
	likes, _ := service.Repo.GetLikeByUserId(userId)
	fmt.Println("Ispis lisla likes")
	fmt.Println(len(likes))
	fmt.Println(likes)
	var postsDto []dto.PostDto
	for _, item := range likes {
		post,_ := service.Repo.GetPostById(item.PostId)

		var images []dto.ImageDTO
		for _, item := range post.Images {
			images = append(images, dto.ImageDTO{Filename: item.Filename, Img: service.ConvertImgToBytes(item.Filepath)})
		}

		l, _ := service.Repo.GetLocationById(post.LocationId)
		location := dto.LocationDTO{Place: l.Place, City: l.City, Country: l.Country}

		var hashtags []dto.HashtagDTO
		for _, item := range post.HashTagsIdList {
			h,_ := service.Repo.GetHashtagById(item.HashtagId)
			hashtags = append(hashtags, dto.HashtagDTO{Name: h.Name})
		}

		var links []dto.LinkDTO
		for _, item := range post.TagsLink {
			links = append(links, dto.LinkDTO{Name: item.Name, LinkType: model.ConvertLinkTypeToString(item.LinkType)})
		}

		var comments []dto.CommentDTO
		for _, item := range post.Comments {
			comments = append(comments, dto.CommentDTO{Content: item.Content, Username: item.Username})
		}
		stringVar := strconv.Itoa(int(post.UserId))
		postsDto = append(postsDto, dto.PostDto{Id: post.ID,Images: images,Comments: comments,UserId: stringVar,Description: post.Description,
			TagsLink: links,HashTags: hashtags,Location: location, CloseFriends: false, PostType: model.ConvertPostTypeToString(post.PostType)})
	}

	return postsDto,nil
}

func (service *PostService) GetStoriesByUserId(userId uint) ([]dto.PostDto,error) {

	fmt.Println("Usao u service")
	stories, _ := service.Repo.GetStoriesByUserId(userId)
	fmt.Println(len(stories))
	var postsDto []dto.PostDto
	for _, item := range stories {
		timein := item.CreatedAt.Add(time.Hour * 24)
		//timein := item.CreatedAt.AddDate(0, 0, 1)(time.Hour * 24 + time.Minute * 1 + time.Second * 1)
		fmt.Println("Ispis timr in")
		fmt.Println(timein)
		t := timein.Sub(time.Now())
		if t>0 {

			//post,_ := service.Repo.GetStoriesById(item.PostId)

			var images []dto.ImageDTO
			for _, item := range item.Images {
				images = append(images, dto.ImageDTO{Filename: item.Filename, Img: service.ConvertImgToBytes(item.Filepath)})
			}

			l, _ := service.Repo.GetLocationById(item.LocationId)
			location := dto.LocationDTO{Place: l.Place, City: l.City, Country: l.Country}

			var hashtags []dto.HashtagDTO
			for _, item1 := range item.HashTagsIdList {
				h,_ := service.Repo.GetHashtagById(item1.HashtagId)
				hashtags = append(hashtags, dto.HashtagDTO{Name: h.Name})
			}

			var links []dto.LinkDTO
			for _, item1 := range item.TagsLink {
				links = append(links, dto.LinkDTO{Name: item1.Name, LinkType: model.ConvertLinkTypeToString(item1.LinkType)})
			}

			var comments []dto.CommentDTO
			for _, item1 := range item.Comments {
				comments = append(comments, dto.CommentDTO{Content: item1.Content, Username: item1.Username})
			}
			stringVar := strconv.Itoa(int(item.UserId))
			postsDto = append(postsDto, dto.PostDto{Id: item.ID,Images: images, Comments: comments, UserId: stringVar, Description: item.Description,
				TagsLink: links, HashTags: hashtags, Location: location, CloseFriends: false, PostType: model.ConvertPostTypeToString(item.PostType)})
		}
	}

	return postsDto,nil
}

func (service *PostService) SearchLocation(name string) ([] dto.LocationDTO, error) {
	exists ,err:= service.Repo.LocationSearch("%"+name+"%")
	var locations []dto.LocationDTO
	for _, item := range exists {
		locations = append(locations, dto.LocationDTO{Id: item.ID,Place: item.Place, City: item.City, Country: item.Country})
	}
	fmt.Print("u repozitorijumu")
	return locations,err
}

func (service *PostService) SearchHashtag(name string) ([] dto.HashtagDTO, error) {
	exists ,err:= service.Repo.HashtagSearch("%"+name+"%")
	var hashtags []dto.HashtagDTO
	for _, item := range exists {
		hashtags = append(hashtags, dto.HashtagDTO{Id: item.ID,Name: item.Name})
	}
	fmt.Print("u repozitorijumu")
	return hashtags,err
}

func (service *PostService) GetPostsByUserId(userId uint) ([]dto.PostDto,error) {

	fmt.Println("Usao u service")
	posts, _ := service.Repo.GetPostsByUserId(userId)
	fmt.Println("Ispis lisla likes")
	fmt.Println(len(posts))
	fmt.Println(posts)
	var postsDto []dto.PostDto
	for _, it := range posts {

		var images []dto.ImageDTO
		for _, item := range it.Images {
			images = append(images, dto.ImageDTO{Filename: item.Filename, Img: service.ConvertImgToBytes(item.Filepath)})
			fmt.Println("Ispis slika")
			fmt.Println(images[0].Img)
		}

		l, _ := service.Repo.GetLocationById(it.LocationId)
		location := dto.LocationDTO{Place: l.Place, City: l.City, Country: l.Country}

		var hashtags []dto.HashtagDTO
		for _, item := range it.HashTagsIdList {
			h,_ := service.Repo.GetHashtagById(item.HashtagId)
			hashtags = append(hashtags, dto.HashtagDTO{Name: h.Name})
		}

		var likesDto []string
		likes, _ := service.GetLikeByPostId(it.ID, model.LikeType(0))
		for _,item := range likes{
			likesDto = append(likesDto, item.Username)
		}

		var dislikesDto []string
		dislikes, _ := service.GetLikeByPostId(it.ID, model.LikeType(1))
		for _,item := range dislikes{
			dislikesDto = append(dislikesDto, item.Username)
		}

		var links []dto.LinkDTO
		for _, item := range it.TagsLink {
			links = append(links, dto.LinkDTO{Name: item.Name, LinkType: model.ConvertLinkTypeToString(item.LinkType)})
		}

		var comments []dto.CommentDTO
		for _, item := range it.Comments {
			comments = append(comments, dto.CommentDTO{Content: item.Content, Username: item.Username})
		}
		stringVar := strconv.Itoa(int(it.UserId))
		postsDto = append(postsDto, dto.PostDto{Id: it.ID,Images: images,Comments: comments,UserId: stringVar,Description: it.Description,
			TagsLink: links,HashTags: hashtags,Location: location, CloseFriends: false, PostType: model.ConvertPostTypeToString(it.PostType),
			Likes: likesDto, Dislikes: dislikesDto})
	}

	return postsDto,nil
}

//func (handler *PostService) ConvertAllPhotos(posts []*model.Post,photos []*model.Image) []*dto.PostDto {
//	var res []*dto.PostDto
//	for _,post := range posts{
//		postDTO := dto.PostDto{PostID: post.PostID,Username: post.Username}
//		for _, pp := range photos {
//			if(pp.PostID == post.PostID) {
//				println("*****")
//				postDTO.PhotoDTO = append(postDTO.PhotoDTO, &dto.PhotoDTO{Img: handler.ConvertImgToBytes(pp.Path)})
//			}
//		}
//		res = append(res, &postDTO)
//	}
//
//	return res
//}

func (handler *PostService) ConvertImgToBytes(filePath string)  []byte{
	fmt.Println("start")
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Ispis greske err u convert")
		fmt.Println(err)
		return nil

	}
	fmt.Println("hshhshdh")
	defer f.Close()
	image, _, err := image.Decode(f)
	image1, _, err1 :=
	buf := new(bytes.Buffer)
	err = png.Encode(buf, image)
	if err != nil{
		fmt.Println(err)

	}
	send_s3 := buf.Bytes()
	return send_s3
}
