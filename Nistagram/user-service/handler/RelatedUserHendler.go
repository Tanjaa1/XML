package handler

import (
	"user-service/service"
)

type RelatedUserHendler struct {
	Service *service.RelatedUserService
}


/*
func (handler *RelatedUserHendler) FolowerExist(w http.ResponseWriter, r *http.Request) {
	fmt.Println("usla uu hendler")
	vars := mux.Vars(r)
	id := vars["registerUserId"]
	username := vars["username"]

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	exists, err := handler.Service.UserExists(username,id)
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
*/