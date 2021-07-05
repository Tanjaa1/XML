package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
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
		filepat = "\\post-service\\files\\"+ filename

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

	err = json.Unmarshal([]byte(data[0]), &post)
	if err != nil{
		fmt.Println(err)
		w.Write([]byte("{\"success\":\"error\"}"))
		return
	}
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

func (handler *PostHandler) AddIntoCollection(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("Ispisuje se id")
	fmt.Println(id)
	id2,err := strconv.ParseInt(id, 10, 64)
	if err != nil{
		fmt.Println(err)
	}
	id3 := int(id2)

	collectionName :=  vars["name"]
	fmt.Println("creating")
	err = handler.Service.AddIntoCollection(id3,collectionName)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

//func uploadFile(w http.ResponseWriter, r *http.Request) (filename string, path string) {
//	r.ParseMultipartForm(1 << 2)
//	file, handler, err := r.FormFile("file")
//	if err != nil {
//		fmt.Println("Error Retrieving the File")
//		fmt.Println(err)
//		return
//	}
//
//	fmt.Println("Prosao FormFile---------------------")
//
//	defer file.Close()
//
//	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
//	fmt.Printf("File Size: %+v\n", handler.Size)
//	fmt.Printf("MIME Header: %+v\n", handler.Header)
//
//	//dst, err := os.OpenFile("./images/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
//	dst, err := os.Create("./files/"+handler.Filename)
//	fmt.Println("Ispis greske******************************")
//	fmt.Println(err)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	fmt.Println("Prosao Create-------------------------------------------")
//
//	defer dst.Close()
//	if _, err := io.Copy(dst, file); err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	filepat := "\\post-service\\files\\"+handler.Filename
//	fmt.Println(filepat)
//	//database := initDB()
//	fmt.Println("izmedju db i insert")
//	//var img = model.I{Filename: handler.Filename, Filepath: filepat}
//	//fmt.Println("Ispis slike:///////////////////////////////////////////////")
//	//fmt.Println(img)
//	//database.Create(&img)
//	//if err != nil {
//	//	panic(err.Error())
//	//}
//	fmt.Println(handler.Filename)
//	fmt.Println(filepat)
//	fmt.Fprintf(w, "Successfully Uploaded File\n"+"")
//	return handler.Filename,filepat
//}
