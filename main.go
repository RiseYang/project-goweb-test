package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(context *gin.Context) {
	context.String(http.StatusOK, "hello world!")
}

func _Json(c *gin.Context) {
	//json响应结构体
	type UserInfo struct {
		UserName string `json:"user_name"`
		Age      int    `json:"age"`
		Sex      string `json:"sex"`
		Password string `json:"-"` //忽略密码
	}
	//user := UserInfo{"张三", 18, "男", "123456"}

	//json响应map
	//userMap := map[string]string{
	//	"username": "张三",
	//	"age":      "18",
	//}

	//直接响应json
	c.JSON(http.StatusOK, gin.H{
		"username": "张三",
		"age":      18,
	})
}

func _xml(c *gin.Context) {
	c.XML(http.StatusOK, gin.H{"user": "张三", "pwd": 123456, "sex": "男"})
}

func _yaml(c *gin.Context) {
	c.YAML(http.StatusOK, gin.H{"user": "张三", "pwd": 123456, "sex": "男"})
}

func _html(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

// 重定向
func _redirect(c *gin.Context) {
	c.Redirect(302, "https://www.baidu.com")
}
func main() {
	//创建一个默认路由
	router := gin.Default()
	//响应模板
	router.LoadHTMLGlob("templates/*")
	//在goland中,没有相对文件的路径,它只有相对项目的路径
	//配置单个文件,网页请求的路由,文件的路径
	router.StaticFile("./static/GOPR0767.JPG", "./static/GOPR0767.JPG")

	//绑定路由和路由函数,访问/index路由,将由对应的函数去处理
	router.GET("/index", Index)

	router.GET("/json", _Json)

	router.GET("/xml", _xml)

	router.GET("/yaml", _yaml)

	router.GET("/html", _html)

	router.GET("/baidu", _redirect)

	//两种启动方式
	//启动监听, gin会把web服务运行在本机的端口
	router.Run(":8080")
	//用原生http服务的方式
	//http.ListenAndServe(":8080", router)
}
