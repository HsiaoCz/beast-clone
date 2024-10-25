package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UserRegisterHandler(c *gin.Context) {
	var userReg UserRegister
	if err := c.ShouldBind(&userReg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "something wrong",
			"error":   err,
		})
		return
	}
	userInfo := UserInfo{
		Email:   userReg.Email,
		IsAdmin: userReg.IsAdmin,
		UserID:  uuid.NewString(),
	}
	ctx := context.WithValue(c.Request.Context(), UserInfoKey, userInfo)
	c.Request = c.Request.WithContext(ctx)
	c.JSON(http.StatusOK, gin.H{
		"message": "welcome to the ceberpunk",
		"userID":  userInfo.UserID,
	})
}
