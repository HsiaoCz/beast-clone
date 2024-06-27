package biz

import (
	"context"
	"time"

	v1 "github.com/HsiaoCz/beast-clone/latency/api/helloworld/v1"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Greeter is a Greeter model.
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string             `bson:"firstName"`
	LastName  string             `bson:"lastName"`
	Email     string             `bson:"lastName"`
	Content   string             `bson:"content"`
	Avatar    string             `bson:"avatar"`
	CreatedAt time.Time          `bson:"createdAt"`
	Password  string             `bson:"password"`
	IsAdmin   bool               `bson:"isAdmin"`
}

// GreeterRepo is a Greater repo.
type GreeterRepo interface {
	Save(context.Context, *User) (*User, error)
	Update(context.Context, *User) (*User, error)
	FindByID(context.Context, int64) (*User, error)
	ListAll(context.Context) ([]*User, error)
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
func (uc *GreeterUsecase) CreateGreeter(ctx context.Context, u *User) (*User, error) {
	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", u.FirstName)
	return uc.repo.Save(ctx, u)
}
