package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"product.com/m/database"
	"product.com/m/tool"
)

func GetKindImgController(c *gin.Context) {
	claims := c.MustGet("token").(*tool.Claims)
	if claims == nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "用户凭证不正确"})
	}
	var kindImg []database.KindImg
	database.DB.Where("now_kind_img=?", 1).Find(&kindImg)
	c.JSON(http.StatusOK, gin.H{"msg": "请求成功", "kind_img": kindImg})
}
