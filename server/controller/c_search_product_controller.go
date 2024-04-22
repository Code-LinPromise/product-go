package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"product.com/m/database"
)

func CSearchProductController(c *gin.Context) {
	var body struct {
		ProductName string `json:"product_name"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	var ProductList []database.Product
	database.DB.Debug().Where(fmt.Sprintf(" name like %q ", ("%" + body.ProductName + "%"))).Find(&ProductList)
	for i, value := range ProductList {
		var ProductDetail []database.ProdctDetailsImage
		database.DB.Where("product_id=?", value.ID).Find(&ProductDetail)
		ProductList[i].ProdctDetails = ProductDetail
	}
	c.JSON(http.StatusOK, gin.H{"msg": "请求成功", "data": ProductList})
}
