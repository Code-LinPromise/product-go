package database

import "gorm.io/gorm"

type Phone struct {
	gorm.Model
	PhoneContent string `json:"phone_content"`
	NowPhone     bool   `json:"now_phone"`
}

func MigratePhone() *Phone {
	return &Phone{}
}
