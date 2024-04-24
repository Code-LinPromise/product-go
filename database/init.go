package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	username := "root"
	password := "123456"
	host := "172.21.80.1"
	port := 3306
	DBname := "product"
	timeout := "10s"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, DBname, timeout)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	DB = db
	if err != nil {
		panic(err)
	}
}

func MigrateDB() {
	DB.AutoMigrate(MigrateAdminDB())
	DB.AutoMigrate(MigrateProduct(), MigrateProdctDetailsImage())
	DB.AutoMigrate(MigrateCarousel())
	DB.AutoMigrate(MigrateVideo())
	DB.AutoMigrate(MigratePhone())
	DB.AutoMigrate(MigrateKindImage())
	DB.AutoMigrate(MigrateUser(), MigrateLiked())
}
