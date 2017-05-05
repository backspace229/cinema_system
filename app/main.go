package app

import (
	"net/http"
	"src"
)

func init() {
	http.HandleFunc("/hello", src.Hello)
}
