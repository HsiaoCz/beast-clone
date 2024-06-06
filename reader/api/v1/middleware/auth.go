package middleware

import (
	"net/http"
	"strings"

	"github.com/HsiaoCz/beast-clone/reader/models"
	"github.com/gin-gonic/gin"
)

const (
	CtxUserInfoKey = "userInfo"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// Authorization Bearer Token
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, "user need login")
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.Split(authHeader, " ")
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusBadRequest, "invalid token")
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusBadRequest, "invalid token")
			c.Abort()
			return
		}
		userInfo := models.UserInfo{
			UserID:  mc.UserID,
			Email:   mc.Email,
			IsAdmin: mc.IsAdmin,
		}
		// 将当前请求的userID信息保存到请求的上下文c上
		c.Set(CtxUserInfoKey, &userInfo)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
