package app_api

import (
	"net/http"

	"io/ioutil"

	"encoding/json"
	"src/logic"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func UserInfoResister(w http.ResponseWriter, r *http.Request) {

	c := appengine.NewContext(r)

	//リクエストからパラメータを取得する
	bytestr, _ := ioutil.ReadAll(r.Body)
	body := string(bytestr)

	log.Infof(c, "[BODY] %v", body)

	var test TestModel

	err := json.Unmarshal(bytestr, &test)
	if err != nil {
		log.Infof(c, "[ERROR] %v", err)
	}

	log.Infof(c, "[UserID]%v, [Password]%v", test.UserID, test.Password)
	logic.UserInfoPut()
}

type TestModel struct {
	UserID   string `json:"user_id"`
	Password string `json:"password"`
}
