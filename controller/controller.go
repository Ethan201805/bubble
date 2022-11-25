package controller

import (
	"bubble/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateTodo(c *gin.Context) {
	//1 从请求中取数据
	var todo models.Todo
	c.BindJSON(&todo)
	//2 存入数据库
	err := models.CreateTodo(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func GetTodoList(c *gin.Context) {
	todo, err := models.GetTodoList()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func UpdateTodo(c *gin.Context) {
	//获取id
	var id, ok = c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	todo := make(map[string]interface{})
	c.BindJSON(&todo)
	if err := models.UpdateTodo(todo, id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteTodo(c *gin.Context) {
	//获取id
	var id, ok = c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}

	if err := models.DeleteTodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"id": id})
	}
}

func GetEndpointList(c *gin.Context) {
	var singleNode = []interface{}{"Melting point / freezing point", "Boiling point", "Density", "Particle size distribution (Granulometry)", "Vapour pressure", "Partition coefficient (Log Pow)", "Water solubility", "Surface tension", "Flash point", "Auto flammability", "Flammability", "Explosiveness", "Oxidising properties", "Stability in organic solvents and identity of relevant degradation products", "pH", "Dissociation constant", "Viscosity", "Carcinogenicity"}
	var sinfo []*models.Substance
	todo, err := models.GetEndpointList()
	for _, substance := range todo {
		flag := models.CheckEndpoint(substance.Sid)
		if flag {
			sinfo = append(sinfo, substance)
		}
	}

	//整理数据
	for _, v := range sinfo {
		if models.InArray(v.Info_key, singleNode) {
			v.Info_value = v.Info_key
		}
	}
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, sinfo)
	}
}
