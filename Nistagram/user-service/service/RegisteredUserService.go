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

var mySigningKey = []byte("superTajnaLozinka")

func(service *RegisteredUserService) GenerateJWT( username string, password string) (string,error){
	fmt.Println("12345")
	account,err := service.Repo.GetRegisteredUserByUsername(username)

	if(account != nil) {
		fmt.Println(account.Account.Password)
		if(account.Account.Password == password) {
			token := jwt.New(jwt.SigningMethodES256)
			claims := token.Claims.(jwt.MapClaims)
			claims["authorized"] = true
			claims["role"] = "TODO"
			claims["exp"] = time.Now().Add(time.Minute * 45).Unix()

			tokenString, err := token.SignedString(mySigningKey)
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
	account := model.Account{Name: dto.Account.Name, Surname: dto.Account.Surname, DateOfBirth: time.Now(),
			Email: dto.Account.Email,Username: dto.Account.Username, Password: dto.Account.Password, Gender: model.ConvertGender(dto.Account.Gender),
	 		PhoneNumber: dto.Account.PhoneNumber}
	fmt.Println("kreiran akaunt")
	fmt.Println(account)
	//service.Repo.CreateAccount(&account)
	registeredUser := model.RegisteredUser{Account: account,Description: dto.Description, Website: dto.Website, IsVerified: dto.IsVerified, IsPrivate: dto.IsPrivate,
		AcceptingMessage: dto.AcceptingMessage, AcceptingTag: dto.AcceptingTag,/* UserType: model.NONE*//*, FollowingRequestIdList: dto.FollowingRequestIdList,*/
		/*RelatedUsers: dto.RelatedUsers*//*[]model.RelatedUser{},*/ /*CollectionsIdList: dto.CollectionsIdList, CooperationRequestIdList: dto.CooperationRequestIdList,
		MessageRequestIdList: dto.MessageRequestIdList, HighlightsIdList: dto.HighlightsIdList*/}

	service.Repo.CreateRegisteredUser(&registeredUser)
	return nil
}

func (service *RegisteredUserService) GetMyPersonalData(userId uint) (dto.MyProfileDTO, error) {
	registeredUser, err := service.Repo.GetRegisteredUserByID(userId)
	if err != nil {
		return dto.MyProfileDTO{}, err
	}
	account := registeredUser.Account
	ret := dto.MyProfileDTO{Username: account.Username, Name: account.Name, Surname: account.Surname,
		Email: account.Email, PhoneNumber: account.PhoneNumber, Gender: model.ConvertGenderToString(account.Gender),
		DateOfBirth: "treba uraditi"/*account.DateOfBirth*/, Description: registeredUser.Description,
		Website: registeredUser.Website, IsVerified: registeredUser.IsVerified, IsPrivate: registeredUser.IsPrivate,
		AcceptingMessage: registeredUser.AcceptingMessage, AcceptingTag: registeredUser.AcceptingTag,
		UserType: model.ConvertUserTypeToString(registeredUser.UserType)}
	return ret, nil
}

//func (service *ProfileService) ChangePersonalData(dto dto.PersonalDataDTO, loggedUserId uint) error {
//	profile, err := service.ProfileRepository.GetProfileByID(loggedUserId)
//	if err != nil {
//		return err
//	}
//	callAuth := bool(false)
//	if profile.Email != dto.Email {
//		callAuth = true
//	}
//	if profile.Username != dto.Username {
//		//TODO: change data in other ms
//		err = service.changeUsernameInPostService(loggedUserId, dto.Username)
//		if err != nil { return err }
//	}
//	profile.Username = dto.Username
//	profile.Website = dto.Website
//	profile.Biography = dto.Biography
//	profile.Email = dto.Email
//	profile.PersonalData.Name = dto.Name
//	profile.PersonalData.BirthDate = dto.BirthDate
//	profile.PersonalData.Gender = dto.Gender
//	profile.PersonalData.Surname = dto.Surname
//	profile.PersonalData.Telephone = dto.Telephone
//	err = service.ProfileRepository.UpdateProfile(profile)
//	if err != nil {
//		return err
//	}
//	if callAuth {
//		postBody, _ := json.Marshal(map[string]string{
//			"profileId": util.Uint2String(profile.ID),
//			"email":     profile.Email,
//		})
//		responseBody := bytes.NewBuffer(postBody)
//		authHost, authPort := util.GetAuthHostAndPort()
//		_, err = http.Post("http://"+authHost+":"+authPort+"/update-user", "application/json", responseBody)
//		if err != nil {
//			fmt.Println(err)
//			return err
//		}
//	}
//	err = service.ProfileRepository.UpdatePersonalData(profile.PersonalData)
//	return err
//}

//func (service *ConsumerService) UserExists(consumerId string) (bool, error) {
//	id, err := uuid.Parse(consumerId)
//	if err != nil {
//		print(err)
//		return false, err
//	}
//	exists := service.Repo.ConsumerExists(id)
//	return exists, nil
//}
