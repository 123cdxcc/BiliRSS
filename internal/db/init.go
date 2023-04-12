package db

import (
	"bilibiliRSS/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	db, err := gorm.Open(sqlite.Open("rss.db"), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败")
	}
	err = db.AutoMigrate(&model.Subscription{})
	if err != nil {
		panic("初始化数据表失败")
	}
	DB = db
}
