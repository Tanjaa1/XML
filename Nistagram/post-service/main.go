package main

import (
	"bufio"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/handlers"
	"gorm.io/driver/mysql"
	"io"
	"log"
	"net/http"
	"os"
	"post-service/model"
	"strings"
	"time"

	//"text/template"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"post-service/handler"
	"post-service/repository"
	"post-service/service"
)

var mySigningKey = []byte("mysupersecretkey")


func initDB() *gorm.DB {
	fmt.Println("Usao u initDB")
	time.Sleep(time.Duration(20) *time.Second)
	dsn := "root:root@tcp(host.docker.internal:3306)/postdb?parseTime=True&charset=utf8&autocommit=false"
	//dsn := "root:root@tcp(127.0.0.1:3306)/mydb?parseTime=True&charset=utf8&autocommit=false"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	//database.AutoMigrate(&Img{})
	database.AutoMigrate(&model.Post{})
	database.AutoMigrate(&model.Image{})
	database.AutoMigrate(&model.Hashtag{})
	database.AutoMigrate(&model.Location{})
	database.AutoMigrate(&model.Link{})
	database.AutoMigrate(&model.Collection{})
	database.AutoMigrate(&model.Comment{})
	database.AutoMigrate(&model.PostIdList{})
	database.AutoMigrate(&model.Like{})
	database.AutoMigrate(&model.Highlight{})
	database.AutoMigrate(&model.HighlightStory{})

	//database.Migrator().CreateConstraint(&model.Collection{}, "postList")
	//database.Migrator().CreateConstraint(&model.Collection{}, "fk_collections_posts")

	database.AutoMigrate(&model.Post{},&model.Location{})

	return database

}

func initRepo(database *gorm.DB) *repository.PostRepository {
	return &repository.PostRepository{Database: database}
}

func initServices(repo *repository.PostRepository) *service.PostService {
	return &service.PostService{Repo: repo}
}

func initHandler(service *service.PostService) *handler.PostHandler {
	return &handler.PostHandler{Service: service}
}

func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Authorization"] == nil {
			fmt.Println("TOKEN JE NIL")
			err := ((http.StatusUnauthorized))
			json.NewEncoder(w).Encode(err)
			return
		}


		token, err := jwt.Parse(strings.Split(r.Header["Authorization"][0], "Bearer ")[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error in parsing")
			}
			return mySigningKey, nil
		})


		if err != nil {
			fmt.Println(err)
			err := ((http.StatusUnauthorized))
			json.NewEncoder(w).Encode(err)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["role"] == "admin" {

				r.Header.Set("Role", "admin")
				handler.ServeHTTP(w, r)
				return

			} else if claims["role"] == "user" {

				r.Header.Set("Role", "user")
				handler.ServeHTTP(w, r)
				return
			} else
			{
				r.Header.Set("Role", "user")
				handler.ServeHTTP(w, r)
				return
			}

			fmt.Println("ULOGA")

		}

		json.NewEncoder(w).Encode((http.StatusUnauthorized))
	}
}


