package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var mysqlLogger logger.Interface

func init() {
	username := "root"
	password := "root"
	host := "127.0.0.1"
	port := 3306
	dbname := "gorm"
	timeout := "10s"

	mysqlLogger = logger.Default.LogMode(logger.Info)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, dbname, timeout)
	// 连接MySQL，获得DB类型实例，用于后面的数据库读写操作。
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		//Logger:                 mysqlLogger,
	})
	//db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("连接数据库失败，error=" + err.Error())
	}
	fmt.Println("连接成功")
	DB = db
	fmt.Println(db)

}

//type Student struct {
//	ID    uint    `gorm:"size:10"`
//	Name  string  `gorm:"size:16"`
//	Age   int     `gorm:"size:3"`
//	Email *string `gorm:"size:128"`
//}
//
//func main() {
//	DB = DB.Session(&gorm.Session{
//		Logger: mysqlLogger,
//	})
//	// 自动迁移并创建表
//	if err := DB.AutoMigrate(&Student{}); err != nil {
//		panic("自动迁移（创建表）失败，error=" + err.Error())
//	}
//	DB.Debug().AutoMigrate(&Student{})
//
//}
