package service

import (
	"context"

	v1 "demon/api/helloworld/v1"
	"demon/internal/biz"
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

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) CreateUser(ctx context.Context, in *v1.CreateUserRequest) (*v1.CreateUserResponse, error) {
	return &v1.CreateUserResponse{}, nil
}

func (s *GreeterService) GetUserByID(ctx context.Context, in *v1.CreateUserRequest) (*v1.CreateUserResponse, error) {
	return &v1.CreateUserResponse{}, nil
}
