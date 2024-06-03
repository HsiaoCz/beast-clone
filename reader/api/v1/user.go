package v1

import (
	"net/http"

	"github.com/HsiaoCz/beast-clone/reader/models"
	"github.com/HsiaoCz/beast-clone/reader/storage"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	store *storage.Store
}

func NewUserHandler(store *storage.Store) *UserHandler {
	return &UserHandler{
		store: store,
	}
}

func (u *UserHandler) HandleCreateUser(c *gin.Context) {
	user := models.User{}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "please check the request params",
		})
		return
	}
	result, err := u.store.User.CreateUser(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "insert the user failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "create user success",
		"result":  result,
	})
}
