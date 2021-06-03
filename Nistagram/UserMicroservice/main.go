package AdminRequestMicroservice

import (
	"XML/Nistagram/UserMicroservice/Handler"
	"XML/Nistagram/UserMicroservice/Model1"
	"XML/Nistagram/UserMicroservice/Repository"
	"XML/Nistagram/UserMicroservice/Service"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	. "gorm.io/gorm"

)

func initDB() *DB {
	database, err := Open(sqlite.Open("consumers.db"), &Config{})
	if err != nil {
		log.Fatal(err)
	}
	database.AutoMigrate(&Model1.Account{})

	/*Loading test data*/
	consumers := []Model1.Account{
		{Email: "petar.petrovic@mail.cc", Password: "petar", Name: "petar", Surname: "petrovic"},
		{Email: "ivan.ivanovic@example.cc", Password: "ivan", Name: "ivan", Surname: "ivanovic"},
	}
	for _, consumer := range consumers {
		database.Create(&consumer)
	}
	return database
}

func initRepo(database *DB) *Repository.AccountRepository {
	return &Repository.AccountRepository{Database: database}
}

func initServices(repo *Repository.AccountRepository) *Service.UserService {
	return &Service.UserService{Repo: repo}
}

func initHandler(service *Service.UserService) *Handler.AccountHandler {
	return &Handler.AccountHandler{Service: service}
}
func handleFunc(handler *Handler.AccountHandler) {
	router := mux.NewRouter().StrictSlash(true)

	//router.HandleFunc("/", handler.Hello).Methods("GET")
	router.HandleFunc("/", handler.CreateAccount).Methods("POST")
	//router.HandleFunc("/verify/{consumerId}", handler.Verify).Methods("GET")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router))
}

func main() {
	database := initDB()
	repo := initRepo(database)
	service := initServices(repo)
	handler := initHandler(service)
	handleFunc(handler)
}
