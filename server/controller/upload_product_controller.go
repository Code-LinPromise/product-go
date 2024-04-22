package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"product.com/m/database"
	"product.com/m/tool"
)

func UploadProductController(c *gin.Context) {
	var body struct {
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
		var ProductDetailsImages []database.ProdctDetailsImage
		for _, value := range body.ProductDetails {
			ProductDetailsImages = append(ProductDetailsImages, database.ProdctDetailsImage{
				Images: value,
			})
		}
		database.DB.Create(&database.Product{
			Name:          body.Name,
			ProductKind:   body.Kind,
			Price:         body.Price,
			ProductCover:  body.ProductCover,
			ProdctDetails: ProductDetailsImages,
		})
		c.JSON(http.StatusOK, gin.H{"msg": "上传成功"})
	}
}
