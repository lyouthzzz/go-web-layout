package domain

import "context"

type User struct {
	Name     string
	Password string
	Email    string
}

type UserUsecase interface {
	Get(ctx context.Context, id int64) (User, error)
	Create(ctx context.Context, username, password, email string) (User, error)
	Update(ctx context.Context, id int64, user *User) error
	Delete(ctx context.Context, id int64) error
}

type UserRepository interface {
	Get(ctx context.Context, id int64) (User, error)
	Create(ctx context.Context, username, password, email string) (User, error)
	Update(ctx context.Context, id int64, user *User) error
	Delete(ctx context.Context, id int64) error
}
