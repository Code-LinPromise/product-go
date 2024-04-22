package database

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	VideoAddr string `json:"video_addr"`
	NowVideo  bool   `json:"now_video"`
}

func MigrateVideo() *Video {
	return &Video{}
}
