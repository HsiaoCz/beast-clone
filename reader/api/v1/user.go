package v1

import (
	"fmt"
	"net/http"

	"github.com/HsiaoCz/beast-clone/reader/api/v1/middleware"
	"github.com/HsiaoCz/beast-clone/reader/models"
	"github.com/HsiaoCz/beast-clone/reader/storage"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	createUserParam := models.UserCreateParams{}
	if err := c.BindJSON(&createUserParam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "please check the request params",
		})
		return
	}
	msg := createUserParam.ValidateParams()
	if len(msg) != 0 {
		c.JSON(http.StatusBadRequest, msg)
		return
	}
	user := models.NewUserFromParams(createUserParam)
	result, err := u.store.User.CreateUser(c.Request.Context(), user)
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

func (u *UserHandler) HandleUpdateUser(c *gin.Context) {
	id := c.Param("uid")
	uid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "please check the request params",
		})
		return
	}
	updateUserParams := models.UserUpdateParams{}
	if err := c.BindJSON(&updateUserParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "please check the request params",
		})
		return
	}
	result, err := u.store.User.UpdateUser(c.Request.Context(), uid, &updateUserParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "update user success",
		"user":    result,
	})
}

func (u *UserHandler) HandleGetUserByID(c *gin.Context) {
	id := c.Param("uid")
	uid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "please check the request param",
		})
		return
	}
	result, err := u.store.User.GetUserByID(c.Request.Context(), uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "get user success",
		"result":  result,
	})
}

func (u *UserHandler) HandleDeleteUserByID(c *gin.Context) {
	id := c.Param("uid")
	uid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "please check the request param",
		})
		return
	}
	if err := u.store.User.DeleteUserByID(c.Request.Context(), uid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "delete user success",
	})
}

func (u *UserHandler) HandleUserLogin(c *gin.Context) {
	userLoginParams := models.UserLoginParams{}
	if err := c.BindJSON(&userLoginParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "please check login params",
		})
		return
	}
	// log.Println(userLoginParams) we dont need this shit
	// who fucking change my config ????
	newUserLoginParams := models.NewUserLoginParams(userLoginParams)
	// log.Println(newUserLoginParams) and this anymore
	user, err := u.store.User.GetUserByEmailAndPassword(c.Request.Context(), &newUserLoginParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "please check the username or password",
		})
		return
	}
	token, err := middleware.GenToken(user.ID, user.Email, user.IsAdmin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "something failed")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"user":   user,
		"token":  token,
	})
}

func (u *UserHandler) HandleGetUsersByString(c *gin.Context) {
	str := c.Query("search")
	fmt.Println(str)
}
