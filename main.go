package main

import (
	"bubble/dao"
	"bubble/routers"
)

func main() {
	//连接数据库
	err := dao.InitMysql()
	if err != nil {
		panic(err)
	}
	//关闭数据库
	defer dao.Close()
	//模型绑定
	//dao.DB.AutoMigrate(&models.Todo{})
	//注册路由
	r := routers.SetupRouter()
	r.Run(":9091")
}
