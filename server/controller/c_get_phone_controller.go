package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"product.com/m/database"
)

func CGetPhoneController(c *gin.Context) {
	var phone_content database.Phone
	database.DB.Where("now_phone=?", 1).Take(&phone_content)
	c.JSON(http.StatusOK, gin.H{"msg": "请求成功", "phone_content": phone_content.PhoneContent})
}
