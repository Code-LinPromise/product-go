package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"product.com/m/database"
	"product.com/m/tool"
)

func GetPhoneController(c *gin.Context) {
	claims := c.MustGet("token").(*tool.Claims)
	if claims == nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "用户凭证不正确"})
	}
	var phone_content database.Phone
	database.DB.Where("now_phone=?", 1).Take(&phone_content)
	c.JSON(http.StatusOK, gin.H{"msg": "请求成功", "phone_content": phone_content})
}
