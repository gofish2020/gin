package main

import (
	"fmt"
	"net/http"

	"github.com/gofish2020/gin"
)

func handlerTest1(c *gin.Context) {
	c.String(http.StatusOK, c.Request.URL.Path)
}

func main() {
	// 初始化 *Engine 对象
	router := gin.New()

	router.GET("/a", handlerTest1) // 路径 /a
	router.GET("/b", handlerTest1) // 路径 /b
	router.GET("/", handlerTest1)  // 路径 /

	// 这里相当于公共的前缀 /ab
	group := router.Group("/ab")
	{
		group.GET("/a", handlerTest1) // 路径 /ab/a
		group.GET("/b", handlerTest1) // 路径 /ab/b
	}
	// "/static" 表示访问的路由
	// "." 访问的文件在服务器的存储目录
	// 访问范例: http://127.0.0.1:8080/static/test.txt
	router.Static("/static", ".")

	// 遍历压缩前缀树
	list := router.Routes()
	for _, l := range list {
		fmt.Println(l.Method, l.Path)
	}

	// 启动服务
	router.Run(":8080")
}
