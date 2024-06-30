package server

import (
	"github.com/gin-gonic/gin"
	"product.com/m/middleware"
	"product.com/m/server/controller"
	"product.com/m/tool"
)

func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	tool.InitCors(r)
	authApi := r.Group("/api/v1/auth")
	authApi.Use(middleware.JWT())
	{
		authApi.POST("/upload_product", controller.UploadProductController)
		authApi.POST("/get_product", controller.GetProductController)
		authApi.DELETE("/delete_product", controller.DeleteProductController)
		authApi.POST("/token_validate", controller.TokenValidateController)
		authApi.POST("/upload_carousel", controller.UploadCarouselController)
		authApi.GET("/get_carousel", controller.GetCarouselController)
		authApi.POST("/upload_video", controller.UploadVideoController)
		authApi.GET("/get_video", controller.GetVideoController)
		authApi.POST("/upload_phone", controller.UploadPhoneController)
		authApi.GET("/get_phone", controller.GetPhoneController)
		authApi.POST("/upload_kindimg", controller.UploadKindImgController)
		authApi.GET("/get_kindimg", controller.GetKindImgController)
		authApi.POST("/get_id_product", controller.GetProductIdController)
		authApi.POST("update_product", controller.UpdateProductController)
	}
	authCApi := r.Group("/api/v1/c/auth")
	authCApi.Use(middleware.JWT())
	{
		authApi.POST("/upload_liked", controller.CUploadLikedController)
		authApi.GET("/get_liked", controller.CGetLikedController)
	}
	r.POST("api/v1/admin_login", controller.AdminLoginController)
	r.GET("api/v1/c/get_kind_info", controller.CGetKindImgController)
	r.POST("api/v1/c/get_product", controller.CGetProductController)
	r.GET("api/v1/c/get_carousel", controller.CGetCarouselController)
	r.GET("api/v1/c/get_video", controller.CGetVideoController)
	r.GET("api/v1/c/get_phone", controller.CGetPhoneController)
	r.POST("api/v1/c/search_product", controller.CSearchProductController)
	r.POST("api/v1/c/login", controller.CLoginController)
	r.GET("/api/v1/c/get_liked_product", controller.CGetLikedController)
	r.POST("/api/v1/c/upload_liked_product", controller.CUploadLikedController)
	return r
}
