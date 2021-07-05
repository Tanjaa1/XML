module user-service

replace user-service => ./

go 1.15

require (
	github.com/HdrHistogram/hdrhistogram-go v1.1.0 // indirect
	github.com/google/uuid v1.2.0
	github.com/gorilla/handlers v1.5.1 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/lib/pq v1.10.2 // indirect
	github.com/mattn/go-sqlite3 v1.14.7 // indirect
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pkg/errors v0.9.1 // indirect
	github.com/uber/jaeger-client-go v2.29.1+incompatible
	github.com/uber/jaeger-lib v2.4.1+incompatible
	go.uber.org/atomic v1.7.0 // indirect
	gorm.io/driver/mysql v1.1.0
	gorm.io/driver/sqlite v1.1.4 // indirect
	gorm.io/gorm v1.21.9
)
