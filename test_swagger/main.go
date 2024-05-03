package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files" // swagger embed files
	gs "github.com/swaggo/gin-swagger"     // gin-swagger middleware
	_ "project-goweb-test/test_swagger/docs"
)

func Index(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "hello,go!"})
}

// @title API文档
// @version 1.0
// @description API文档
// @host 127.0.0.1:8088
// @BasePath /
func main() {
	r := gin.Default()
	r.GET("/", Index)
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	r.Run(":8088")
}
