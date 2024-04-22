package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"product.com/m/database"
)

func CGetProductController(c *gin.Context) {
	var body struct {
		ProductKind string `json:"product_kind"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	var ProductList []database.Product
	fmt.Println(body)
	queryProduct := database.Product{
		ProductKind: body.ProductKind,
	}
	database.DB.Debug().Where(&queryProduct).Find(&ProductList)
	for i, value := range ProductList {
		var ProductDetail []database.ProdctDetailsImage
		database.DB.Where("product_id=?", value.ID).Find(&ProductDetail)
		ProductList[i].ProdctDetails = ProductDetail
	}
	c.JSON(http.StatusOK, gin.H{"msg": "请求成功", "data": ProductList})
}
