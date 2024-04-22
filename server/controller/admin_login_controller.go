package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"product.com/m/database"
	"product.com/m/tool"
)

func AdminLoginController(c *gin.Context) {
	var body struct {
		Username string
		Password string
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		var findResult database.Admin
		username := body.Username
		password := body.Password
		//查询数据库是否已经存在用户
		database.DB.Where("user_name = ?", username).Find(&findResult)
		fmt.Println(findResult)
		fmt.Println(body)
		if findResult.UserName == "" {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "用户不存在"})
		} else {
			if findResult.UserName == username && findResult.Password == password {
				token, err := tool.GenerateToken(username)
				if err != nil {
					panic(err)
				}
				c.JSON(http.StatusOK, gin.H{"msg": "登录成功", "token": token})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"msg": "密码错误"})
			}
		}

	}
}
