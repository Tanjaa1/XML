package service

import (
	"campaign-service/model"
	"campaign-service/repository"
	"github.com/google/uuid"
)

type ConsumerService struct {
	Repo *repository.ConsumerRepository
}

func (service *ConsumerService) CreateConsumer(consumer *model.Campaign) error {
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
