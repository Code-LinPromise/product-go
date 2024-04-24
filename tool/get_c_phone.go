package tool

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

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

func GetCPhone(code string, url string) PostResponseData {
	type Payload struct {
		code string
	}
	payload := Payload{
		code: code,
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
		url,
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

	var postResponseData PostResponseData
	err = json.Unmarshal(postBodyBytes, &postResponseData)
	if err != nil {
		log.Fatalf("Error unmarshaling response body: %v", err)
	}
	return postResponseData
}
