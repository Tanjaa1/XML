package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"campaign-service/handler"
	"campaign-service/model"
	"campaign-service/repository"
	"campaign-service/service"
)

func initDB() *gorm.DB {
	database, err := gorm.Open(sqlite.Open("consumers.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	database.AutoMigrate(&model.Campaign{})

	/*Loading test data*/
	consumers := []model.Campaign{
		{Repetition: 1},
		{Repetition: 4},
	}
	for _, consumer := range consumers {
		database.Create(&consumer)
	}
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

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router))
}

func main() {
	//database := initDB()
	//repo := initRepo(database)
	//service := initServices(repo)
	//handler := initHandler(service)
	//handleFunc(handler)
}
