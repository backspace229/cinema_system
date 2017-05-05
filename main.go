package src

import (
	"src/app_api"

	"github.com/gorilla/mux"
)

func init() {
	r := mux.NewRouter()
	version := "v1"

	r.HandleFunc(version+"/hello", app_api.Hello).Methods("GET")
	r.HandleFunc(version+"/user_resister", app_api.UserInfoResister).Methods("POST")
}
