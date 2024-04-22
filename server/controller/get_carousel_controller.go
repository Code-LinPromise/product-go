package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"product.com/m/database"
	"product.com/m/tool"
)

func GetCarouselController(c *gin.Context) {
	claims := c.MustGet("token").(*tool.Claims)
	if claims == nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "用户凭证不正确"})
	}
	var carouselList []database.Carousel
	database.DB.Where("now_carousel=?", 1).Find(&carouselList)
	c.JSON(http.StatusOK, gin.H{"msg": "请求成功", "carousel_list": carouselList})
}
