package repository

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"user-service/model"
)

type RegisteredUserRepository struct {
	Database *gorm.DB
}

func New() (*RegisteredUserRepository, error) {
	ts := &RegisteredUserRepository{}

	host := os.Getenv("DBHOST")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")
	dbport := os.Getenv("DBPORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, user, password, dbname, dbport)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	ts.Database = db
	ts.Database.AutoMigrate(&model.RegisteredUser{}, &model.Account{})

	return ts, nil
}

func (repo *RegisteredUserRepository) CreateRegisteredUser(registeredUser *model.RegisteredUser) error {
	result := repo.Database.Create(registeredUser)
	//TODO convert to logs
	fmt.Println(result.RowsAffected)
	return nil
}

//func (repo *ConsumerRepository) ConsumerExists(consumerId uuid.UUID) bool {
//	var count int64
//	repo.Database.Where("id = ?", consumerId).Find(&model.Account{}).Count(&count)
//	return count != 0
//}

func (ts *RegisteredUserRepository) Close() error {
	db, err := ts.Database.DB()
	if err != nil {
		return err
	}

	db.Close()
	return nil
}