package service

import (
	"github.com/google/uuid"
	"story-service/model"
	"story-service/repository"
)

type ConsumerService struct {
	Repo *repository.ConsumerRepository
}

func (service *ConsumerService) CreateConsumer(consumer *model.Story) error {
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
