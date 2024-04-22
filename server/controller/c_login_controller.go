package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

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
	access_token_url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=%s&appid=%s&secret=%s", "client_credential", appid, secret)
	resp, err := http.Get(access_token_url)
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
		Access_Token string `json:"access_token"`
		ExpiresIn    int    `json:"expires_in"`
	}
	var responseData ResponseData
	err = json.Unmarshal(bodyBytes, &responseData)
	if err != nil {
		log.Fatalf("Error unmarshaling response body: %v", err)
	}
	// 打印响应体（通常需要进一步解析成所需的数据结构）
	getPhoneNumberUrl := fmt.Sprintf("https://api.weixin.qq.com/wxa/business/getuserphonenumber?access_token=%s", responseData.Access_Token)
	type Payload struct {
		code string
	}
	payload := Payload{
		code: body.Code,
	}
	// 将请求数据编码为JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Error marshaling payload to JSON: %v", err)
	}
	// 创建POST请求
	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		getPhoneNumberUrl,
		bytes.NewBuffer(jsonPayload),
	)
	if err != nil {
		log.Fatalf("Error creating POST request: %v", err)
	}
	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 创建HTTP客户端
	client := &http.Client{
		Timeout: time.Second * 10, // 设置请求超时为10秒
	}

	// 发送POST请求
	postResp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending POST request: %v", err)
	}
	defer postResp.Body.Close()

	// 检查响应状态码是否表示成功（通常为200 OK或201 Created等）
	if postResp.StatusCode < 200 || postResp.StatusCode >= 300 {
		log.Fatalf("Unexpected status code: %d", postResp.StatusCode)
	}

	// 读取响应体
	postBodyBytes, err := ioutil.ReadAll(postResp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}
	type WaterMarkType struct {
		TimeStamp int    `timestamp`
		AppId     string `appid`
	}
	type PhoneInfoType struct {
		PhoneNumber     string        `json:"phoneNumber"`
		PurePhoneNumber string        `json:"purePhoneNumber"`
		CountryCode     uint          `json:"countryCode"`
		WaterMark       WaterMarkType `json:"watermark"`
	}
	type PostResponseData struct {
		ErrCode   uint          `json:"errcode"`
		ErrMsg    string        `json:"errmsg"`
		PhoneInfo PhoneInfoType `json:"phone_info"`
	}
	var postResponseData PostResponseData
	err = json.Unmarshal(postBodyBytes, &postResponseData)
	if err != nil {
		log.Fatalf("Error unmarshaling response body: %v", err)
	}
}
