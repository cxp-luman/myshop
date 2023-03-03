package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

func genMD5(code string) string {
	MD5 := md5.New()
	_, _ = io.WriteString(MD5, code)
	return hex.EncodeToString(MD5.Sum(nil))
}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	//dsn := "root:Cxp19990707.@tcp(localhost:3306)/myshop_users_srv?charset=utf8mb4&parseTime=True&loc=Local"
	//newLogger := logger.New(
	//	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
	//	logger.Config{
	//		SlowThreshold: time.Second, // 慢 SQL 阈值
	//		LogLevel:      logger.Info, // 日志级别
	//		Colorful:      true,        // 禁用彩色打印
	//	},
	//)
	//
	//// 全局模式
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	//	NamingStrategy: schema.NamingStrategy{
	//		SingularTable: true,
	//	},
	//	Logger: newLogger,
	//})
	//if err != nil {
	//	fmt.Println(err)
	//	panic("连接失败")
	//}
	//db.AutoMigrate(&model.User{})
	fmt.Println(genMD5("12345"))

}
