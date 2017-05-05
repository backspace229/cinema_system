package model

import "time"

const (
	USER_INFO_MODEL = "UserInfo"
)

type UserInfo struct {
	UserID       string    `json:"user_id"`
	Password     string    `json:"password"`
	ResisterDate time.Time `json:"resister_date"`
	UpdateDate   time.Time `json:"update_date"`
	DeleteFlag   bool      `json:"delete_flag"`
}
