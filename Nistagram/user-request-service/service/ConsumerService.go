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

func (service *ConsumerService) UserExists(consumerId string) (bool, error) {
	id, err := uuid.Parse(consumerId)
	if err != nil {
		print(err)
		return false, err
	}
	exists := service.Repo.ConsumerExists(id)
	return exists, nil
}
