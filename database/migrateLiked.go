package database

import "gorm.io/gorm"

type Liked struct {
	gorm.Model
	UserId    uint `json:"user_id"`
	ProductId uint `json:"product_id"`
}

func MigrateLiked() *Liked {
	return &Liked{}
}
