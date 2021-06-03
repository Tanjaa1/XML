package Repository

import (
	"XML/Nistagram/UserMicroservice/Model1"
	"fmt"
	"gorm.io/gorm"
)

type AccountRepository struct {
	Database *gorm.DB
}

func (repo *AccountRepository) CreateAccount(user *Model1.Account) error {
	result := repo.Database.Create(user)
	//TODO convert to logs
	fmt.Println(result.RowsAffected)
	return nil
}
