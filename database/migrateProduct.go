package database

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name          string               `json:"product_name"`
	ProductKind   string               `json:"product_kind"`
	Price         uint                 `json:"product_price"`
	ProductCover  string               `json:"product_cover"`
	ProdctDetails []ProdctDetailsImage `gorm:"foreignKey:ProductID;references:ID" json:"product_details"`
	// 添加别名以控制 gorm.Model 字段的 JSON 表现
	ID        uint           `gorm:"<-:create;->:ignore" json:"product_id"` // 主键
	CreatedAt time.Time      `gorm:"<-:create;->:ignore" json:"created_at"` // 创建时间
	UpdatedAt time.Time      `gorm:"<-:create;->:ignore" json:"updated_at"` // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"<-:create;->:ignore" json:"deleted_at"` // 删除标记（软删除）
}

type ProdctDetailsImage struct {
	gorm.Model
	ProductID uint   `gorm:"not null" json:"product_id"`
	Images    string `json:"detail_url"`
}

func MigrateProduct() *Product {
	return &Product{}
}
func MigrateProdctDetailsImage() *ProdctDetailsImage {
	return &ProdctDetailsImage{}
}
