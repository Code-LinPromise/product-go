package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"product.com/m/database"
	"product.com/m/tool"
)

func CUploadLikedController(c *gin.Context) {
	var body struct {
		ProductId uint `json:"product_id"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		claims := c.MustGet("token").(*tool.Claims)
		if claims == nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "用户凭证不正确"})
		}
		user_phone := claims.Username
		var user database.User
		var product_liked database.Liked
		database.DB.Where("phone=?", user_phone).Take(&user)
		product_liked.ProductId = body.ProductId
		product_liked.UserId = user.ID
		database.DB.Create(&product_liked)
		c.JSON(http.StatusOK, gin.H{"msg": "收藏成功"})
	}
}
