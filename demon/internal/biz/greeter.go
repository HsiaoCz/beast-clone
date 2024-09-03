package biz

import (
	"context"

	v1 "demon/api/helloworld/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Greeter is a Greeter model.
type User struct {
	gorm.Model
	UserID          string `gorm:"column:user_id"`
	Username        string `gorm:"column:username"`
	Password        string `gorm:"column:password"`
	Email           string `gorm:"column:email"`
	Synopsis        string `gorm:"column:synopsis"`
	Avatar          string `gorm:"column:avatar"`
	BackgroundImage string `gorm:"column:background_image"`
}

// GreeterRepo is a Greater repo.
type GreeterRepo interface {
	CreateUser(context.Context, *User) (*User, error)
	FindByID(context.Context, string) (*User, error)
}

// GreeterUsecase is a Greeter usecase.
type GreeterUsecase struct {
	repo GreeterRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewGreeterUsecase(repo GreeterRepo, logger log.Logger) *GreeterUsecase {
	return &GreeterUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *GreeterUsecase) CreateGreeter(ctx context.Context, g *User) (*User, error) {
	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g)
	return uc.repo.CreateUser(ctx, g)
}

