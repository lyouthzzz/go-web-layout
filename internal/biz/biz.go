package biz

import (
	"context"
	"github.com/google/wire"
)

/**
biz.go文件有如下作用:
1. 定义Repo层接口
2. 定义三方防腐层接口
3. 维护biz层wire provider set
*/

var ProviderSet = wire.NewSet(NewUserUsecase)

// UserRepo 用户接口
type UserRepo interface {
	GetUser(ctx context.Context, id int64) (*User, error)
	CreateUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, user *User) error
	DeleteUser(ctx context.Context, id int64) error
	ListUser(ctx context.Context, offset, limit int64) (*UserPage, error)
}
