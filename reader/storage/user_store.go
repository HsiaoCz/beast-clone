package storage

import (
	"context"

	"github.com/HsiaoCz/beast-clone/reader/models"
)

type UserStorer interface {
	CreateUser(context.Context, *models.User) (*models.User, error)
}
