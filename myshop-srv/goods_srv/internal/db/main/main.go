package main

import (
	"fmt"
	"log"
	db1 "myshop/goods_srv/internal/db"
	// "myshop/goods_srv/utils"
	"os"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func main() {
	dsn := "root:Cxp19990707.@tcp(127.0.0.1:3306)/myshop_goods_srv?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second,   // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,         // 禁用彩色打印
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		zap.S().Fatal("init mysql connect faild!")
		panic(err)
	}
	brands := []db1.Brands{}
	res := db.Find(&brands)
	fmt.Printf("res.RowsAffected: %v\n", res.RowsAffected)
	fmt.Println(brands)
	// db.Scopes(utils.Paginate(1, 100)).Find(&brands)
	// fmt.Println(brands)
}