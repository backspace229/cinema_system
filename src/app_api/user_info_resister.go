package app_api

import (
	"net/http"
	"src/logic"

	//"google.golang.org/appengine"
)

func UserInfoResister(w http.ResponseWriter, r *http.Request) {

	//c := appengine.NewContext(r)

	//リクエストからパラメータを取得する

	logic.UserInfoPut()
}
