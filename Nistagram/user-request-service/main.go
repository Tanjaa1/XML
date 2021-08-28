package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"log"
	"net/http"
	"os"
	"time"
	"user-request-service/model"

	"gorm.io/gorm"

	"user-request-service/handler"

	"user-request-service/repository"
	"user-request-service/service"
)


func initDB() *gorm.DB {
	time.Sleep(time.Duration(20) *time.Second)
	dsn := "root:root@tcp(host.docker.internal:3306)/mydb3?parseTime=True&charset=utf8&autocommit=false"
	//dsn := "root:root@tcp(127.0.0.1:3306)/mydb?parseTime=True&charset=utf8&autocommit=false"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	database.AutoMigrate(&model.Request{})
	database.AutoMigrate(&model.MessageRequest{})
	database.AutoMigrate(&model.CooperationRequest{})


	//database.Migrator().CreateConstraint(&model.Collection{}, "postList")
	//database.Migrator().CreateConstraint(&model.Collection{}, "fk_collections_posts")


	return database
}

func initRepo(database *gorm.DB) *repository.ConsumerRepository {
	return &repository.ConsumerRepository{Database: database}
}

func initServices(repo *repository.ConsumerRepository) *service.ConsumerService {
	return &service.ConsumerService{Repo: repo}
}

func initHandler(service *service.ConsumerService) *handler.ConsumerHandler {
	return &handler.ConsumerHandler{Service: service}
}
func handleFunc(handler *handler.ConsumerHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", handler.Hello).Methods("GET")
	router.HandleFunc("/", handler.CreateConsumer).Methods("POST")
	router.HandleFunc("/verify/{consumerId}", handler.Verify).Methods("GET")
	router.HandleFunc("/createRequest/{idFollowing}/{idFollower}", handler.CreateRequest).Methods("POST")
	router.HandleFunc("/getRequestById/{idRequest}", handler.GetRequestById).Methods("GET")
	router.HandleFunc("/deleteRequest/{idRequest}", handler.DeleteRequest).Methods("GET")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router))
}

func main() {
	fmt.Print("usao u main /n")
	database := initDB()
	repo := initRepo(database)
	service := initServices(repo)
	handler := initHandler(service)
	handleFunc(handler)
	r := mux.NewRouter()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), r))
}
