package service

import (
	"github.com/google/uuid"
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
func (service *ConsumerService) CreateRequest(request *model.Request) error {
	service.Repo.CreateRequest(request)
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