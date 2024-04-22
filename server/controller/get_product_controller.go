package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"product.com/m/database"
	"product.com/m/tool"
)

func GetProductController(c *gin.Context) {
	claims := c.MustGet("token").(*tool.Claims)
	if claims == nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "用户凭证不正确"})
	}
	var body struct {
		Current   int    `json:"current"`
		PageSize  int    `json:"pageSize"`
		Name      string `json:"product_name"`
		StartTime string `json:"startTime"`
		EndTime   string `json:"endTime"`
		Kind      string `json:"product_kind"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	offset := (body.Current - 1) * body.PageSize
	var ProductList []database.Product

	queryProduct := database.Product{
		Name:        body.Name,
		ProductKind: body.Kind,
	}
	db := database.DB.Debug().Where(&queryProduct)

	// 执行分页查询
	db.Limit(body.PageSize).Offset(offset).Find(&ProductList)
	c.JSON(http.StatusOK, gin.H{"msg": "请求成功", "data": ProductList})
}
