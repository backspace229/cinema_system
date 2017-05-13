package user_resister

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func init() {
	//ルータの生成
	r := mux.NewRouter()

	version := "/v1"

	//r.HandleFunc(version+"/hello", app_api.HelloHandler).Methods("GET")
	r.HandleFunc(version+"/user_resister", userInfoResister).Methods("POST")

	//Gorilla/muxを使用する上で必要なルーティング
	http.Handle("/", r)
}

func userInfoResister(w http.ResponseWriter, r *http.Request) {

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
	key, err := logic(c, test.UserID, test.Password)
	if err != nil {
		log.Errorf(c, "ERROR=%v", err)
	}

	log.Infof(c, "KEY=%v", key)
}

type TestModel struct {
	UserID   string
	Password string
}
