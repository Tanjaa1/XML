package repository

import (
	"fmt"
	"gorm.io/gorm"
	"user-service/model"
)

type RelatedUserRepository struct {
	Database *gorm.DB
}

func (repo *RelatedUserRepository) CreateRelatedUser(relatedUser *model.RelatedUser) error {

	result := repo.Database.Create(relatedUser)
	//TODO convert to logs
	fmt.Println(result.RowsAffected)
	fmt.Println( "POKUSAJ UISA RELATED USERA")
	return nil
}

func (repo *RelatedUserRepository) FollowerExists(username string,registerUserId int) bool {
	fmt.Println("repozitorijum")
	/*if err := repo.Database.First("username = ? and registered_user_id = ?", username,registerUserId).Find(&model.RelatedUser{}).Error; err == nil {
		fmt.Println("IPronadjen isti korisnik")
		fmt.Println(err)
		return false
	}
*/
	return true
}
