package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MemeryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.Request.Header.Get("id")
		userInfo, ok := GetMemery(userID)
		if !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "miss user_id",
			})
			ctx.Abort()
			return
		}
		if !userInfo.IsAdmin {
			ctx.JSON(http.StatusNotAcceptable, gin.H{
				"message": "You have no right to accept this data",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
