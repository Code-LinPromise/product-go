package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"product.com/m/tool"
)

func TokenValidateController(c *gin.Context) {
	claims := c.MustGet("token").(*tool.Claims)
	if claims == nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "用户凭证不正确"})
	}
	c.JSON(http.StatusOK, gin.H{"msg": "请求成功", "userInfo": claims.Username})
}
