package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// gin 中间件
// 1.内置中间件
// 2.自定义中间件

// 自定义中间件
func CustomMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user_info, ok := ctx.Request.Context().Value(UserInfoKey).(UserInfo)
		if !ok {
			ctx.JSON(http.StatusNonAuthoritativeInfo, gin.H{
				"message": "You have no author infomation,please login",
			})
			ctx.Abort()
			return
		}
		if !user_info.IsAdmin {
			ctx.JSON(http.StatusNotAcceptable, gin.H{
				"message": "You are not admin,so you have no rights to accept the infomations",
			})
			ctx.Abort()
			return
		}
		logrus.WithFields(logrus.Fields{
			"user_id": user_info.UserID,
			"email":   user_info.Email,
		}).Info("user accept the resource")
		ctx.Next()
	}
}
