package src

import (
	"net/http"
	"src/app_api"
)

func init() {
	http.HandleFunc("/hello", app_api.Hello)
}
