package models

import (
	"log"

	"github.com/iot-platform/define"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDB() {
	dsn := define.MysqlDSN + "/iot-platform?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("[DB ERROR]: ", err.Error())
	}

	if err := db.AutoMigrate(
		&DeviceBasic{},
		&ProductBasic{},
		&UserBasic{},
	); err != nil {
		log.Fatalln("[DB ERROR]: ", err.Error())
	}

	DB = db
}
