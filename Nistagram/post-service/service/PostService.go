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

func (service *PostService) AddIntoCollection(c *dto.CollectionD) error {
	fmt.Println("Usao u service insert collection")
	fmt.Println(c.Name)
	fmt.Println(c.PostId)
	result,_ := service.Repo.GetCollectionByName(c.Name)
	//var posts []model.PostIdList
	fmt.Println("duzina liste postova kod kolekcije")
	fmt.Println(len(result.Posts))
	for _,p := range result.Posts{
		fmt.Println("Ispisuje p.PostId i c.PostId")
		fmt.Println(p.PostId)
		fmt.Println(c.PostId)
		if p.PostId == c.PostId{
			fmt.Println("Usao u if")
			return fmt.Errorf("This post exists in collection!")
		}
	}
	result.Posts = append(result.Posts, model.PostIdList{PostId: c.PostId})
	fmt.Println(len(result.Posts))
	//collection := model.Collection{Name: result.Name, UserId: result.UserId, Posts: result.}

	r := service.Repo.AddIntoCollection(result)
	return r
}

func (service *PostService) RemoveFromCollection(c *dto.CollectionD) error {
	fmt.Println("Usao u service remove from collection")
	result,_ := service.Repo.GetCollectionByName(c.Name)
	//var posts []model.PostIdList
	fmt.Println("duzina liste postova kod kolekcije")
	fmt.Println(len(result.Posts))
	var index int
	for i,p := range result.Posts{
		if p.PostId == c.PostId{
			index = i
		}
	}
	result.Posts = append(result.Posts[:index], result.Posts[index+1:]...)
	fmt.Println(len(result.Posts))

	r := service.Repo.RemoveFromCollection(result)
	var posts []model.PostIdList
	for _, s := range result.Posts {
		posts = append(posts, model.PostIdList{PostId: s.PostId})
	}
	result2 := model.Collection{Name: result.Name, Posts: posts, UserId: result.UserId}
	service.Repo.CreateCollection(&result2)
	return r
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
	fmt.Println("Ispis broj kolekcija")
	fmt.Println(len(collections))
	if err != nil {
		return []dto.CollectionDTOO{}, err
	}
	var result []dto.CollectionDTOO
	var postList []dto.PostDto
	for _,item := range collections{
		//var post &model.Post{}
		fmt.Println("ispis broj postova u kolekciji")
		fmt.Println(len(item.Posts))
		for _, item1 := range item.Posts {
			post,_ := service.Repo.GetPostById(item1.PostId)

			fmt.Println("Ispis broj slika u postu")
			fmt.Println(len(post.Images))
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

			var likesDto []string
			likes, _ := service.GetLikeByPostId(post.ID, model.LikeType(0))
			for _,item3 := range likes{
				likesDto = append(likesDto, item3.Username)
			}

			var dislikesDto []string
			dislikes, _ := service.GetLikeByPostId(post.ID, model.LikeType(1))
			for _,item3 := range dislikes{
				dislikesDto = append(dislikesDto, item3.Username)
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
				TagsLink: links, HashTags: hashtags, Location: location, CloseFriends: false, PostType: model.ConvertPostTypeToString(post.PostType),
				Likes: likesDto, Dislikes: dislikesDto}

			postList = append(postList, postDTO)

			}


		result =  append(result, dto.CollectionDTOO{Name: item.Name, UserId: item.UserId, Posts: postList})
	}

	return result, nil
}

func (service *PostService) GetCollectionsForProfileByUserId(userId uint) ([]dto.CollectionDTOO, error) {
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

			postDTO := dto.PostDto{Id: post.ID}
			postList = append(postList, postDTO)
		}

		result =  append(result, dto.CollectionDTOO{Name: item.Name, UserId: item.UserId, Posts: postList})
	}

	return result, nil
}

func (service *PostService) CreateLike(dtoo *dto.LikeDTO) error {

	l, _ := service.Repo.GetLikeByUserIdAndPostId(dtoo.UserId,dtoo.PostId)
	if l != nil {
		fmt.Println("ispis convert")
		fmt.Println(dtoo.LikeType)
		fmt.Println(model.ConvertLikeType(dtoo.LikeType))
		fmt.Println(l.LikeType)
		if l.LikeType != model.ConvertLikeType(dtoo.LikeType){
			fmt.Println("Uslo u like")
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
	fmt.Println("Ispis like type service")
	var likesDto []dto.LikeDTO
	for _, item := range likes {
		likesDto = append(likesDto, dto.LikeDTO{PostId: item.PostId, UserId: item.UserId, Username: item.Username,
			LikeType: model.ConvertLikeTypeToString(item.LikeType)})
		fmt.Println("Ispis like tupe u for")
		fmt.Println(item.LikeType)
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

		var likesDto []string
		likes, _ := service.GetLikeByPostId(post.ID, model.LikeType(0))
		for _,item := range likes{
			likesDto = append(likesDto, item.Username)
		}

		var dislikesDto []string
		dislikes, _ := service.GetLikeByPostId(post.ID, model.LikeType(1))
		for _,item := range dislikes{
			dislikesDto = append(dislikesDto, item.Username)
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
			TagsLink: links,HashTags: hashtags,Location: location, CloseFriends: false, PostType: model.ConvertPostTypeToString(post.PostType),
			Likes: likesDto, Dislikes: dislikesDto})
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

			stringVar := strconv.Itoa(int(item.UserId))
			postsDto = append(postsDto, dto.PostDto{Id: item.ID,Images: images, UserId: stringVar, Description: item.Description,
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
			fmt.Println("Ispis ispis like")
			fmt.Println(item.LikeType)
		}

		var dislikesDto []string
		dislikes, _ := service.GetLikeByPostId(it.ID, model.LikeType(1))
		for _,item := range dislikes{
			dislikesDto = append(dislikesDto, item.Username)
			fmt.Println("Ispis ispis dislike")
			fmt.Println(item.LikeType)
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
	//image1, _, err1 :=
	buf := new(bytes.Buffer)
	err = png.Encode(buf, image)
	if err != nil{
		fmt.Println(err)

	}
	send_s3 := buf.Bytes()
	return send_s3
}
