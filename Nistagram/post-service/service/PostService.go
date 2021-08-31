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

//func (service *PostService) AddIntoHighlight(c *dto.HighlightD) error {
//	fmt.Println("Usao u service insert highlight")
//	fmt.Println(c.Name)
//	fmt.Println(c.PostId)
//	result,_ := service.Repo.GetHighlightByName(c.Name)
//	//var posts []model.PostIdList
//	fmt.Println("duzina liste postova kod hajlajta")
//	fmt.Println(len(result.Posts))
//	for _,p := range result.Posts{
//		fmt.Println("Ispisuje p.PostId i c.PostId")
//		fmt.Println(p.PostId)
//		fmt.Println(c.PostId)
//		if p.PostId == c.PostId{
//			fmt.Println("Usao u if")
//			return fmt.Errorf("This post exists in highlight!")
//		}
//	}
//	result.Posts = append(result.Posts, model.PostIdList{PostId: c.PostId})
//	fmt.Println(len(result.Posts))
//	//collection := model.Collection{Name: result.Name, UserId: result.UserId, Posts: result.}
//
//	r := service.Repo.AddIntoHighlight(result)
//	return r
//}

func (service *PostService) AddIntoHighlight(c *dto.HighlightD) error {
	fmt.Println("Usao u service insert highlight")
	fmt.Println(c.Name)
	fmt.Println(c.PostId)
	//result,_ := service.Repo.GetHighlightByName(c.Name)
	////var posts []model.PostIdList
	//fmt.Println("duzina liste postova kod hajlajta")
	//fmt.Println(len(result.Posts))
	//for _,p := range result.Posts{
	//	fmt.Println("Ispisuje p.PostId i c.PostId")
	//	fmt.Println(p.PostId)
	//	fmt.Println(c.PostId)
	//	if p.PostId == c.PostId{
	//		fmt.Println("Usao u if")
	//		return fmt.Errorf("This post exists in highlight!")
	//	}
	//}
	//result.Posts = append(result.Posts, model.PostIdList{PostId: c.PostId})
	//fmt.Println(len(result.Posts))
	////collection := model.Collection{Name: result.Name, UserId: result.UserId, Posts: result.}
	result := model.HighlightStory{PostId: c.PostId, CollectionName: c.Name}

	r := service.Repo.AddIntoHighlightStory(&result)
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

//func (service *PostService) RemoveFromHighlight(c *dto.HighlightD) error {
//	fmt.Println("Usao u service remove from highlight")
//	result,_ := service.Repo.GetHighlightByName(c.Name)
//	//var posts []model.PostIdList
//	fmt.Println("duzina liste postova kod hajlajta")
//	fmt.Println(len(result.Posts))
//	var index int
//	for i,p := range result.Posts{
//		if p.PostId == c.PostId{
//			index = i
//		}
//	}
//	result.Posts = append(result.Posts[:index], result.Posts[index+1:]...)
//	fmt.Println(len(result.Posts))
//
//	r := service.Repo.RemoveFromHighlight(result)
//	var posts []model.PostIdList
//	for _, s := range result.Posts {
//		posts = append(posts, model.PostIdList{PostId: s.PostId})
//	}
//	result2 := model.Highlight{Name: result.Name, Posts: posts, UserId: result.UserId}
//	service.Repo.CreateHighlight(&result2)
//	return r
//}

func (service *PostService) RemoveFromHighlight(c *dto.HighlightD) error {
	fmt.Println("Usao u service remove from highlight")
	fmt.Println(c.Name)
	fmt.Println(c.PostId)
	result,_ := service.Repo.GetHighlightStoriesByName(c.Name)
	fmt.Println(len(result))
	var err error
	for _,item := range result{
		if item.PostId == c.PostId {
			err = service.Repo.RemoveHighlightStory(item.ID)
			if err != nil {
				return err
			}
			return err
		}
	}
	return err
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

func (service *PostService) GetHighlightsByUserId(userId uint) ([]dto.HighlightDTOO, error) {
	fmt.Println("Usao u GetHighlightsByUserId")
	var result []dto.HighlightDTOO
	highlights, err := service.Repo.GetHighlightsByUserId(userId)
	fmt.Println("Ispis broj hajlajta")
	fmt.Println(len(highlights))
	if err != nil {
		return []dto.HighlightDTOO{}, err
	}
	//var highlightStorise []model.HighlightStory
	ha:= service.Repo.GetHighlightStories()
	fmt.Println("Ispis duzina ha ++++++++++++++")
	fmt.Println(ha)
	for _,it := range highlights{
		var h []model.HighlightStory
		for _,it2 := range ha{
			if it.Name == it2.CollectionName {
				h = append(h, it2)
			}
		}
		fmt.Println("Ispis duzina h")
		fmt.Println(len(h))
		if len(h) == 0{
			result = append(result, dto.HighlightDTOO{Name: it.Name, UserId: userId})
		}else {
			var postList []dto.PostDto
			for _, item1 := range h {
				post, _ := service.Repo.GetPostById(item1.PostId)

				fmt.Println("Ispis broj slika u postu")
				fmt.Println(len(post.Images))
				var images []dto.ImageDTO
				for _, item2 := range post.Images {
					images = append(images, dto.ImageDTO{Filename: item2.Filename, Img: service.ConvertImgToBytes(item2.Filepath)})
				}

				l, _ := service.Repo.GetLocationById(post.LocationId)
				location := dto.LocationDTO{Place: l.Place, City: l.City, Country: l.Country}
				var hashtags []dto.HashtagDTO
				for _, item3 := range post.HashTagsIdList {
					h, _ := service.Repo.GetHashtagById(item3.HashtagId)
					hashtags = append(hashtags, dto.HashtagDTO{Name: h.Name})
				}

				var links []dto.LinkDTO
				for _, item4 := range post.TagsLink {
					links = append(links, dto.LinkDTO{Name: item4.Name, LinkType: model.ConvertLinkTypeToString(item4.LinkType)})
				}

				stringVar := strconv.Itoa(int(post.UserId))
				postDTO := dto.PostDto{Id: post.ID, Images: images, UserId: stringVar, Description: post.Description,
					TagsLink: links, HashTags: hashtags, Location: location, CloseFriends: false,
					PostType: model.ConvertPostTypeToString(post.PostType)}



					postList = append(postList, postDTO)
			}
			result = append(result, dto.HighlightDTOO{Name: it.Name, UserId: userId, Posts: postList})
		}

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

func (service *PostService) GetHighlightsForProfileByUserId(userId uint) ([]dto.HighlightDTOO, error) {
	fmt.Println("Usao u service GetHighlightsForProfileByUserId")
	var result []dto.HighlightDTOO
	highlights, err := service.Repo.GetHighlightsByUserId(userId)
	fmt.Println("Ispis broj hajlajta")
	fmt.Println(len(highlights))
	if err != nil {
		return []dto.HighlightDTOO{}, err
	}
	//var highlightStorise []model.HighlightStory
	ha:= service.Repo.GetHighlightStories()
	fmt.Println("Ispis duzina ha ++++++++++++++")
	fmt.Println(ha)
	for _,it := range highlights{
		var h []model.HighlightStory
		for _,it2 := range ha{
			if it.Name == it2.CollectionName {
				h = append(h, it2)
			}
		}
		fmt.Println("Ispis duzina h")
		fmt.Println(len(h))
		if len(h) == 0{
			result = append(result, dto.HighlightDTOO{Name: it.Name, UserId: userId})
		}else {
			var postList []dto.PostDto
			for _, item1 := range h {
				postDTO := dto.PostDto{Id: item1.PostId}
				postList = append(postList, postDTO)
			}
			result = append(result, dto.HighlightDTOO{Name: it.Name, UserId: userId, Posts: postList})
		}
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
		fmt.Println("Ispis t")
		fmt.Println(t.Hours())
		if t.Hours()>0 {
			fmt.Println("Usao u if")
			var images []dto.ImageDTO
			for _, item := range item.Images {
				images = append(images, dto.ImageDTO{Filename: item.Filename, Img: service.ConvertImgToBytes(item.Filepath)})
			}

			fmt.Println("nakon slike")

			l, _ := service.Repo.GetLocationById(item.LocationId)
			location := dto.LocationDTO{Place: l.Place, City: l.City, Country: l.Country}

			fmt.Println("Nakon lokacije")

			var hashtags []dto.HashtagDTO
			for _, item1 := range item.HashTagsIdList {
				h,_ := service.Repo.GetHashtagById(item1.HashtagId)
				hashtags = append(hashtags, dto.HashtagDTO{Name: h.Name})
			}

			fmt.Println("Nakon hashtag")

			var links []dto.LinkDTO
			for _, item1 := range item.TagsLink {
				links = append(links, dto.LinkDTO{Name: item1.Name, LinkType: model.ConvertLinkTypeToString(item1.LinkType)})
			}

			fmt.Println("Nakon links")

			stringVar := strconv.Itoa(int(item.UserId))
			postsDto = append(postsDto, dto.PostDto{Id: item.ID,Images: images, UserId: stringVar, Description: item.Description,
				TagsLink: links, HashTags: hashtags, Location: location, CloseFriends: false, PostType: model.ConvertPostTypeToString(item.PostType)})

			fmt.Println("Kreirao story")
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

func (service *PostService) CreateHighlight(dtoo *dto.HighlightDTO) error {
	fmt.Println("Ulazak u highlight service")
	result,_ := service.Repo.GetHighlightByName(dtoo.Name)
	fmt.Println("ispis name")
	fmt.Println(dtoo.Name)
	fmt.Println("ispis result")
	fmt.Println(result)
	if result == nil {

		highlight := model.Highlight{Name: dtoo.Name, UserId: dtoo.UserId}
		service.Repo.CreateHighlight(&highlight)

		for _, s := range dtoo.Posts {
			//posts = append(posts, model.PostIdList{PostId: uint(s)})
			highlightStory := model.HighlightStory{PostId: uint(s), CollectionName: dtoo.Name}
			service.Repo.CreateHighlightStory(&highlightStory)

		}

		return nil
	}
	return fmt.Errorf("Highlight with this name already exists")
}

func (service *PostService) GetPostsByLocation(locationId uint, myId uint) ([] dto.PostDto,error) {
	fmt.Println("Usao u getPostsByLocation")
	postsByLocation,_ := service.Repo.GetPostsByLocationId(locationId)
	var postsDto []dto.PostDto
	for _,it := range postsByLocation{
		//resp, err := http.Get("http://"+profileHost+":"+profilePort+"/get-by-id/" + strconv.Itoa(int(profileId)))
		fmt.Println("Get 1 88888888888888888888888888888")
		//resp,err := http.Get("http://localhost:8080/api/user/checkPublic/" + strconv.Itoa(int(myId)) + "/" + strconv.Itoa(int(it.UserId)))
		fmt.Println("Get 2 ffffffffffffffffffffff")
		//client := &http.Client{  Timeout: time.Second * 10,
		//	}
		//	fmt.Println("Get 2 posle")
		//resp1, err := client.Get("http://localhost:8080/api/user/checkPublic/" + strconv.Itoa(int(myId)) + "/" + strconv.Itoa(int(it.UserId)))
		//fmt.Println("izasao")
		//var rez bool
		//rez = true
		//body, err := ioutil.ReadAll(resp.Body)
		//if err != nil{
		//	fmt.Println("Greska -/-/-/-/----/-/----/-/-/-/--")
		//	fmt.Println(err)
		//	return nil,err
		//}
		//fmt.Println("Body: ", body)
		//defer resp.Body.Close()
		//err = json.Unmarshal(body, &rez)
		//fmt.Println("Ispis rez  zzzzzzzzzzzzzzzzzzzzzzzzzzz")
		//fmt.Println(rez)
		//
		//var rez1 bool
		//rez1 = true
		//body1, err := ioutil.ReadAll(resp1.Body)
		//if err != nil{
		//	fmt.Println("Greska1 -/-/-/-/----/-/----/-/-/-/--")
		//	fmt.Println(err)
		//	return nil,err
		//}
		//fmt.Println("Body1: ", body1)
		//defer resp.Body.Close()
		//err = json.Unmarshal(body1, &rez1)
		//fmt.Println("Ispis rez1  zzzzzzzzzzzzzzzzzzzzzzzzzzz")
		//fmt.Println(rez1)

		//if rez{

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
	//}
	return postsDto,nil
}


func (service *PostService) GetStoriesByLocation(locationId uint, myId uint) ([] dto.PostDto,error) {
	fmt.Println("Usao u getPostsByLocation")
	postsByLocation,_ := service.Repo.GetStoryByLocationId(locationId)
	var postsDto []dto.PostDto
	for _,it := range postsByLocation{
		timein := it.CreatedAt.Add(time.Hour * 24)
		//timein := item.CreatedAt.AddDate(0, 0, 1)(time.Hour * 24 + time.Minute * 1 + time.Second * 1)
		fmt.Println("Ispis timr in")
		fmt.Println(timein)
		t := timein.Sub(time.Now())
		fmt.Println("Ispis t")
		fmt.Println(t.Hours())
		if t.Hours()>0 {
			//resp, err := http.Get("http://"+profileHost+":"+profilePort+"/get-by-id/" + strconv.Itoa(int(profileId)))
			fmt.Println("Get 1 88888888888888888888888888888")
			//resp,err := http.Get("http://localhost:8080/api/user/checkPublic/" + strconv.Itoa(int(myId)) + "/" + strconv.Itoa(int(it.UserId)))
			fmt.Println("Get 2 ffffffffffffffffffffff")
			//client := &http.Client{  Timeout: time.Second * 10,
			//	}
			//	fmt.Println("Get 2 posle")
			//resp1, err := client.Get("http://localhost:8080/api/user/checkPublic/" + strconv.Itoa(int(myId)) + "/" + strconv.Itoa(int(it.UserId)))
			//fmt.Println("izasao")
			//var rez bool
			//rez = true
			//body, err := ioutil.ReadAll(resp.Body)
			//if err != nil{
			//	fmt.Println("Greska -/-/-/-/----/-/----/-/-/-/--")
			//	fmt.Println(err)
			//	return nil,err
			//}
			//fmt.Println("Body: ", body)
			//defer resp.Body.Close()
			//err = json.Unmarshal(body, &rez)
			//fmt.Println("Ispis rez  zzzzzzzzzzzzzzzzzzzzzzzzzzz")
			//fmt.Println(rez)
			//
			//var rez1 bool
			//rez1 = true
			//body1, err := ioutil.ReadAll(resp1.Body)
			//if err != nil{
			//	fmt.Println("Greska1 -/-/-/-/----/-/----/-/-/-/--")
			//	fmt.Println(err)
			//	return nil,err
			//}
			//fmt.Println("Body1: ", body1)
			//defer resp.Body.Close()
			//err = json.Unmarshal(body1, &rez1)
			//fmt.Println("Ispis rez1  zzzzzzzzzzzzzzzzzzzzzzzzzzz")
			//fmt.Println(rez1)

			//if rez{

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
	}
	return postsDto,nil
}