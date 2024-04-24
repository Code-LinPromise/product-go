package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"product.com/m/database"
	"product.com/m/tool"
)

func CGetLikedController(c *gin.Context) {
	claims := c.MustGet("token").(*tool.Claims)
	if claims == nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "用户凭证不正确"})
	}
	user_phone := claims.Username
	var user database.User
	database.DB.Where("phone=?", user_phone).Take(&user)
	userId := user.ID
	var liked_list []database.Liked
	var product_list []database.Product
	database.DB.Where("user_id=?", userId).Find(&liked_list)
	for _, value := range liked_list {
		var liked_product database.Product
		database.DB.Where("id=?", value.ProductId).Find(&liked_product)
		product_list = append(product_list, liked_product)
	}
	c.JSON(http.StatusOK, gin.H{"msg": "请求成功", "liked_product_list": product_list})
}
