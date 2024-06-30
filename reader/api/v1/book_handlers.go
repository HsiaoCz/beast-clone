package v1

import (
	"net/http"

	"github.com/HsiaoCz/beast-clone/reader/storage"
	"github.com/gin-gonic/gin"
)

type BookHandlers struct {
	store *storage.Store
}

func NewBookHandlers(store *storage.Store) *BookHandlers {
	return &BookHandlers{
		store: store,
	}
}

func (b *BookHandlers) HandleCreateBook(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "all is well",
	})
}
