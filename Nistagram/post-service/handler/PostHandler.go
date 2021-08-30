package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"post-service/model"
	"strconv"

	//"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"post-service/dto"
	"post-service/service"
)


type PostHandler struct {
	Service *service.PostService
}

func (handler *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request){
	fmt.Println("creating post ****************************************************")
	//filename,filepath :=uploadFile(w,r)
	//r.ParseMultipartForm(1 << 2)

	err := r.ParseMultipartForm(100000)
	if err != nil {
		fmt.Println("error parsing multiplepart form", err)
		return
	}

	files := r.MultipartForm.File["file"]
	fmt.Println("Ispis files")
	fmt.Println(files)

	var filename string
	var filepat string
	var  images []dto.ImageDTO
	for i, _ := range files {

		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			fmt.Println("error opening file ", err)
			return
		}

		//filename = path.Ext(files[i].Filename)
		filename = files[i].Filename
		fmt.Println("Ispis filename: " + filename)
		filepat = "files/"+ filename

		image := dto.ImageDTO{Filename: filename,Filepath: filepat}
		images = append(images, image)

		//filename := GetRandomString(10) + ext

		dst, err := os.Create("./files/"+ filename)
		defer dst.Close()
		if err != nil {
			fmt.Println("error creating destination ", err)
			return
		}

		//copy the uploaded file to the destination file
		if _, err := io.Copy(dst, file); err != nil {
			fmt.Println("error copying file", err)
			return
		}

		fmt.Println("Image upload success: ", files[i].Filename)
	}

	fmt.Println("Filename: " + filename)
	fmt.Println(filepat)
	fmt.Println("all are uploaded")

// Ostali podaci
	var post dto.PostDto
	data := r.MultipartForm.Value["data"]
	fmt.Println("ispis data")
	fmt.Println(data)

	err = json.Unmarshal([]byte(data[0]), &post)
	if err != nil{
		fmt.Println(err)
		w.Write([]byte("{\"success\":\"error\"}"))
		return
	}
	fmt.Println("ispis date param location")
	fmt.Println(post.Location)
	err = handler.Service.CreatePost(&post, images)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *PostHandler) CreateCollection(w http.ResponseWriter, r *http.Request){
	fmt.Println("creating")
	var collection dto.CollectionDTO
	fmt.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(&collection)
	fmt.Println(err)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println("prosao decoder")
	err = handler.Service.CreateCollection(&collection)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *PostHandler) CreateHighlight(w http.ResponseWriter, r *http.Request){
	fmt.Println("creating")
	var highlight dto.HighlightDTO
	fmt.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(&highlight)
	fmt.Println(err)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println("prosao decoder")
	err = handler.Service.CreateHighlight(&highlight)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *PostHandler) AddIntoCollection(w http.ResponseWriter, r *http.Request){
	fmt.Println("Usao u kolekciju")
	var collection dto.CollectionD
	fmt.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(&collection)
	fmt.Println(err)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println("prosao decoder")
	err = handler.Service.AddIntoCollection(&collection)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *PostHandler) AddIntoHighlight(w http.ResponseWriter, r *http.Request){
	fmt.Println("Usao u kolekciju")
	var highlight dto.HighlightD
	fmt.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(&highlight)
	fmt.Println(err)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println("prosao decoder")
	err = handler.Service.AddIntoHighlight(&highlight)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *PostHandler) RemoveFromCollection(w http.ResponseWriter, r *http.Request){
	fmt.Println("Usao u remove from collection")
	var collection dto.CollectionD
	fmt.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(&collection)
	fmt.Println(err)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println("prosao decoder")
	err = handler.Service.RemoveFromCollection(&collection)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *PostHandler) RemoveFromHighlight(w http.ResponseWriter, r *http.Request){
	fmt.Println("Usao u remove from highlight")
	var highlight dto.HighlightD
	fmt.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(&highlight)
	fmt.Println(err)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println("prosao decoder")
	err = handler.Service.RemoveFromHighlight(&highlight)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *PostHandler) AddComment(w http.ResponseWriter, r *http.Request){

	var dto dto.CommentDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("Ispisuje se id")
	fmt.Println(id)
	id2,err := strconv.ParseInt(id, 10, 64)
	if err != nil{
		fmt.Println(err)
	}
	id3 := int(id2)

	fmt.Println("creating")
	err = handler.Service.AddComment(&dto,uint(id3))
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *PostHandler) GetCollectionsByUserId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("Ispisuje se id")
	fmt.Println(id)
	id2,err := strconv.ParseInt(id, 10, 64)
	if err != nil{
		fmt.Println(err)
	}
	id3 := int(id2)
	result, err := handler.Service.GetCollectionsByUserId(uint(id3))
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&result)
	//if result == true {
	//	w.WriteHeader(http.StatusOK)
	//}else{
	//	w.WriteHeader(http.StatusOK)
	//}
	//w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(&result)
}

