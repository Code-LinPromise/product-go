package tool

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetAccessToken() string {
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
	return responseData.Access_Token
}
