package service

import (
	"context"
	v1 "github.com/lyouthzzz/go-web-layout/api/v1"
	"github.com/lyouthzzz/go-web-layout/internal/domain"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

type UserService struct {
	v1.UnimplementedUserServiceServer
	uc domain.IUserUsecase
	sc domain.ISessionUsecase
}

func NewUserService(uc domain.IUserUsecase, sc domain.ISessionUsecase) *UserService {
	return &UserService{uc: uc, sc: sc}
}

func (s *UserService) Login(ctx context.Context, req *v1.UserLoginRequest) (*v1.UserLoginResponse, error) {
	user, err := s.uc.GetByName(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	if user.Password != req.Password {
		return nil, errors.New("password incorrect")
	}
	session, err := s.sc.Create(ctx, user.ID)
	if err != nil {
		return nil, err
	}
	return &v1.UserLoginResponse{
		User:    &v1.User{Name: user.Username, Email: user.Email},
		Session: &v1.Session{UserId: cast.ToUint32(user.ID), Id: session.Id},
	}, nil
}

func (s *UserService) Logout(ctx context.Context, req *v1.UserLogoutRequest) (*v1.Empty, error) {
	err := s.sc.Delete(ctx, req.SessionId)
	if err != nil {
		return &v1.Empty{}, err
	}
	return &v1.Empty{}, nil
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