func handleFunc(handler *handler.PostHandler) {
	fmt.Println("Usao u handleFunc")
	router := mux.NewRouter().StrictSlash(true)
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Access-Control-Allow-Headers", "text/plain"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"http://localhost:8081"})
	credential := handlers.AllowCredentials()
	h := handlers.CORS(headers, methods, origins, credential)
	//router.HandleFunc("/", handler.Hello).Methods("GET")
	//router.HandleFunc("/", handler.CreateConsumer).Methods("POST")
	//router.HandleFunc("/verify/{consumerId}", handler.Verify).Methods("GET")

	router.HandleFunc("/upload",IsAuthorized(handler.CreatePost)).Methods("POST")
	router.HandleFunc("/createCollection",IsAuthorized(handler.CreateCollection)).Methods("POST")
	router.HandleFunc("/createHighlight",IsAuthorized(handler.CreateHighlight)).Methods("POST")
	router.HandleFunc("/addIntoCollection",IsAuthorized(handler.AddIntoCollection)).Methods("POST")
	router.HandleFunc("/addIntoHighlight",IsAuthorized(handler.AddIntoHighlight)).Methods("POST")
	router.HandleFunc("/removeFromCollection",IsAuthorized(handler.RemoveFromCollection)).Methods("POST")
	router.HandleFunc("/removeFromHighlight",IsAuthorized(handler.RemoveFromHighlight)).Methods("POST")
	router.HandleFunc("/addComment/{id}",IsAuthorized(handler.AddComment)).Methods("POST")
	router.HandleFunc("/GetCollectionsByUserId/{id}",IsAuthorized(handler.GetCollectionsByUserId)).Methods("GET")
	router.HandleFunc("/GetHighlightsByUserId/{id}",IsAuthorized(handler.GetHighlightsByUserId)).Methods("GET")
	router.HandleFunc("/GetCollectionsForProfileByUserId/{id}",IsAuthorized(handler.GetCollectionsForProfileByUserId)).Methods("GET")
	router.HandleFunc("/GetHighlightsForProfileByUserId/{id}",IsAuthorized(handler.GetHighlightsForProfileByUserId)).Methods("GET")
	router.HandleFunc("/addLike",IsAuthorized(handler.AddLike)).Methods("POST")
	router.HandleFunc("/getLikesByPostId/{id}",handler.GetLikeByPostId).Methods("GET")
	router.HandleFunc("/getPostsByUserId/{id}",handler.GetLikedPostsByUserId).Methods("GET")
	router.HandleFunc("/getStoriesByUserId/{id}",handler.GetStoriesByUserId).Methods("GET")
	router.HandleFunc("/searchLocation/{name}", IsAuthorized(handler.SearchLocation)).Methods("GET")
	router.HandleFunc("/searchHashtag/{name}", IsAuthorized(handler.SearchHashtag)).Methods("GET")
	router.HandleFunc("/getPByUserId/{id}",IsAuthorized(handler.GetPostsByUserId)).Methods("GET")
	router.HandleFunc("/getPostsByLocation/{locationId}/{myId}",handler.GetPostsByLocation).Methods("GET")
	router.HandleFunc("/getStoriesByLocation/{locationId}/{myId}",handler.GetStoriesByLocation).Methods("GET")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), h(router)))
	//log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8080"), h(router)))
}
//--------------------------------------------------------------------------------------------------------------
type Imgpath struct {
	ID       string   `json:"id"`
	ImagePath string   `json:"imgpath"`
}
func setupCorsResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}
var imgpaths[]Imgpath


