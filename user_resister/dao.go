package user_resister

import (
	"time"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

func dao(c context.Context, userID string, password string) (*datastore.Key, error) {
	log.Infof(c, "dao")

	key := datastore.NewIncompleteKey(c, USER_INFO_MODEL, nil)

	model :=  UserInfo{
		UserID:       userID,
		Password:     password,
		ResisterDate: time.Now(),
		UpdateDate:   time.Now(),
		DeleteFlag:   false,
	}

	newKey, err := datastore.Put(c, key, &model)
	if err != nil {
		log.Errorf(c, "ERROR=%v", err)
	}
	return newKey, err
}
