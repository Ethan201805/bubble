package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var(
	DB *gorm.DB
)

func InitMysql()(err error) {
	dsn := "ruiou:ruiou@2022@(120.77.146.215:3306)/jeecg-boot?charset=utf8mb4&parseTime=True&loc=Local"
	DB,err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

func Close(){
	DB.Close()
}
