package app_api

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	log.Infof(c, "hello, world")
}
