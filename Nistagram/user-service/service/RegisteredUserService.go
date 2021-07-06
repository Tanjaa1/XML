package service

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
	"user-service/dto"
	"user-service/model"
	"user-service/repository"
)

type RegisteredUserService struct {
	Repo *repository.RegisteredUserRepository
}

var myKey = []byte("mysupersecretkey")


func(service *RegisteredUserService) GenerateJWT( username string, password string) (string,error){

	account,err := service.Repo.GetRegisteredUserByUsername(username)

	if(account != nil) {
		if(account.Password == password) {
			token := jwt.New(jwt.SigningMethodHS256)
			claims := token.Claims.(jwt.MapClaims)
			claims["authorized"] = true
			claims["role"] = "TODO"
			claims["exp"] = time.Now().Add(time.Minute * 45).Unix()

			tokenString, err := token.SignedString(myKey)
			if err != nil {
				fmt.Errorf("Greska u pravljenju tokena : %s", err.Error())
				return "", err
			}

			return tokenString, nil
		}
		return "PASSWORD", err
	}
	return "USERNAME", err
}

func (service *RegisteredUserService) CreateRegisteredUser(dto *dto.RequestRegisteredUser) error {
	fmt.Println("DTO NAME:")
	fmt.Println(dto.Account.Name)
	result, _ := service.FindAccountByUsername(dto.Account.Username)
	result1, _ := service.FindAccountByEmail(dto.Account.Email)
	//var err error
	if result == true && result1 == true {
		account := model.Account{Name: dto.Account.Name, Surname: dto.Account.Surname, DateOfBirth: time.Now(),
			Email: dto.Account.Email, Username: dto.Account.Username, Password: dto.Account.Password, Gender: model.ConvertGender(dto.Account.Gender),
			PhoneNumber: dto.Account.PhoneNumber}
		fmt.Println("kreiran akaunt")
		fmt.Println(account)
		//service.Repo.CreateAccount(&account)
		registeredUser := model.RegisteredUser{Account: account, Description: dto.Description, Website: dto.Website, IsVerified: dto.IsVerified, IsPrivate: dto.IsPrivate,
			AcceptingMessage: dto.AcceptingMessage, AcceptingTag: dto.AcceptingTag, /* UserType: model.NONE*//*, FollowingRequestIdList: dto.FollowingRequestIdList,*/
			/*RelatedUsers: dto.RelatedUsers*//*[]model.RelatedUser{},*/ /*CollectionsIdList: dto.CollectionsIdList, CooperationRequestIdList: dto.CooperationRequestIdList,
			MessageRequestIdList: dto.MessageRequestIdList, HighlightsIdList: dto.HighlightsIdList*/}

		service.Repo.CreateRegisteredUser(&registeredUser)
		return nil
	}
	return fmt.Errorf("Username and Email not unique")
}

func (service *RegisteredUserService) GetMyPersonalData(userId uint) (dto.MyProfileDTO, error) {
	registeredUser, err := service.Repo.GetRegisteredUserByID(userId)
	if err != nil {
		return dto.MyProfileDTO{}, err
	}
	fmt.Println("Pol korisnika:")
	account := registeredUser.Account
	fmt.Println(model.ConvertGenderToString(account.Gender))
	ret := dto.MyProfileDTO{Username: account.Username,Password: account.Password,Name: account.Name, Surname: account.Surname,
		Email: account.Email, PhoneNumber: account.PhoneNumber, Gender: model.ConvertGenderToString(account.Gender),
		DateOfBirth: "treba uraditi"/*account.DateOfBirth*/, Description: registeredUser.Description,
		Website: registeredUser.Website, IsVerified: registeredUser.IsVerified, IsPrivate: registeredUser.IsPrivate,
		AcceptingMessage: registeredUser.AcceptingMessage, AcceptingTag: registeredUser.AcceptingTag,
		UserType: model.ConvertUserTypeToString(registeredUser.UserType)}
	return ret, nil
}

func (service *RegisteredUserService) ChangePersonalData(dto dto.MyProfileDTO, userId uint) error {
		registeredUser, err := service.Repo.GetRegisteredUserByID(userId)
		if err != nil {
			return err
		}
		registeredUser.Account.Name = dto.Name
		fmt.Println(dto.Name)
		fmt.Println(registeredUser.Account.Name)
		registeredUser.Account.Surname = dto.Surname
		//registeredUser.Account.DateOfBirth = dto.Account.DateOfBirth
		registeredUser.Account.Email = dto.Email
		registeredUser.Account.Username = dto.Username
		registeredUser.Account.Password = dto.Password
		registeredUser.Account.Gender = model.ConvertGender(dto.Gender)
		registeredUser.Account.PhoneNumber = dto.PhoneNumber
		registeredUser.Description = dto.Description
		registeredUser.Website = dto.Website
		registeredUser.IsVerified = dto.IsVerified
		registeredUser.IsPrivate = dto.IsPrivate
		registeredUser.AcceptingMessage = dto.AcceptingMessage
		registeredUser.AcceptingTag = dto.AcceptingTag
		registeredUser.UserType = model.ConvertUserType(dto.UserType)
		err = service.Repo.UpdateRegisterUser(registeredUser)
	return err
}

func (service *RegisteredUserService) FindAccountByUsername(username string) (bool, error) {
	account, err := service.Repo.FindAccountByUsername(username)
	if err != nil {
		return false, err
	}
	if account == true {
		return true, nil
	}
	return false, nil
}

func (service *RegisteredUserService) FindAccountByEmail(email string) (bool, error) {
	account, err := service.Repo.FindAccountByEmail(email)
	if err != nil {
		return false, err
	}
	if account == true {
		return true, nil
	}
	return false, nil
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
