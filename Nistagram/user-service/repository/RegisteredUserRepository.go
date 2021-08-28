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

//func (repo *RegisteredUserRepository) CreateAccount(account *model.Account) int64 {
//	result := repo.Database.Create(account)
//	fmt.Println("trebalo je da ga kreira")
//	fmt.Println(result)
//	if result.RowsAffected == 0 {
//		return result.RowsAffected
//	}
//	return result.RowsAffected
//}

func (repo *RegisteredUserRepository) GetRegisteredUserByID(id uint) (*model.RegisteredUser, error) {
	registeredUser := &model.RegisteredUser{}
	if err := repo.Database.Preload("Account").First(&registeredUser, "ID = ?", id).Error; err != nil {
		return nil, err
	}
	return registeredUser, nil
}
func (repo *RegisteredUserRepository) GetRegisteredUserByID1(id uint) (*model.RegisteredUser, error) {
	registeredUser := &model.RegisteredUser{}
	if err := repo.Database.Preload("Account").Preload("RelatedUsers").First(&registeredUser, "ID = ?", id).Error; err != nil {
		return nil, err
	}
	return registeredUser, nil
}



func (repo *RegisteredUserRepository) GetRegisteredUserByUsername(username string) (*model.Account, error) {
	registeredUser := &model.Account{}
	repo.Database.Model(&registeredUser)
	repo.Database.First(&registeredUser,"username = ?" , username)
	return registeredUser, nil
}

func (repo *RegisteredUserRepository) UpdateRegisterUser(registeredUser *model.RegisteredUser) error {
	err := repo.Database.Save(registeredUser.Account).Error
	fmt.Println("pokusaj dodavanja pratioca")
	if err == nil {
		err1 := repo.Database.Save(registeredUser).Error
		fmt.Println("u funkciji ")

		if err1 != nil {
			return err
		}
		return nil
	}
	return nil
}
/*
func (repo *RegisteredUserRepository) AddFollower(idRegisteUser string,idRelatedUser string) (bool, error) {
	registerUser := &model.RegisteredUser{}
	relatedUser := &model.RegisteredUser{}

	repo.Database.First(&registerUser, "id = ?", idRegisteUser)
	RelatedUserRepository.Database.First(&relatedUser,"id = ?",idRelatedUser)

	fmt.Println(registerUser.Account.Name + "/n+++++++")
	fmt.Println(relatedUser.Account.Name+"/n-----------")

//	registerUser.RelatedUsers = append(registerUser.RelatedUsers, *relatedUser)

	err := repo.Database.Save(registerUser).Error
	if err==nil{
		return true,err
	}
	return false,nil
}
*/

func (repo *RegisteredUserRepository) FindAccountByUsername(username string) (bool, error) {
	account := &model.Account{}
	if err := repo.Database.First(&account, "username = ?", username).Error; err == nil {
		fmt.Println("Ispis greske u accountu")
		fmt.Println(err)
		fmt.Println(account.Name)
		return false, err
	}
	return true, nil
}

func (repo *RegisteredUserRepository) FindAccountByEmail(email string) (bool, error) {
	fmt.Println("Ulazak u metodu za mail")
	account := &model.Account{}
	if err := repo.Database.First(&account, "email = ?", email).Error; err == nil {
		fmt.Println("Ispis greske u accountu")
		fmt.Println(err)
		fmt.Println(account.Name)
		return false, err
	}
	return true, nil
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