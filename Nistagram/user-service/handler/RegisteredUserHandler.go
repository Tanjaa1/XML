package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"user-service/dto"
	"user-service/service"
)

type RegisteredUserHandler struct {
	Service *service.RegisteredUserService
}

//func (handler *ConsumerHandler) Hello(w http.ResponseWriter, r *http.Request) {
//	addrs, _ := net.InterfaceAddrs()
//	for i, addr := range addrs {
//		fmt.Printf("%d %v\n", i, addr)
//	}
//}

func (handler *RegisteredUserHandler) CreateRegisteredUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("creating")
	var registeredUser dto.RequestRegisteredUser
	fmt.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(&registeredUser)
	fmt.Println(err)
	if err != nil {
		//TODO log
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println("prosao decoder")
	fmt.Println(registeredUser.Account)
	err = handler.Service.CreateRegisteredUser(&registeredUser)
	if err != nil {
		fmt.Println(err)
 	w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *RegisteredUserHandler) GetMyPersonalData(w http.ResponseWriter, r *http.Request) {
	//data, err := handler.Service.GetMyPersonalData(util.GetLoggedUserIDFromToken(r))
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("Ispisuje se id")
	fmt.Println(id)
	id2,err := strconv.ParseUint(id, 10, 64)
	if err != nil{
		fmt.Println(err)
	}
	id3 := uint(id2)
	data, err := handler.Service.GetMyPersonalData(id3)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (handler *RegisteredUserHandler) ChangePersonalData(w http.ResponseWriter, r *http.Request) {
	var dto dto.MyProfileDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//userId := util.GetLoggedUserIDFromToken(r)
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("Ispisuje se id")
	fmt.Println(id)
	id2,err := strconv.ParseUint(id, 10, 64)
	if err != nil{
		fmt.Println(err)
	}
	id3 := uint(id2)
	err = handler.Service.ChangePersonalData(dto, id3)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"success\":\"ok\"}"))
	w.Header().Set("Content-Type", "application/json")
}

//func (ts *postServer) CreateRegisteredUser(w http.ResponseWriter, req *http.Request) {
//	span := tracer.StartSpanFromRequest("cretePostHandler", ts.tracer, req)
//	defer span.Finish()
//
//	span.LogFields(
//		tracer.LogString("handler", fmt.Sprintf("handling post create at %s\n", req.URL.Path)),
//	)
//
//	// Enforce a JSON Content-Type.
//	contentType := req.Header.Get("Content-Type")
//	mediatype, _, err := mime.ParseMediaType(contentType)
//	if err != nil {
//		tracer.LogError(span, err)
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//	if mediatype != "application/json" {
//		http.Error(w, "expect application/json Content-Type", http.StatusUnsupportedMediaType)
//		return
//	}
//
//	ctx := tracer.ContextWithSpan(context.Background(), span)
//	rt, err := decodeBodyRegisteredUser(ctx, req.Body)
//	if err != nil {
//		tracer.LogError(span, err)
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	id := ts.store.CreateRegisteredUser(ctx, rt)
//	renderJSON(ctx, w, ResponseId{Id: id})
//}

//func (handler *ConsumerHandler) Verify(w http.ResponseWriter, r *http.Request) {
//	fmt.Println("verifying")
//	vars := mux.Vars(r)
//	id := vars["consumerId"]
//	if id == "" {
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//	exists, err := handler.Service.UserExists(id)
//	if err != nil {
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//	if exists {
//		w.WriteHeader(http.StatusOK)
//	} else {
//		w.WriteHeader(http.StatusNotFound)
//	}
//}

