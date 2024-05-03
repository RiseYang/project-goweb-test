package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name string `json:"name" form:"name" uri:"name"`
	Age  int    `json:"age" form:"age" uri:"age"`
	Sex  string `json:"sex" form:"sex" uri:"sex"`
}

func main() {

	router := gin.Default()

	router.POST("/index", func(c *gin.Context) {
		var user User
		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.JSON(200, gin.H{"msg": "错误的参数显示"})
			return
		}
		c.JSON(200, user)
	})
	//绑定查询参数
	router.POST("/query", func(c *gin.Context) {
		var user User
		err := c.ShouldBindQuery(&user)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"msg": "你错了"})
			return
		}
		c.JSON(http.StatusOK, user)
	})
	//绑定动态参数
	router.POST("/uri/:name/:age/:sex", func(c *gin.Context) {
		var user User
		err := c.ShouldBindUri(&user)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"msg": "错误参数"})
			return
		}
		c.JSON(http.StatusOK, user)

	})
	//绑定formData
	router.POST("/form", func(c *gin.Context) {
		var user User
		err := c.ShouldBind(&user)
		if err != nil {
			c.JSON(200, gin.H{"msg": "错误的参数显示"})
			return
		}
		c.JSON(200, user)
	})
	router.Run(":8080")
}