//// Compile templates on start of the application
//var templates = template.Must(template.ParseFiles("public/upload.html"))
//
//// Display the named template
//func display(w http.ResponseWriter, page string, data interface{}) {
//	templates.ExecuteTemplate(w, page+".html", data)
//}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	// Maximum upload of 10 MB files
	r.ParseMultipartForm(1 << 2)

	// Get handler for filename, size and headers
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	fmt.Println("Prosao FormFile---------------------")

	defer file.Close()

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create file
	//dst, err := os.Create(handler.Filename)
	////////////////////////////////////////////////////
	//dst, err := os.Create(filepath.Join("C:/Users/Sasa/Desktop/aleksa novo/XML/Nistagram/post-service/temp-img", filepath.Base(handler.Filename)))
	//dst, err := os.Create(filepath.Join("C:/Users/Sasa/Desktop/aleksa novo/XML/Nistagram/post-service/images", filepath.Base(handler.Filename)))
	//var outfile *os.File
	//dst, err := os.OpenFile("./images/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	dst, err := os.Create("./files/"+handler.Filename)
	fmt.Println("Ispis greske******************************")
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Prosao Create-------------------------------------------")

	defer dst.Close()
	//if _, err = io.Copy(dst, file); err != nil {
	//	fmt.Println(err)
	//	return
	//}
	////////////////////////////////////////////////////
	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//var written int64
	//if _, err := io.Copy(outfile, dst); err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}

	//w.Write([]byte("uploaded file:" + handler.Filename + ";length:" + strconv.Itoa(int(written))))
	//w.Write([]byte("uploaded file:" + handler.Filename))



	// create a new buffer base on file size
	fInfo, _ := dst.Stat()
	var size int64 = fInfo.Size()
	buf := make([]byte, size)

	// read file content into buffer
	fReader := bufio.NewReader(dst)
	fReader.Read(buf)

	//convert the buffer bytes to base64 string - use buf.Bytes() for new image
	imgBase64Str := base64.StdEncoding.EncodeToString(buf)
	fmt.Fprintf(w,imgBase64Str)
	//fmt.Fprintf(w, imgBase64Str)------------------------------------------------------

	fmt.Println("Izmedju encode i decode")

	//Decoding
	sDec, _ := base64.StdEncoding.DecodeString(imgBase64Str)
	fmt.Println(sDec)
	filepat := "\\post-service\\files\\"+handler.Filename
	fmt.Println(filepat)
	//db, _ := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306)/imgupdb")
	//dsn := "root:root@tcp(host.docker.internal:3306)/mydb?parseTime=True&charset=utf8&autocommit=false"
	//db, _ := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306)/mydb")
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//db.AutoMigrate(&Img{})
	database := initDB()
	fmt.Println("izmedju db i insert")
	var img = Img{Filename: handler.Filename, Filepath: filepat}
	fmt.Println("Ispis slike:///////////////////////////////////////////////")
	fmt.Println(img)
	database.Create(&img)
	//insert, err := db.Query("INSERT INTO img (filename,filepath) VALUES (?,?,?)", handler.Filename,filepat)
	if err != nil {
		panic(err.Error())
	}
	//defer insert.Close()


	fmt.Fprintf(w, "Successfully Uploaded File\n"+"")
}

type Img struct {
	gorm.Model
	Filename        string    `json:"filename"`
	Filepath        string    `json:"filepath"`
}

type BazaSlike struct {
	Database *gorm.DB
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
		uploadFile(w, r)
}

func getimg(w http.ResponseWriter, r *http.Request) {
	var imagepath Imgpath

	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)

	//dsn := "root:root@tcp(host.docker.internal:3306)/mydb?parseTime=True&charset=utf8&autocommit=false"
	//database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//db, _ := sql.Open("mysql", "root:roote@tcp(127.0.0.1:3306)/imgupdb")
	//db, _ := sql.Open("mysql", "root:roote@tcp(host.docker.internal:3306)/imgupdb")
	db, _ := sql.Open("mysql", "root:roote@tcp(host.docker.internal:3306)/mydb")
	row, err := db.Query("select id,filepath from img where id=?", params["id"])
	if err != nil {
		panic(err.Error())
	}else{
		for row.Next(){
			var id string
			var imgpath string
			err2 := row.Scan(&id, &imgpath)
			row.Columns()
			if err2 != nil{
				panic(err2.Error())
			}else{
				imgpath := Imgpath{
					ID:      id,
					ImagePath:    imgpath,
				}
				imagepath = imgpath
			}
		}
	}
	defer row.Close()
	json.NewEncoder(w).Encode(imagepath)

}

func main() {
	fmt.Println("Usao u main")
	database := initDB()
	repo := initRepo(database)
	service := initServices(repo)
	handler := initHandler(service)
	handleFunc(handler)

	fmt.Println("Usao u main")
	// Upload route
	r := mux.NewRouter()
	//r.HandleFunc("/upload", uploadHandler).Methods("POST")
	//load route

	r.HandleFunc("/load/{id}", getimg).Methods("GET","OPTIONS")
	//Listen on port 8080
	//log.Fatal(http.ListenAndServe(":8000", r))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), r))

}
