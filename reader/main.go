package main

import (
	v1 "github.com/HsiaoCz/beast-clone/reader/api/v1"
	"github.com/gin-gonic/gin"
)

func main() {
	var (
		router = gin.Default()
	)

	router.GET("/read", v1.ReadBook)

	router.Run(":3009")
}
