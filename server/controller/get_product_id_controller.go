package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"product.com/m/database"
	"product.com/m/tool"
)

func GetProductIdController(c *gin.Context) {
	claims := c.MustGet("token").(*tool.Claims)
	if claims == nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "用户凭证不正确"})
		return
	}
	var body struct {
		ProductId uint `json:"id"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var product database.Product
	var ProductDetail []database.ProdctDetailsImage
	database.DB.Where("id=?", body.ProductId).Take(&product)
	database.DB.Where("product_id=?", product.ID).Find(&ProductDetail)
	product.ProdctDetails = ProductDetail
	c.JSON(http.StatusOK, gin.H{"msg": "请求成功", "data": product})
}
