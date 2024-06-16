package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/HsiaoCz/beast-clone/hotel/types"
)

type CtxKey string

const (
	CtxUserInfoKey CtxKey = "userInfo"
)

// JWT Middleware
func JwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		// Split the header to get the token part
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}
		tokenString := parts[1]

		// Parse and validate the token
		mc, err := ParseToken(tokenString)
		if err != nil {
			http.Error(w, "Invalid token string", http.StatusUnauthorized)
			return
		}
		userInfo := &types.UserInfo{
			UserID:  mc.UserID,
			Email:   mc.Email,
			IsAdmin: mc.IsAdmin,
		}
		ctx := context.WithValue(r.Context(), CtxUserInfoKey, userInfo)
		r = r.WithContext(ctx)
		// Pass the request to the next handler
		next.ServeHTTP(w, r)
	})
}
