package model

import (
	"goblog/pkg/logger"
	"gorm.io/gorm"
	// GORM 的 MySQL 数据库驱动导入
	"gorm.io/driver/mysql"
)

// config := mysql.Config{
// 	User:                 "admin",
// 	Passwd:               "Admin_12",
// 	Addr:                 "180.163.53.44:3360",
// 	Net:                  "tcp",
// 	DBName:               "goblog",
// 	AllowNativePasswords: true,
// }
// DB gorm.DB 对象
var DB *gorm.DB
// ConnectDB 初始化模型
func ConnectDB() *gorm.DB {
	var err error
	config := mysql.New(mysql.Config{
		DSN: "admin:Admin_12@tcp(180.163.53.44:3360)/goblog?charset=utf8mb4&parseTime=True&loc=Local",
	})

	// 准备数据库连接池
	DB,err = gorm.Open(config,&gorm.Config{})
	logger.LogError(err)
	return DB
}