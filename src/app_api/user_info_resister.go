package app_api

import (
	"net/http"
	"src/logic"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func UserInfoResister(w http.ResponseWriter, r *http.Request) {

	c := appengine.NewContext(r)

	//リクエストからパラメータを取得する
	//vars := mux.Vars(r)

	log.Infof(c, "[r.Body] %v", r.Body)

	logic.UserInfoPut()
}
