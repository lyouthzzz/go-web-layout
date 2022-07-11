package user

import "context"

//go:generate mockgen -source user.go -destination ../../mock/domain/user/user_mock.go -package=user

type UserRepo interface {
	GetUser(context.Context, uint) (*User, error)
	CreateUser(context.Context, *User) (*User, error)
	UpdateUser(context.Context, *User) (*User, error)
	DeleteUser(context.Context, uint) error
}

type User struct {
	UID      uint
	Username string
	Email    string
}
