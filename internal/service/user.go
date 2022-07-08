package service

import (
	"context"
	userbiz "github.com/lyouthzzz/go-web-layout/internal/biz/user"

	userV1 "github.com/lyouthzzz/go-web-layout/api/user/v1"
	userdomain "github.com/lyouthzzz/go-web-layout/internal/domain/user"
)

type UserService struct {
	userV1.UnimplementedUserServiceServer
	user *userbiz.Usecase
}

func NewUserService(uc *userbiz.Usecase) *UserService {
	return &UserService{user: uc}
}

func (svc *UserService) GetUser(ctx context.Context, req *userV1.GetUserRequest) (*userV1.GetUserReply, error) {
	_user, err := svc.user.GetUser(ctx, uint(req.Uid))
	if err != nil {
		return nil, err
	}
	return &userV1.GetUserReply{Uid: int64(_user.UID), Username: _user.Username, Email: _user.Email}, nil
}

func (svc *UserService) CreateUser(ctx context.Context, req *userV1.CreateUserRequest) (*userV1.CreateUserReply, error) {
	_user, err := svc.user.CreateUser(ctx, &userdomain.User{
		Username: req.Username,
		Email:    req.Email,
	})
	if err != nil {
		return nil, err
	}
	return &userV1.CreateUserReply{Uid: int64(_user.UID)}, nil
}

func (svc *UserService) UpdateUser(ctx context.Context, req *userV1.UpdateUserRequest) (*userV1.UpdateUserReply, error) {
	user, err := svc.user.UpdateUser(ctx, &userdomain.User{
		Username: req.Username,
		Email:    req.Email,
	})
	if err != nil {
		return nil, err
	}
	return &userV1.UpdateUserReply{Uid: int64(user.UID)}, nil
}

func (svc *UserService) DeleteUser(ctx context.Context, req *userV1.DeleteUserRequest) (*userV1.Empty, error) {
	err := svc.user.DeleteUser(ctx, uint(req.Uid))
	if err != nil {
		return nil, err
	}
	return &userV1.Empty{}, nil
}
