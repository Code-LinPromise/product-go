package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"product.com/m/database"
	"product.com/m/tool"
)

func UploadVideoController(c *gin.Context) {
	var body struct {
		VideoAddr []string `json:"video_addr"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		claims := c.MustGet("token").(*tool.Claims)
		if claims == nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "用户凭证不正确"})
		}
		var videoList []database.Video
		for _, value := range body.VideoAddr {
			videoList = append(videoList, database.Video{
				VideoAddr: value,
				NowVideo:  true,
			})
		}
		var nowVideo []database.Video
		tx := database.DB.Begin()
		tx.Find(&nowVideo, "now_video", 1).Update("now_video", 0)
		tx.Create(&videoList)
		if err := tx.Commit().Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "上传失败"})
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{"msg": "上传成功"})
	}
}
