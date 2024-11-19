package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// gin静态文件处理
	// 初始化 redis
	InitRdb()
	r := gin.Default()

	// r.Static("/static", "./static/public") // html js css
	// r.StaticFS("/file", http.Dir("./picture"))

	//
	r.GET("/some", MemeryMiddleware(), func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "hello",
		})
	})
	r.POST("/user", UserMemRegisterHandler)
	r.Run(":3001")
}
