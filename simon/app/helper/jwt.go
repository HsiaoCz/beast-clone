package helper

import (
	"errors"
	"log/slog"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// customer the jwt claims
// contain the userID and email
// we can use userID verify the user
// question: the email hava to exist?
type myClaims struct {
	UserID  primitive.ObjectID `json:"userID"`
	Email   string             `json:"email"`
	IsAdmin bool               `json:"isAdmin"`
	jwt.StandardClaims
}

// define expire time
const TokenExpireDuration = time.Hour * 24 * 3

var mySecret = []byte(os.Getenv("JWTSECRET"))

// GenToken generate JWT
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
