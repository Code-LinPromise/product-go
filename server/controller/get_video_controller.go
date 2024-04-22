package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"product.com/m/database"
	"product.com/m/tool"
)

func GetVideoController(c *gin.Context) {
	claims := c.MustGet("token").(*tool.Claims)
	if claims == nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "用户凭证不正确"})
	}
	var videolList []database.Video
	database.DB.Where("now_video=?", 1).Find(&videolList)
	c.JSON(http.StatusOK, gin.H{"msg": "请求成功", "video_list": videolList})
}
