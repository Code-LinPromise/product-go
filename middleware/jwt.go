package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"product.com/m/tool"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusNotFound, gin.H{"msg": "token is null"})
		} else {
			claims, err := tool.ParseToken(token)
			if err != nil || time.Now().Unix() > claims.ExpiresAt {
				c.JSON(http.StatusBadRequest, gin.H{"msg": "token is bad"})
				c.Abort()
				return
			} else {
				c.Set("token", claims)
			}
		}

		c.Next()
	}
}
