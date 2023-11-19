package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//vender
// 1. 一定要将代码新建到gopath目录下的src
// 2. 要记得设置GO111MODULE=off , 开始go modules 要记得 GO111MODULE=on
// 先查找 gopath/src 这个目录下的包是否有 goroot/src目录下找
// 其实就是不做包管理

//包管理 - 异常处理 泛型

// 能用go modules 就用moduls 不用去考虑以前的开发模式
// 即使你使用了以前的模式， 也可以自动设置为现在的modules模式
// go modules 是一个统一的做法
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
