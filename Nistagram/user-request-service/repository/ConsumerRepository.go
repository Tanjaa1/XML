package repository

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"user-request-service/model"
)

type ConsumerRepository struct {
	Database *gorm.DB
}

func (repo *ConsumerRepository) CreateConsumer(consumer *model.Request) error {
	result := repo.Database.Create(consumer)
	//TODO convert to logs
	fmt.Println(result.RowsAffected)
	return nil
}
func (repo *ConsumerRepository) CreateRequest(request *model.Request) error {
	result := repo.Database.Create(request)
	//TODO convert to logs
	fmt.Println(result.RowsAffected)
	fmt.Println("pokusaj kreiranja zahtjeva")
	return nil
}

func (repo *ConsumerRepository) ConsumerExists(consumerId uuid.UUID) bool {
	var count int64
	repo.Database.Where("id = ?", consumerId).Find(&model.Request{}).Count(&count)
	return count != 0
}

func (repo *ConsumerRepository) GetRequestById(id string) (*model.Request, error) {
	request := &model.Request{}
	if err := repo.Database.First(&request, "id = ?", id).Error; err == nil {
		fmt.Println("Pronalazim zahtjev")
		return request, err
	}
	return nil, nil
}

func (repo *ConsumerRepository) GetRequestsByReceiverId(id string) ([]model.Request, error) {
	var requests []model.Request
	result:=repo.Database.Table("requests").Find(&requests,"receiver_id like ?",id)
	return requests,result.Error
}

func (repo *ConsumerRepository) DeleteRequest(id string) (bool, error) {
	if err := repo.Database.Where("id like ?", id).Delete(&model.Request{}).Error; err == nil {
		fmt.Println("Brisem zahtjev")
		return true, err
	}
	return false, nil
}

func (repo *ConsumerRepository) DeleteRequest2(idSender string, idReceiver string) (bool, error) {
	if err := repo.Database.Where("sender_id like ? and receiver_id like ?", idSender, idReceiver).Delete(&model.Request{}).Error; err == nil {
		fmt.Println("Brisem zahtjev")
		return true, err
	}
	return false, nil
}

func (repo *ConsumerRepository) IsRequestExists(myId string, userId string) (*model.Request, error) {
	requests := &model.Request{}
	result:=repo.Database.Table("requests").First(&requests,"receiver_id like ? and sender_id like ?",userId, myId)
	return requests,result.Error
}