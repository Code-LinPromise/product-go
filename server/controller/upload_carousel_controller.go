package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"product.com/m/database"
	"product.com/m/tool"
)

func UploadCarouselController(c *gin.Context) {
	var body struct {
		ImgAddr []string `json:"img_addr"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		claims := c.MustGet("token").(*tool.Claims)
		if claims == nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "用户凭证不正确"})
		}
		var carousel []database.Carousel
		serial := uuid.New().ID()
		for _, value := range body.ImgAddr {
			carousel = append(carousel, database.Carousel{
				ImgAddr:     value,
				Serial:      serial,
				NowCarousel: true,
			})
		}
		var nowCarousel []database.Carousel
		tx := database.DB.Begin()
		tx.Find(&nowCarousel, "now_carousel", 1).Update("now_carousel", 0)
		tx.Create(&carousel)
		if err := tx.Commit().Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "上传失败"})
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{"msg": "上传成功"})
	}
}
