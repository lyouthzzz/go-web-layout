package user

import (
	"context"
	userdomain "github.com/lyouthzzz/go-web-layout/internal/domain/user"
)

type Usecase struct {
	user userdomain.UserRepo
}

func NewUserUsecase(user userdomain.UserRepo) *Usecase {
	return &Usecase{user: user}
}

func (uc *Usecase) GetUser(ctx context.Context, uid uint) (*userdomain.User, error) {
	return uc.user.GetUser(ctx, uid)
}

func (uc *Usecase) CreateUser(ctx context.Context, user *userdomain.User) (*userdomain.User, error) {
	return uc.user.CreateUser(ctx, user)
}

func (uc *Usecase) UpdateUser(ctx context.Context, user *userdomain.User) (*userdomain.User, error) {
	return uc.user.UpdateUser(ctx, user)
}

func (uc *Usecase) DeleteUser(ctx context.Context, uid uint) error {
	return uc.user.DeleteUser(ctx, uid)
}
