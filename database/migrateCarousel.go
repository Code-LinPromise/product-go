package database

import "gorm.io/gorm"

type Carousel struct {
	gorm.Model
	Serial      uint32 `json:"serial"`
	ImgAddr     string `json:"img_addr"`
	NowCarousel bool   `json:"now_carousel"`
}

func MigrateCarousel() *Carousel {
	return &Carousel{}
}
