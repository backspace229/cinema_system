package src

import (
	"net/http"
	"src/app_api"

	"github.com/gorilla/mux"
)

func init() {
	//ルータの生成
	r := mux.NewRouter()

	version := "/v1"

	r.HandleFunc(version+"/hello", app_api.HelloHandler).Methods("GET")
	r.HandleFunc(version+"/user_resister", app_api.UserInfoResister).Methods("POST")

	//Gorilla/muxを使用する上で必要なルーティング
	http.Handle("/", r)
}
