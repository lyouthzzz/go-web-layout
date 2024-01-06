package data

import (
	"context"
	"github.com/lyouthzzz/go-web-layout/internal/biz"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
}

func NewUserRepo(data *Data) biz.UserRepo {
	return &userRepo{data: data}
}

func (repo *userRepo) GetUser(ctx context.Context, id int64) (*biz.User, error) {
	panic("GetUser implement me")
}

func (repo *userRepo) CreateUser(ctx context.Context, user *biz.User) error {
	panic("CreateUser implement me")
}

func (repo *userRepo) UpdateUser(ctx context.Context, user *biz.User) error {
	panic("UpdateUser implement me")
}

func (repo *userRepo) DeleteUser(ctx context.Context, uid int64) error {
	panic("DeleteUser implement me")
}

func (repo *userRepo) ListUser(ctx context.Context, offset, limit int64) (*biz.UserPage, error) {
	panic("ListUser implement me")
}
