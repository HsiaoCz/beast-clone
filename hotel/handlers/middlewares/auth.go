package middlewares

import (
	"errors"
	"strings"

	"github.com/HsiaoCz/beast-clone/hotel/types"
	"github.com/gofiber/fiber/v2"
)

const (
	CtxUserInfoKey = "userInfo"
)

func JWTAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.GetReqHeaders()["Authorization"]
		if len(authHeader) == 0 {
			return errors.New("user unlogin")
		}
		authStr := authHeader[0]
		tokenStr := strings.Split(authStr, " ")
		if tokenStr[0] != "Bearer" {
			return errors.New("invalid token,the token string need Bearer for prefix")
		}
		mc, err := ParseToken(tokenStr[1])
		if err != nil {
			return errors.New("invalid Authorization")
		}
		userInfo := &types.UserInfo{
			UserID:  mc.UserID,
			Email:   mc.Email,
			IsAdmin: mc.IsAdmin,
		}
		c.Context().SetUserValue(CtxUserInfoKey, userInfo)
		return c.Next()
	}
}
