package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"product.com/m/database"
	"product.com/m/tool"
)

func UploadKindImgController(c *gin.Context) {
	var body struct {
		KindName string `json:"kind_name"`
		ImgAddr  string `json:"img_addr"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		claims := c.MustGet("token").(*tool.Claims)
		if claims == nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "用户凭证不正确"})
		}
		var kindImg database.KindImg
		kindImg.NowKindImg = true
		kindImg.KindName = body.KindName
		kindImg.ImageAddr = body.ImgAddr
		var nowkindImg []database.KindImg
		tx := database.DB.Begin()
		tx.Find(&nowkindImg, "kind_name", body.KindName).Update("now_kind_img", 0)
		tx.Create(&kindImg)
		if err := tx.Commit().Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "上传失败"})
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{"msg": "上传成功"})
	}
}
