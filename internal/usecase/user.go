package usecase

import (
	"context"
	"github.com/lyouthzzz/go-web-layout/internal/domain"
)

type UserUsecase struct {
	repo domain.UserRepository
}

func NewUserUsecase(repo domain.UserUsecase) domain.UserUsecase {
	return &UserUsecase{repo: repo}
}

func (u *UserUsecase) GetByName(ctx context.Context, name string) (domain.User, error) {
	return u.repo.GetByName(ctx, name)
}

func (u *UserUsecase) Get(ctx context.Context, id int64) (domain.User, error) {
	return u.repo.Get(ctx, id)
}

func (u *UserUsecase) Create(ctx context.Context, user *domain.User) (domain.User, error) {
	return u.repo.Create(ctx, user)
}

func (u *UserUsecase) Update(ctx context.Context, id int64, user *domain.User) error {
	return u.repo.Update(ctx, id, user)
}

func (u *UserUsecase) Delete(ctx context.Context, id int64) error {
	return u.repo.Delete(ctx, id)
}
