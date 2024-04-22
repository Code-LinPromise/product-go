package database

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	UserName string `json:"username"`
	Password string `json:"password"`
}

func MigrateAdminDB() *Admin {
	return &Admin{}
}
