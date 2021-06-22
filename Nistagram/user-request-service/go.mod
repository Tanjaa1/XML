module user-request-service

replace user-request-service => ./

go 1.15

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/google/uuid v1.2.0
	github.com/gorilla/mux v1.8.0
	gorm.io/gorm v1.21.8
)
