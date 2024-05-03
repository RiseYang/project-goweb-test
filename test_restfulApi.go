package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

type ArticleModel struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func _bindJson(c *gin.Context, obj any) (err error) {
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

// 创建文章
func _create(c *gin.Context) {
	//接收前端传来的json
	var article ArticleModel
	err := _bindJson(c, &article)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(200, Response{0, article, "添加成功"})
}

// 编辑文章
func _update(c *gin.Context) {
	fmt.Println(c.Param("id"))
	//接收前端传来的json
	var article ArticleModel
	err := _bindJson(c, &article)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(200, Response{0, article, "修改成功"})
}

// 删除文章
func _delete(c *gin.Context) {
	fmt.Println(c.Param("id"))
	c.JSON(200, Response{0, map[string]string{}, "删除成功"})
}

// 文章列表
func _getList(c *gin.Context) {
	articleList := []ArticleModel{
		{"这是一个GO列子", "这是一个GO列子"},
		{"这是一个java列子", "这是一个java列子"},
		{"这是一个sql列子", "这是一个sql列子"},
	}

	c.JSON(200, Response{0, articleList, "成功"})
}

// 文章详情
func _getDetail(c *gin.Context) {
	//获取param的id
	fmt.Println(c.Param("id"))
	article := ArticleModel{
		"这是一个GO列子", "这是一个GO列子",
	}
	c.JSON(200, Response{0, article, "成功"})
}
func main() {

	router := gin.Default()

	router.GET("/articles", _getList)       //文章列表
	router.GET("/articles/:id", _getDetail) //文章详情
	router.POST("/articles", _create)       //添加文章
	router.PUT("/articles/:id", _update)    //修改文章
	router.DELETE("/articles/:id", _delete) //删除文章

	router.Run(":8080")
}
