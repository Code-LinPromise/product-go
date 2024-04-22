package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"product.com/m/database"
	"product.com/m/tool"
)

func UploadPhoneController(c *gin.Context) {
	var body struct {
		PhoneContent string `json:"phone_content"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		claims := c.MustGet("token").(*tool.Claims)
		if claims == nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "用户凭证不正确"})
		}
		var havePhone database.Phone
		PhoneContent := database.Phone{
			PhoneContent: body.PhoneContent,
			NowPhone:     true,
		}
		tx := database.DB.Begin()
		tx.Find(&havePhone, "now_phone", 1).Update("now_phone", 0)
		tx.Create(&PhoneContent)
		if err := tx.Commit().Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "上传失败"})
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{"msg": "上传成功"})
	}
}
