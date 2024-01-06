package biz

import (
	"context"
)

type Usecase struct {
	user UserRepo
}

func NewUserUsecase(user UserRepo) *Usecase {
	return &Usecase{user: user}
}

func (uc *Usecase) GetUser(ctx context.Context, uid int64) (*User, error) {
	return uc.user.GetUser(ctx, uid)
}

func (uc *Usecase) CreateUser(ctx context.Context, user *User) error {
	return uc.user.CreateUser(ctx, user)
}

func (uc *Usecase) UpdateUser(ctx context.Context, user *User) error {
	return uc.user.UpdateUser(ctx, user)
}

func (uc *Usecase) DeleteUser(ctx context.Context, uid int64) error {
	return uc.user.DeleteUser(ctx, uid)
}
