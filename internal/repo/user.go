package repo

import (
	"context"
	"github.com/lyouthzzz/go-web-layout/internal/domain"
)

type User struct {
	db    interface{}
	redis interface{}
}

func (u *User) Get(ctx context.Context, id int64) (domain.User, error) {
	panic("implement me")
}

func (u *User) Create(ctx context.Context, username, password, email string) (domain.User, error) {
	panic("implement me")
}

func (u *User) Update(ctx context.Context, id int64, user *domain.User) error {
	panic("implement me")
}

func (u *User) Delete(ctx context.Context, id int64) error {
	panic("implement me")
}
