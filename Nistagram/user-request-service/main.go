package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"user-request-service/model"

	"gorm.io/gorm"

	"user-request-service/handler"

	"user-request-service/repository"
	"user-request-service/service"
)

var mySigningKey = []byte("mysupersecretkey")

func initDB() *gorm.DB {
	time.Sleep(time.Duration(20) *time.Second)
	dsn := "root:root@tcp(host.docker.internal:3306)/requestdb?parseTime=True&charset=utf8&autocommit=false"
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

func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Ispis token")
		fmt.Println(r.Header["Authorization"])
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
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Access-Control-Allow-Headers", "text/plain"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"http://localhost:8081"})
	credential := handlers.AllowCredentials()
	h := handlers.CORS(headers, methods, origins, credential)

	router.HandleFunc("/", handler.Hello).Methods("GET")
	router.HandleFunc("/", handler.CreateConsumer).Methods("POST")
	router.HandleFunc("/verify/{consumerId}", handler.Verify).Methods("GET")
	router.HandleFunc("/createRequest", IsAuthorized(handler.CreateRequest)).Methods("POST")
	router.HandleFunc("/getRequestById/{idRequest}", IsAuthorized(handler.GetRequestById)).Methods("GET")
	router.HandleFunc("/deleteRequest/{idRequest}", IsAuthorized(handler.DeleteRequest)).Methods("GET")
	router.HandleFunc("/deleteRequest2/{idSender}/{idReceiver}", IsAuthorized(handler.DeleteRequest2)).Methods("GET")
	router.HandleFunc("/isRequestExists/{myId}/{userId}", IsAuthorized(handler.IsRequestExists)).Methods("GET")
	router.HandleFunc("/getRequestsByReceiverId/{idRequest}", IsAuthorized(handler.GetRequestsByReceiverId)).Methods("GET")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), h(router)))
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
