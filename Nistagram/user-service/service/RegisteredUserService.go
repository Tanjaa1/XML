package service

import (
	"user-service/dto"
	"user-service/model"
	"user-service/repository"
)

type RegisteredUserService struct {
	Repo *repository.RegisteredUserRepository
}

func (service *RegisteredUserService) CreateRegisteredUser(dto *dto.RequestRegisteredUser) error {
	registeredUser := model.RegisteredUser{Description: dto.Description, Website: dto.Website, IsVerified: dto.IsVerified, IsPrivate: dto.IsPrivate,
		AcceptingMessage: dto.AcceptingMessage, AcceptingTag: dto.AcceptingTag, UserType: model.NONE, FollowingRequestIdList: dto.FollowingRequestIdList,
		RelatedUsers: []model.RelatedUser{}, CollectionsIdList: dto.CollectionsIdList, CooperationRequestIdList: dto.CooperationRequestIdList,
		MessageRequestIdList: dto.MessageRequestIdList, HighlightsIdList: dto.HighlightsIdList}

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
