package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"product.com/m/database"
)

func CGetKindImgController(c *gin.Context) {
	var KIND = [...]string{
		"本月爆款产品",
		"实木楼梯",
		"玻璃楼梯",
		"金属楼梯",
		"水泥楼梯",
		"铁艺护栏",
		"铝合金护栏",
		"铜艺护栏",
		"玻璃护栏",
		"实木护栏",
		"铝合金庭院门",
		"阳台护栏",
		"屏风隔断",
		"车库门",
		"阁楼梯",
		"电梯",
	}
	var kindImg []database.KindImg
	err := database.DB.Where("now_kind_img=?", 1).Find(&kindImg).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": "请求失败", "err": err})
	}
	type KindInfo struct {
		Name   string `json:"name"`
		PicUrl string `json:"pic_url"`
	}
	var kindInfo []KindInfo
	for _, kind := range KIND {
		// 查找kindImg中是否有与kind匹配的Name
		var matchingPicUrl string
		for _, img := range kindImg {
			if img.KindName == kind {
				matchingPicUrl = img.ImageAddr
				break
			}
		}
		// 将匹配到的PicUrl写入结果，否则写入空字符串（这里假设null用空字符串表示）
		kindInfo = append(kindInfo, KindInfo{
			Name:   kind,
			PicUrl: matchingPicUrl,
		})

	}
	c.JSON(http.StatusOK, gin.H{"msg": "请求成功", "product_kind": kindInfo})
}
