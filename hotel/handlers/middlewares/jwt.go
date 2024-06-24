package middlewares

import (
	"errors"
	"log/slog"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CustomCliams 自定义声明类型 并内嵌jwt.RegisteredClaims
// jwt包自带的jwt.RegisteredClaims只包含了官方字段

type myClaims struct {
	UserID  primitive.ObjectID `json:"userID"`
	Email   string             `json:"email"`
	IsAdmin bool               `json:"isAdmin"`
	jwt.StandardClaims
}

// 定义过期时间
const TokenExpireDuration = time.Hour * 24 * 3

var mySecret = []byte(os.Getenv("JWTSECRET"))

// GenToken 生成JWT
func GenToken(userID primitive.ObjectID, email string, isAdmin bool) (string, error) {
	claims := myClaims{
		userID,
		email,
		isAdmin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "hotel-hsiaol1",
		},
	}
	// use jwt generate
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(mySecret)
	if err != nil {
		slog.Error("token failed", "err", err)
		return "", err
	}
	return token, nil
}

// ParseToken parse JWT
func ParseToken(tokenString string) (*myClaims, error) {
	// parse
	token, err := jwt.ParseWithClaims(tokenString, &myClaims{}, func(t *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*myClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
