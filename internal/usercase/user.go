package usercase

import (
	"context"
	"github.com/lyouthzzz/go-web-layout/internal/domain"
)

type User struct {
	repo domain.UserRepository
}

func (u *User) Get(ctx context.Context, id int64) (domain.User, error) {
	return u.repo.Get(ctx, id)
}

func (u *User) Create(ctx context.Context, username, password, email string) (domain.User, error) {
	return u.repo.Create(ctx, username, password, email)
}

func (u *User) Update(ctx context.Context, id int64, user *domain.User) error {
	return u.repo.Update(ctx, id, user)
}

func (u *User) Delete(ctx context.Context, id int64) error {
	return u.repo.Delete(ctx, id)
}
