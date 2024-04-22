package tool

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitCors(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "PUT", "PATCH", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if origin == "http://127.0.0.1:8000" || origin == "http://localhost:8000" {
				return true
			} else {
				log.Printf("%v is now allowed", origin)
				return false
			}
		},
		MaxAge: 12 * time.Hour,
	}))
}
