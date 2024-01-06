package service

import (
	"context"
	apiV1 "github.com/lyouthzzz/go-web-layout/api/v1"
	"github.com/lyouthzzz/go-web-layout/internal/biz"
)

type UserService struct {
	apiV1.UnimplementedUserServiceServer
	user *biz.Usecase
}

func NewUserService(uc *biz.Usecase) *UserService {
	return &UserService{user: uc}
}

func (svc *UserService) GetUser(ctx context.Context, req *apiV1.GetUserRequest) (*apiV1.User, error) {
	user, err := svc.user.GetUser(ctx, req.Uid)
	if err != nil {
		return nil, err
	}
	return &apiV1.User{Uid: user.Uid, Username: user.Username, Email: user.Email}, nil
}

func (svc *UserService) CreateUser(ctx context.Context, req *apiV1.CreateUserRequest) (*apiV1.Empty, error) {
	if err := svc.user.CreateUser(ctx, &biz.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}); err != nil {
		return nil, err
	}
	return &apiV1.Empty{}, nil
}

func (svc *UserService) UpdateUser(ctx context.Context, req *apiV1.UpdateUserRequest) (*apiV1.Empty, error) {
	if err := svc.user.UpdateUser(ctx, &biz.User{
		Password: req.Password,
		Email:    req.Email,
	}); err != nil {
		return nil, err
	}
	return &apiV1.Empty{}, nil
}

func (svc *UserService) DeleteUser(ctx context.Context, req *apiV1.DeleteUserRequest) (*apiV1.Empty, error) {
	if err := svc.user.DeleteUser(ctx, req.Uid); err != nil {
		return nil, err
	}
	return &apiV1.Empty{}, nil
}
