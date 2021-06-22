package repository

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"user-service/model"
)

type ConsumerRepository struct {
	Database *gorm.DB
}

func (repo *ConsumerRepository) CreateConsumer(consumer *model.Account) error {
	result := repo.Database.Create(consumer)
	//TODO convert to logs
	fmt.Println(result.RowsAffected)
	return nil
}

func (repo *ConsumerRepository) ConsumerExists(consumerId uuid.UUID) bool {
	var count int64
	repo.Database.Where("id = ?", consumerId).Find(&model.Account{}).Count(&count)
	return count != 0
}
