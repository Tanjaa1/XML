package Service

import (
	"XML/Nistagram/UserMicroservice/Model1"
	"XML/Nistagram/UserMicroservice/Repository"
)

type UserService struct {
	Repo *Repository.AccountRepository
}

func (service *UserService) CreateAccount(account *Model1.Account) error {
	service.Repo.CreateAccount(account)
	return nil
}
