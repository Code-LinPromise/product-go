package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"product.com/m/database"
)

func CGetCarouselController(c *gin.Context) {
	var carouselList []database.Carousel
	err := database.DB.Where("now_carousel=?", 1).Find(&carouselList).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": "请求失败", "err": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "请求成功", "carousel_list": carouselList})
}
