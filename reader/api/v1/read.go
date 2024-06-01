package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadBook(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  http.StatusOK,
		"message": "all is well",
	})
}
