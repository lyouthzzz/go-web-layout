package data

import (
	"context"
	userdomain "github.com/lyouthzzz/go-web-layout/internal/domain/user"
)

var _ userdomain.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
}

func NewUserRepo(data *Data) userdomain.UserRepo {
	return &userRepo{data: data}
}

func (repo *userRepo) GetUser(ctx context.Context, u uint) (*userdomain.User, error) {
	panic("implement me")
}

func (repo *userRepo) CreateUser(ctx context.Context, user *userdomain.User) (*userdomain.User, error) {
	panic("implement me")
}

func (repo *userRepo) UpdateUser(ctx context.Context, user *userdomain.User) (*userdomain.User, error) {
	panic("implement me")
}

func (repo *userRepo) DeleteUser(ctx context.Context, u uint) error {
	panic("implement me")
}
