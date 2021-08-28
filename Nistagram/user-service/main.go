package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"time"
	"user-service/handler"
	"user-service/model"
	"user-service/repository"
	"user-service/service"
)

var mySigningKey = []byte("mysupersecretkey")


func initDB() *gorm.DB {
	time.Sleep(time.Duration(20) *time.Second)
	dsn := "root:root@tcp(host.docker.internal:3306)/mydb?parseTime=True&charset=utf8&autocommit=false"
	//dsn := "root:root@tcp(127.0.0.1:3306)/mydb?parseTime=True&charset=utf8&autocommit=false"
	//var dbhost, dbport, dbusername, dbpassword string = "localhost", "3306", "root", "root"
	//database, err := gorm.Open(mysql.Open(dbusername + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/profile?charset=utf8mb4&parseTime=True&loc=local&serverTimezone=UTC"))
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//database, err := gorm.Open("db_relational", &gorm.Config{})
	//database, err := sql.Open("db_relational", "root:root@tcp(docker.for.mac.localhost:3306)/mydb?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	database.AutoMigrate(&model.Account{})
	database.AutoMigrate(&model.RegisteredUser{})
	database.AutoMigrate(&model.RelatedUser{})
	database.AutoMigrate(&model.RegisteredUser{},&model.Account{})

	return database

}

func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Authorization"] == nil {
			fmt.Println("TOKEN JE NIL")
			err := ((http.StatusUnauthorized))
			json.NewEncoder(w).Encode(err)
			return
		}


		token, err := jwt.Parse(r.Header["Authorization"][0], func(token *jwt.Token) (interface{}, error) {
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

func initRepo(database *gorm.DB) *repository.RegisteredUserRepository {
	return &repository.RegisteredUserRepository{Database: database}
}

func initServices(repo *repository.RegisteredUserRepository) *service.RegisteredUserService {
	return &service.RegisteredUserService{Repo: repo}
}

func initHandler(service *service.RegisteredUserService) *handler.RegisteredUserHandler {
	return &handler.RegisteredUserHandler{Service: service}
}
func initRepo1(database *gorm.DB) *repository.RelatedUserRepository {
	return &repository.RelatedUserRepository{Database: database}
}

func initServices1(repo *repository.RelatedUserRepository) *service.RelatedUserService {
	return &service.RelatedUserService{Repo: repo}
}

func initHandler1(service *service.RelatedUserService) *handler.RelatedUserHendler {
	return &handler.RelatedUserHendler{Service: service}
}
func handleFunc(handler *handler.RegisteredUserHandler) {
	router := mux.NewRouter().StrictSlash(true)
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Access-Control-Allow-Headers", "text/plain"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"http://localhost:8081"})
	credential := handlers.AllowCredentials()
	h := handlers.CORS(headers, methods, origins, credential)


	//router.HandleFunc("/", handler.Hello).Methods("GET")
	//router.HandleFunc("/", handler.CreateConsumer).Methods("POST")
	//router.HandleFunc("/verify/{consumerId}", handler.Verify).Methods("GET")
	router.HandleFunc("/userRegistration", handler.CreateRegisteredUser).Methods("POST")
	router.HandleFunc("/getMyPersonalData/{id}", handler.GetMyPersonalData).Methods("GET")
	router.HandleFunc("/changeMyPersonalData/{id}", handler.ChangePersonalData).Methods("POST")
	router.HandleFunc("/getAccountByUsername/{username}", handler.GetAccountByUsername).Methods("GET")
	router.HandleFunc("/login/{username}/{password}", handler.Login).Methods("GET")
	router.HandleFunc("/addFollower/{idRegisterUser}/{idRelatedUser}", handler.AddFollower).Methods("PUT")
	router.HandleFunc("/followerExist/{username}/{registerUserId}", handler.FolowerExist).Methods("GET")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), h(router)))
}
func main() {
	database := initDB()
	repo := initRepo(database)
	service := initServices(repo)
	handler := initHandler(service)
	handleFunc(handler)

}

//func main() {
//	quit := make(chan os.Signal)
//	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
//
//	router := mux.NewRouter()
//	router.StrictSlash(true)
//	server, err := NewPostServer()
//	if err != nil {
//		log.Fatal(err.Error())
//	}
//
//	defer server.CloseTracer()
//	defer server.CloseDB()
//
//	//router.HandleFunc("/post/", server.createPostHandler).Methods("POST")
//	//router.HandleFunc("/post/", server.getAllPostsHandler).Methods("GET")
//	//router.HandleFunc("/post/{id:[0-9]+}/", server.getPostHandler).Methods("GET")
//	//router.HandleFunc("/post/{id:[0-9]+}/", server.deletePostHandler).Methods("DELETE")
//
//	router.HandleFunc("/userRegistration/", server.CreateRegisteredUser).Methods("POST")
//
//	// start server
//	srv := &http.Server{Addr: "0.0.0.0:8000", Handler: router}
//	go func() {
//		log.Println("server starting")
//		if err := srv.ListenAndServe(); err != nil {
//			if err != http.ErrServerClosed {
//				log.Fatal(err)
//			}
//		}
//	}()
//
//	<-quit
//
//	log.Println("service shutting down ...")
//
//	// gracefully stop server
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//
//	if err := srv.Shutdown(ctx); err != nil {
//		log.Fatal(err)
//	}
//
//	log.Println("server stopped")
//}
