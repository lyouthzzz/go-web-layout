package service

import (
	"context"
	v1 "github.com/lyouthzzz/go-web-layout/api/v1"
	"github.com/lyouthzzz/go-web-layout/internal/domain"
)

type UserService struct {
	v1.UnimplementedUserServiceServer
	uc domain.UserUsecase
}

func NewUserService(uc domain.UserUsecase) *UserService {
	return &UserService{uc: uc}
}

func (s *UserService) Create(ctx context.Context, req *v1.CreateUserRequest) (*v1.User, error) {
	user, err := s.uc.Create(ctx, &domain.User{Username: req.User.Name, Password: req.User.Password, Email: req.User.Email})
	if err != nil {
		return nil, err
	}
	return &v1.User{Name: user.Username, Password: user.Password, Email: user.Email}, nil
}

func (s *UserService) Update(ctx context.Context, req *v1.UpdateUserRequest) (*v1.User, error) {
	user := &domain.User{Username: req.User.Name, Password: req.User.Password, Email: req.User.Email}
	err := s.uc.Update(ctx, int64(req.Id), user)
	if err != nil {
		return nil, err
	}
	return &v1.User{}, nil
}

func (s *UserService) Get(ctx context.Context, req *v1.GetUserRequest) (*v1.User, error) {
	user, err := s.uc.Get(ctx, int64(req.Id))
	if err != nil {
		return nil, err
	}
	return &v1.User{Name: user.Username, Password: user.Password, Email: user.Email}, nil
}

func (s *UserService) Delete(ctx context.Context, req *v1.DeleteUserRequest) (*v1.Empty, error) {
	if err := s.uc.Delete(ctx, int64(req.Id)); err != nil {
		return nil, err
	}
	return &v1.Empty{}, nil
}
