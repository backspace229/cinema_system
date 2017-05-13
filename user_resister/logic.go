package user_resister

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/datastore"
)

func logic(c context.Context, userID string, password string) (*datastore.Key, error) {
	log.Infof(c, "logic")
	return dao(c, userID, password)
}