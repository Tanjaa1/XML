package service

import (
	"github.com/google/uuid"
	"user-request-service/dto"
	"user-request-service/model"
	"user-request-service/repository"
)

type ConsumerService struct {
	Repo *repository.ConsumerRepository
}

func (service *ConsumerService) CreateConsumer(consumer *model.Request) error {
	service.Repo.CreateConsumer(consumer)
	return nil
}
func (service *ConsumerService) CreateRequest(request *dto.RequestDTO) error {
	service.Repo.CreateRequest(&model.Request{SenderId: int(request.FollowerId), ReceiverId: int(request.FollowingId)})
	return nil
}
func (service *ConsumerService) UserExists(consumerId string) (bool, error) {
	id, err := uuid.Parse(consumerId)
	if err != nil {
		print(err)
		return false, err
	}
	exists := service.Repo.ConsumerExists(id)
	return exists, nil
}


func (service *ConsumerService) GetRequestById(id string) (*model.Request, error) {
	request, err := service.Repo.GetRequestById(id)
	return request, err
}
func (service *ConsumerService) DeleteRequest(id string) (bool, error) {
	request, err := service.Repo.DeleteRequest(id)
	return request, err
}

func (service *ConsumerService) DeleteRequest2(idSender string, idReceiver string) (bool, error) {
	request, err := service.Repo.DeleteRequest2(idSender, idReceiver)
	return request, err
}

func (service *ConsumerService) GetRequestByReciverId(id string) ([]dto.RequestDTO, error) {
	request, err := service.Repo.GetRequestsByReceiverId(id)
	var requestDTO []dto.RequestDTO
	for _,r := range request{
		requestDTO = append(requestDTO, dto.RequestDTO{Id: r.ID, FollowingId: uint(r.SenderId), FollowerId: uint(r.ReceiverId)})
	}
	return requestDTO, err
}

func (service *ConsumerService) IsRequestExists(myId string, userId string) (bool, error) {
	request, err := service.Repo.IsRequestExists(myId, userId)
	if request != nil{
		return true,err
	}
	return false, err
}