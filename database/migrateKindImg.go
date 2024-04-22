package database

import "gorm.io/gorm"

type KindImg struct {
	gorm.Model
	KindName   string `json:"name"`
	ImageAddr  string `json:"pic_url"`
	NowKindImg bool   `json:"now_kind_img"`
}

func MigrateKindImage() *KindImg {
	return &KindImg{}
}
