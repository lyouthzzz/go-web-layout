package usecase

import (
	"context"

	"github.com/douyu/jupiter/pkg/xlog"
	v1 "github.com/lyouthzzz/go-web-layout/api/user/v1"
	"github.com/lyouthzzz/go-web-layout/internal/data"
	"github.com/lyouthzzz/go-web-layout/internal/domain"
	"github.com/lyouthzzz/go-web-layout/pkg/ecode"
)

type UserUsecase struct {
	repo data.UserRepo
	log  *xlog.Logger
}

func NewUserUsecase(repo data.UserRepo, logger *xlog.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: logger.With(xlog.String("pkg", "usecase/user"))}
}

func (uc *UserUsecase) GetUser(ctx context.Context, userId uint64) (*domain.User, error) {
	return uc.repo.GetUser(ctx, userId)
}

func (uc *UserUsecase) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	return uc.repo.CreateUser(ctx, user)
}

func (uc *UserUsecase) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	return uc.repo.UpdateUser(ctx, user)
}

func (uc *UserUsecase) DeleteUser(ctx context.Context, userId uint64) error {
	err := uc.repo.DeleteUser(ctx, userId)
	if err != nil && ecode.Reason(err) == v1.ErrorReason_USER_NOT_FOUND.String() {
		return nil
	}
	return err
}
