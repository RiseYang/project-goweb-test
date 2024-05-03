package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func main() {

	router := gin.Default()

	router.GET("/header", func(c *gin.Context) {
		//获取请求头 字母不含大小写
		fmt.Println(c.GetHeader("User-Agent"))
		//fmt.Println(c.GetHeader("user-agent"))
		//fmt.Println(c.GetHeader("user-AGent"))
		fmt.Println(c.Request.Header)
		fmt.Println(c.Request.Header.Get("User-Agent"))
		//自定义请求头token, 字母不含大小写
		fmt.Println(c.Request.Header.Get("Token"))
		fmt.Println(c.Request.Header.Get("token"))

		c.JSON(200, gin.H{"msg": "成功"})
	})

	//爬虫和用户的区别对待
	router.GET("/index", func(c *gin.Context) {
		userAgent := c.GetHeader("User-Agent")
		//取正则去匹配
		//判断字符串是否包含
		if strings.Contains(userAgent, "python") {
			//爬虫来
			c.JSON(0, gin.H{"data": "这是爬虫的数据"})
			return
		}
		c.JSON(0, gin.H{"data": "这是用户的数据"})
	})

	router.Run(":8080")
}
