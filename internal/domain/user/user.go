package user

import "context"

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
