package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net"
	"net/http"
	"strconv"
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
	vars := mux.Vars(r)
	id1 := vars["idFollowing"]
	if id1 == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id2 := vars["idFollower"]
	if id2 == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var request model.Request
	var idd1, error =strconv.Atoi(id1)
	if error !=nil{
		fmt.Println("greska 1")
	}
	var idd2, error2 =strconv.Atoi(id2)
	if error2 !=nil{
		fmt.Println("greska 2")
	}
	request.SenderId=idd1
	request.ReceiverId=idd2
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		//TODO log
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(request)
	err = handler.Service.CreateRequest(&request)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
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


