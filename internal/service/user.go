package service

import (
	"context"

	"github.com/douyu/jupiter/pkg/xlog"
	v1 "github.com/lyouthzzz/go-web-layout/api/user/v1"
	"github.com/lyouthzzz/go-web-layout/internal/domain"
	"github.com/lyouthzzz/go-web-layout/internal/usecase"
)

type UserRepo interface {
	GetUser(ctx context.Context, userId uint64) (*domain.User, error)
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	DeleteUser(ctx context.Context, userId uint64) error
	ListUser(ctx context.Context, pageNum, pageSize uint64) ([]*domain.User, error)
}

type UserService struct {
	v1.UnimplementedUserServer
	uc  *usecase.UserUsecase
	log *xlog.Logger
}

func NewUserService(uc *usecase.UserUsecase, logger *xlog.Logger) *UserService {
	return &UserService{uc: uc, log: logger.With(xlog.String("pkg", "service/user"))}
}

func (svc *UserService) GetUser(ctx context.Context, getUserReq *v1.GetUserRequest) (*v1.GetUserReply, error) {
	user, err := svc.uc.GetUser(ctx, uint64(getUserReq.Id))
	if err != nil {
		return nil, err
	}
	return &v1.GetUserReply{Id: int64(user.Id)}, nil
}

func (svc *UserService) CreateUser(ctx context.Context, createUserReq *v1.CreateUserRequest) (*v1.CreateUserReply, error) {
	user, err := svc.uc.CreateUser(ctx, &domain.User{
		Username: createUserReq.Username,
		Password: createUserReq.Password,
		Email:    createUserReq.Email,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateUserReply{Id: int64(user.Id)}, nil
}

func (svc *UserService) UpdateUser(ctx context.Context, updateUserReq *v1.UpdateUserRequest) (*v1.UpdateUserReply, error) {
	user, err := svc.uc.UpdateUser(ctx, &domain.User{
		Username: updateUserReq.Username,
		Password: updateUserReq.Password,
		Email:    updateUserReq.Email,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UpdateUserReply{Id: int64(user.Id)}, nil
}

func (svc *UserService) DeleteUser(ctx context.Context, deleteUserReq *v1.DeleteUserRequest) (*v1.Empty, error) {
	err := svc.uc.DeleteUser(ctx, uint64(deleteUserReq.Id))
	if err != nil {
		return nil, err
	}
	return &v1.Empty{}, nil
}
