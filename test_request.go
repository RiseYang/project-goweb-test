package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 查询参数
func _req(c *gin.Context) {
	fmt.Println(c.Query("user"))
	fmt.Println(c.GetQuery("user"))
	fmt.Println(c.QueryArray("user")) //拿到多个查询参数 ?user=zs&user=ls

}

// 动态参数
func _param(c *gin.Context) {
	fmt.Println(c.Param("user_id"))
	fmt.Println(c.Param("book_id"))
}

// 表单参数
func _form(c *gin.Context) {
	fmt.Println(c.PostForm("name"))
	fmt.Println(c.PostFormArray("name"))
	fmt.Println(c.DefaultPostForm("addr", "四川")) //用户没传值,为默认值
	fmt.Println(c.MultipartForm())               //接受所有参数

}

// 解析json
func bindJson(c *gin.Context, obj any) (err error) {
	body, _ := c.GetRawData()
	contentType := c.GetHeader("Content-Type")
	switch contentType {
	case "application/json":
		err := json.Unmarshal(body, &obj)
		if err != nil {
			fmt.Println(err.Error())
			err.Error()
		}
	}
	return nil
}

// 原始参数
func _raw(c *gin.Context) {
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	var user User
	err := bindJson(c, &user)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(user)

}

func main() {
	router := gin.Default()

	router.GET("/query", _req)
	router.GET("/param/:user_id/", _param)
	router.GET("/param/:user_id/:book_id", _param)
	router.POST("/form", _form)
	router.POST("/raw", _raw)
	router.Run(":8080")
}
