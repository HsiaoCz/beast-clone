package service

import (
	"context"

	v1 "github.com/HsiaoCz/beast-clone/latency/api/helloworld/v1"
	"github.com/HsiaoCz/beast-clone/latency/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

func (s *GreeterService) CreateUser(ctx context.Context, in *v1.CreateUserRequest) (*v1.CreateUserResponse, error) {
	return &v1.CreateUserResponse{}, nil
}

func (s *GreeterService) UpdateUser(ctx context.Context, in *v1.UpdateUserRequest) (*v1.UpdateUserResponse, error) {
	return &v1.UpdateUserResponse{}, nil
}

func (s *GreeterService) DeleteUser(ctx context.Context, in *v1.DeleteUserRequest) (*v1.DeleteUserResponse, error) {
	return &v1.DeleteUserResponse{}, nil
}

func (s *GreeterService) GetUserByID(ctx context.Context, in *v1.GetUserByIDRequest) (*v1.GetUserByIDResponse, error) {
	return &v1.GetUserByIDResponse{}, nil
}
