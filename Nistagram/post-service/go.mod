module post-service

replace post-service => ./

go 1.15

require (
	github.com/google/uuid v1.2.0
	github.com/gorilla/mux v1.8.0
	github.com/mattn/go-sqlite3 v1.14.7 // indirect
	gorm.io/driver/mysql v1.1.1 // indirect
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.21.9
)
