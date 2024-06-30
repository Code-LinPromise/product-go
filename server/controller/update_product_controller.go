package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"product.com/m/database"
	"product.com/m/tool"
)

func UpdateProductController(c *gin.Context) {
	var body struct {
		Id             uint     `json:"id"`
		Name           string   `json:"name"`
		Kind           string   `json:"kind"`
		Price          uint     `json:"price"`
		ProductCover   string   `json:"product_cover"`
		ProductDetails []string `json:"product_details"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		claims := c.MustGet("token").(*tool.Claims)
		if claims == nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "用户凭证不正确"})
		}
		tx := database.DB.Begin()
		tx.Where("id = ?", body.Id).Delete(&database.Product{})
		var ProductDetailsImages []database.ProdctDetailsImage
		for _, value := range body.ProductDetails {
			ProductDetailsImages = append(ProductDetailsImages, database.ProdctDetailsImage{
				Images: value,
			})
		}
		tx.Create(&database.Product{
			Name:          body.Name,
			ProductKind:   body.Kind,
			Price:         body.Price,
			ProductCover:  body.ProductCover,
			ProdctDetails: ProductDetailsImages,
		})
		if err := tx.Commit().Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "更新失败"})
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{"msg": "更新成功"})
	}
}
