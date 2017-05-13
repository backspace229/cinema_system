package user_resister

import "time"

const (
	USER_INFO_MODEL = "UserInfo"
)

type UserInfo struct {
	UserID       string
	Password     string
	ResisterDate time.Time
	UpdateDate   time.Time
	DeleteFlag   bool
}
