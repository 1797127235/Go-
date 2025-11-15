package initialize

import (
	"blog-server/global"
	"blog-server/model"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql() {
	m := global.Config.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.User, m.Password, m.Host, m.Port, m.DB,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("connect mysql failed:", err)
	}

	// 自动建表
	if err := db.AutoMigrate(&model.User{}, &model.Post{}, &model.Category{}, &model.Comment{}); err != nil {
		log.Fatal("auto migrate failed:", err)
	}

	global.DB = db
}
