package service

import (
	"fmt"
	"strconv"
	"user-service/model"
	"user-service/repository"
)

type RelatedUserService struct {
	Repo *repository.RelatedUserRepository
}

func (service *RelatedUserService) CreateRelatedUser(consumer *model.RelatedUser) error {
	service.Repo.CreateRelatedUser(consumer)
	return nil
}

func (service *RelatedUserService) UserExists(username string,registerUserId string) (bool, error) {
	intVar, err := strconv.Atoi(registerUserId)
	 fmt.Println("u servisu")
	fmt.Println(string(intVar) + "55555555555")
	if err != nil {
		print(err)
		return false, err
	}
	exists := service.Repo.FollowerExists(username,intVar)

	return exists, nil

}
