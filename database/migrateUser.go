package database

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Phone     string  `json:"phone"`
	UserLiked []Liked `json:"user_liked" gorm:"many2many:user_liked"`
}

func MigrateUser() *User {
	return &User{}
}
