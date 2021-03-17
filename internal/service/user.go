package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/lyouthzzz/framework/pkg/auth/authn"
	v1 "github.com/lyouthzzz/go-web-layout/api/v1"
	"github.com/lyouthzzz/go-web-layout/internal/domain"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"strings"
)

type UserService struct {
	v1.UnimplementedUserServiceServer
	authN authn.Authenticator
	uc    domain.IUserUsecase
}

func NewUserService(uc domain.IUserUsecase, authN authn.Authenticator) *UserService {
	return &UserService{uc: uc, authN: authN}
}

func (s *UserService) Login(ctx context.Context, req *v1.UserLoginRequest) (*v1.UserLoginResponse, error) {
	user, err := s.uc.GetByName(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	if user.Password != req.Password {
		return nil, errors.New("password incorrect")
	}
	tokenAuth := &authn.TokenAuthentication{Token: strings.ReplaceAll(uuid.New().String(), "-", "")}
	if err := s.authN.WriteAuthentication(ctx, tokenAuth, user); err != nil {
		return nil, err
	}
	return &v1.UserLoginResponse{
		User:    &v1.User{Name: user.Username, Email: user.Email},
		Session: &v1.Session{UserId: cast.ToUint32(user.ID), Id: tokenAuth.Token},
	}, nil
}

func (s *UserService) Logout(ctx context.Context, req *v1.UserLogoutRequest) (*v1.Empty, error) {
	tokenAuth := authn.TokenAuthentication{Token: req.SessionId}
	if err := s.authN.DeleteAuthentication(ctx, tokenAuth); err != nil {
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
