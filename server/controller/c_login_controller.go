package controller

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm/clause"
	"product.com/m/database"
	"product.com/m/redisModule"
	"product.com/m/tool"
)

func CLoginController(c *gin.Context) {
	var body struct {
		Code string `json:"code"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	redis_access_token, errCode := redisModule.Redis.Get(context.Background(), "access_token").Result()
	if (errCode == redis.Nil || errCode != nil) || redis_access_token == "" {
		access_token := tool.GetAccessToken()
		errCode := redisModule.Redis.Set(context.Background(), "access_token", access_token, 120*time.Minute).Err()
		if errCode != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "登录失败"})
			return
		}
	}
	getPhoneNumberUrl := fmt.Sprintf("https://api.weixin.qq.com/wxa/business/getuserphonenumber?access_token=%s", redis_access_token)
	phone_info := tool.GetCPhone(body.Code, getPhoneNumberUrl)
	fmt.Println("phone_info.PhoneInfo.PhoneNumber", phone_info.PhoneInfo.PhoneNumber)
	fmt.Println("phone_info", phone_info)
	if phone_info.PhoneInfo.PhoneNumber == "" {
		c.JSON(http.StatusOK, gin.H{"msg": "登陆失败", "phone_info": phone_info, "code": body.Code})
		return
	}
	user := database.User{
		Phone: phone_info.PhoneInfo.PhoneNumber,
	}
	database.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "phone"}},
		DoUpdates: clause.AssignmentColumns([]string{"updated_at"})}).Create(&user)
	token, err := tool.GenerateToken(phone_info.PhoneInfo.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": "登陆失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "登陆成功", "token": token})
}
