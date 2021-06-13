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
	//r := repo.CreateAccount(&registeredUser.Account)
	//r := repo.CreateAccount(registeredUser.Account)
	//fmt.Println("greska posle kreiranja accounta")
	//fmt.Println(r)
	fmt.Println("kreira se registrovani korisnik")
	//if r != 0{
		fmt.Println("Usaoooooo")
		result := repo.Database.Create(registeredUser)
		if result.RowsAffected == 0 {
			return fmt.Errorf("User not created")
		}else{
			fmt.Println("User created")
			return  nil
			}
	//}
	//return fmt.Errorf("Greska prilikom kreiranja accounta")
}

func (repo *RegisteredUserRepository) CreateAccount(account *model.Account) int64 {
	result := repo.Database.Create(account)
	fmt.Println("trebalo je da ga kreira")
	fmt.Println(result)
	//TODO convert to logs
	if result.RowsAffected == 0 {
		return result.RowsAffected
	}
	return result.RowsAffected
}

func (repo *RegisteredUserRepository) GetRegisteredUserByID(id uint) (*model.RegisteredUser, error) {
	registeredUser := &model.RegisteredUser{}
	if err := repo.Database.Preload("Account").First(&registeredUser, "ID = ?", id).Error; err != nil {
		return nil, err
	}
	return registeredUser, nil
}

func (repo *RegisteredUserRepository) UpdateRegisterUser(registeredUser *model.RegisteredUser) error {
	err := repo.Database.Save(registeredUser.Account).Error
	if err == nil {
		err1 := repo.Database.Save(registeredUser).Error
		if err1 != nil {
			return err
		}
		return nil
	}
	return nil
}

//func (repo *ConsumerRepository) ConsumerExists(consumerId uuid.UUID) bool {
//	var count int64
//	repo.Database.Where("id = ?", consumerId).Find(&model.Account{}).Count(&count)
//	return count != 0
//}

func (repo *RegisteredUserRepository) Close() error {
	db, err := repo.Database.DB()
	if err != nil {
		return err
	}

	db.Close()
	return nil
}