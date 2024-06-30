package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CLoginController(c *gin.Context) {
	var body struct {
		Code string `json:"code"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	appid := "wx303d3a2eed3d65e4"
	secret := "db8f84ec47ae749db7d5a10fb2aa2c8e"
	session_token_url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=%s", appid, secret, body.Code, "authorization_code")
	resp, err := http.Get(session_token_url)
	if err != nil {
		log.Fatalf("Error sending GET request: %v", err)
	}
	defer resp.Body.Close()
	// 检查响应状态码是否表示成功（通常为200 OK）
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected status code: %d", resp.StatusCode)
	}
	// 读取响应体
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}
	type ResponseData struct {
		OpenId     string `json:"openid"`
		SessionKey string `json:"session_key"`
		UnionidId  string `json:"unionid"`
		ErrCode    uint   `json:"errcode"`
		ErrMsg     string `json:"errmsg"`
	}
	var responseData ResponseData
	err = json.Unmarshal(bodyBytes, &responseData)
	if err != nil {
		log.Fatalf("Error unmarshaling response body: %v", err)
	}
	fmt.Println("response", responseData)
	c.JSON(http.StatusOK, gin.H{"msg": "登陆成功", "response": responseData})
	// var redis_access_token string
	// var errCode error
	// redis_access_token, errCode = redisModule.Redis.Get(context.Background(), "access_token").Result()
	// if (errCode == redis.Nil || errCode != nil) || redis_access_token == "" {
	// 	redis_access_token = tool.GetAccessToken()
	// 	errCode := redisModule.Redis.Set(context.Background(), "access_token", redis_access_token, 120*time.Second).Err()
	// 	if errCode != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"msg": "登录失败"})
	// 		return
	// 	}
	// }
	// fmt.Println("redis", redis_access_token)
	// getPhoneNumberUrl := fmt.Sprintf("https://api.weixin.qq.com/wxa/business/getuserphonenumber?access_token=%s", redis_access_token)
	// phone_info := tool.GetCPhone(body.Code, getPhoneNumberUrl)
	// fmt.Println("phone_info.PhoneInfo.PhoneNumber", phone_info.PhoneInfo.PhoneNumber)
	// fmt.Println("phone_info", phone_info)
	// if phone_info.PhoneInfo.PhoneNumber == "" {
	// 	c.JSON(http.StatusOK, gin.H{"msg": "登陆失败", "phone_info": phone_info, "code": body.Code})
	// 	return
	// }
	// user := database.User{
	// 	Phone: phone_info.PhoneInfo.PhoneNumber,
	// }
	// database.DB.Clauses(clause.OnConflict{
	// 	Columns:   []clause.Column{{Name: "phone"}},
	// 	DoUpdates: clause.AssignmentColumns([]string{"updated_at"})}).Create(&user)
	// token, err := tool.GenerateToken(phone_info.PhoneInfo.PhoneNumber)
	// if err != nil {
	// 	c.JSON(http.StatusOK, gin.H{"msg": "登陆失败"})
	// 	return
	// }
	// c.JSON(http.StatusOK, gin.H{"msg": "登陆成功", "token": token})
}
