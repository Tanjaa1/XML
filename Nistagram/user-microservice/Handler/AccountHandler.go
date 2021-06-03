package Handler

import (
	"XML/Nistagram/UserMicroservice/Model1"
	"XML/Nistagram/UserMicroservice/Service"
	"encoding/json"
	"fmt"
	"net/http"
)

type AccountHandler struct {
	Service *Service.UserService
}

func (handler *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Println("creating")
	var account Model1.Account
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		//TODO log
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(account)
	err = handler.Service.CreateAccount(&account)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