func (handler *PostHandler) GetHighlightsByUserId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("Ispisuje se id")
	fmt.Println(id)
	id2,err := strconv.ParseInt(id, 10, 64)
	if err != nil{
		fmt.Println(err)
	}
	id3 := int(id2)
	result, err := handler.Service.GetHighlightsByUserId(uint(id3))
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&result)
	//if result == true {
	//	w.WriteHeader(http.StatusOK)
	//}else{
	//	w.WriteHeader(http.StatusOK)
	//}
	//w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(&result)
}

func (handler *PostHandler) GetCollectionsForProfileByUserId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("Ispisuje se id")
	fmt.Println(id)
	id2,err := strconv.ParseInt(id, 10, 64)
	if err != nil{
		fmt.Println(err)
	}
	id3 := int(id2)
	result, err := handler.Service.GetCollectionsForProfileByUserId(uint(id3))
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&result)
	//if result == true {
	//	w.WriteHeader(http.StatusOK)
	//}else{
	//	w.WriteHeader(http.StatusOK)
	//}
	//w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(&result)
}

func (handler *PostHandler) GetHighlightsForProfileByUserId(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Usao u handler GetHighlightsForProfileByUserId")
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("Ispisuje se id")
	fmt.Println(id)
	id2,err := strconv.ParseInt(id, 10, 64)
	if err != nil{
		fmt.Println(err)
	}
	id3 := int(id2)
	result, err := handler.Service.GetHighlightsForProfileByUserId(uint(id3))
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&result)
	//if result == true {
	//	w.WriteHeader(http.StatusOK)
	//}else{
	//	w.WriteHeader(http.StatusOK)
	//}
	//w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(&result)
}

func (handler *PostHandler) AddLike(w http.ResponseWriter, r *http.Request){
	fmt.Println("Usao u handler")
	fmt.Println("creating")
	var like dto.LikeDTO
	fmt.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(&like)
	fmt.Println(err)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println("prosao decoder")
	err = handler.Service.CreateLike(&like)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *PostHandler) GetLikeByPostId(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Usao u handler")
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("Ispisuje se id")
	fmt.Println(id)
	id2,err := strconv.ParseInt(id, 10, 64)
	if err != nil{
		fmt.Println(err)
	}
	id3 := int(id2)
	result, err := handler.Service.GetLikeByPostId(uint(id3),model.LikeType(0))

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (handler *PostHandler) GetLikedPostsByUserId(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Usao u handler")
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("Ispisuje se id")
	fmt.Println(id)
	id2,err := strconv.ParseInt(id, 10, 64)
	if err != nil{
		fmt.Println(err)
	}
	id3 := int(id2)
	result, err := handler.Service.GetLikedPostsByUserId(uint(id3))

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (handler *PostHandler) GetStoriesByUserId(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Usao u handler")
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("Ispisuje se id")
	fmt.Println(id)
	id2,err := strconv.ParseInt(id, 10, 64)
	if err != nil{
		fmt.Println(err)
	}
	id3 := int(id2)
	result, err := handler.Service.GetStoriesByUserId(uint(id3))

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (handler *PostHandler) SearchLocation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	result, err := handler.Service.SearchLocation(vars["name"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	if result !=nil {
		w.WriteHeader(http.StatusOK)
	}else{
		w.WriteHeader(http.StatusOK)

	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&result)
}

func (handler *PostHandler) SearchHashtag(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	result, err := handler.Service.SearchHashtag(vars["name"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	if result !=nil {
		w.WriteHeader(http.StatusOK)
	}else{
		w.WriteHeader(http.StatusOK)

	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&result)
}

func (handler *PostHandler) GetPostsByUserId(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Usao u handler")
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("Ispisuje se id")
	fmt.Println(id)
	id2,err := strconv.ParseInt(id, 10, 64)
	if err != nil{
		fmt.Println(err)
	}
	id3 := int(id2)
	result, err := handler.Service.GetPostsByUserId(uint(id3))

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
