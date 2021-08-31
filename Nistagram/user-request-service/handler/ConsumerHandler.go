package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net"
	"net/http"
	_ "strconv"
	"user-request-service/dto"
	"user-request-service/model"
	"user-request-service/service"
)

type ConsumerHandler struct {
	Service *service.ConsumerService
}

func (handler *ConsumerHandler) Hello(w http.ResponseWriter, r *http.Request) {
	addrs, _ := net.InterfaceAddrs()
	for i, addr := range addrs {
		fmt.Printf("%d %v\n", i, addr)
	}
}

func (handler *ConsumerHandler) CreateConsumer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("creating")
	var consumer model.Request
	err := json.NewDecoder(r.Body).Decode(&consumer)
	if err != nil {
		//TODO log
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(consumer)
	err = handler.Service.CreateConsumer(&consumer)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
func (handler *ConsumerHandler) CreateRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("pravim zahtjev za pracenje /n")
	fmt.Println("creating")
	var registeredUser dto.RequestDTO
	fmt.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(&registeredUser)
	fmt.Println(err)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println("prosao decoder")
	err = handler.Service.CreateRequest(&registeredUser)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")


}


func (handler *ConsumerHandler) Verify(w http.ResponseWriter, r *http.Request) {
	fmt.Println("verifying")
	vars := mux.Vars(r)
	id := vars["consumerId"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	exists, err := handler.Service.UserExists(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if exists {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func (handler *ConsumerHandler) GetRequestById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	result, err := handler.Service.GetRequestById(vars["idRequest"])
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if result != nil {
		w.WriteHeader(http.StatusOK)
		fmt.Println("pronadjen /n")
	}else{
		w.WriteHeader(http.StatusOK)
		fmt.Println("nije pronadjen /n")

	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&result)
}

func (handler *ConsumerHandler) DeleteRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	result, err := handler.Service.DeleteRequest(vars["idRequest"])
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if result == true {
		w.WriteHeader(http.StatusOK)
		fmt.Println("obrisan /n")
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&result)
}

func (handler *ConsumerHandler) DeleteRequest2(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	result, err := handler.Service.DeleteRequest2(vars["idSender"], vars["idReceiver"])
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if result == true {
		w.WriteHeader(http.StatusOK)
		fmt.Println("obrisan /n")
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&result)
}

func (handler *ConsumerHandler) GetRequestsByReceiverId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	result, err := handler.Service.GetRequestByReciverId(vars["idRequest"])
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if result != nil {
		w.WriteHeader(http.StatusOK)
		fmt.Println("pronadjen /n")
	}else{
		w.WriteHeader(http.StatusOK)
		fmt.Println("nije pronadjen /n")

	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&result)
}

func (handler *ConsumerHandler) IsRequestExists(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	result, err := handler.Service.IsRequestExists(vars["myId"], vars["userId"])
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&result)
}


