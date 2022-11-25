package routers

import (
	"bubble/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine{
	r := gin.Default()
	r.LoadHTMLGlob("template/*")
	r.Static("/static","./static")
	//待办事项
	r.GET("/",controller.IndexHandler)

	//v1
	v1Group := r.Group("v1")
	{
		//添加
		v1Group.POST("/todo", controller.CreateTodo)
		//查看
		v1Group.GET("/todo", controller.GetTodoList)
		//修改
		v1Group.PUT("/todo/:id", controller.UpdateTodo)
		//删除
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)

		//查看
		v1Group.GET("/endpoint", controller.GetEndpointList)
	}
	return r
}

