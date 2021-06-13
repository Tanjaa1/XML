package service

import (
	"fmt"
	"time"
	"user-service/dto"
	"user-service/model"
	"user-service/repository"
)

type RegisteredUserService struct {
	Repo *repository.RegisteredUserRepository
}

func (service *RegisteredUserService) CreateRegisteredUser(dto *dto.RequestRegisteredUser) error {
	fmt.Println("DTO NAME:")
	fmt.Println(dto.Account.Name)
	account := model.Account{Name: dto.Account.Name, Surname: dto.Account.Surname, DateOfBirth: time.Now(),
			Email: dto.Account.Email,Username: dto.Account.Username, Password: dto.Account.Password, Gender: model.ConvertGender(dto.Account.Gender),
	 		PhoneNumber: dto.Account.PhoneNumber}
	fmt.Println("kreiran akaunt")
	fmt.Println(account)
	//service.Repo.CreateAccount(&account)
	registeredUser := model.RegisteredUser{Account: account,Description: dto.Description, Website: dto.Website, IsVerified: dto.IsVerified, IsPrivate: dto.IsPrivate,
		AcceptingMessage: dto.AcceptingMessage, AcceptingTag: dto.AcceptingTag,/* UserType: model.NONE*//*, FollowingRequestIdList: dto.FollowingRequestIdList,*/
		/*RelatedUsers: dto.RelatedUsers*//*[]model.RelatedUser{},*/ /*CollectionsIdList: dto.CollectionsIdList, CooperationRequestIdList: dto.CooperationRequestIdList,
		MessageRequestIdList: dto.MessageRequestIdList, HighlightsIdList: dto.HighlightsIdList*/}

	service.Repo.CreateRegisteredUser(&registeredUser)
	return nil
}

//func (service *ConsumerService) UserExists(consumerId string) (bool, error) {
//	id, err := uuid.Parse(consumerId)
//	if err != nil {
//		print(err)
//		return false, err
//	}
//	exists := service.Repo.ConsumerExists(id)
//	return exists, nil
//}
