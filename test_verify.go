package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Password  string `json:"password"`
	Rpassword string `json:"rpassword"`
}

func main() {
	r := gin.Default()

	r.POST("/index", func(c *gin.Context) {
		var user UserInfo
		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.JSON(200, gin.H{"msg": "错误参数!"})
			return
		}
		c.JSON(http.StatusOK, user)
	})

	r.Run(":8088")
}
