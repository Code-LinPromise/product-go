package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"product.com/m/database"
)

func CGetVideoController(c *gin.Context) {
	var videolList []database.Video
	err := database.DB.Where("now_video=?", 1).Find(&videolList).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": "请求失败", "err": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "请求成功", "video_list": videolList})
}
