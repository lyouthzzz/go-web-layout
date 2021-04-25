package domain

import (
	"context"
	"strconv"

	"github.com/lyouthzzz/framework/pkg/gormx"
)

type User struct {
	gormx.Model
	Username string `gorm:"column:username;type:char(20);not null"`
	Nickname string `gorm:"column:nickname;type:char(20)"`
	Password string `gorm:"column:password;type:char(30);not null"`
	Email    string `gorm:"column:email;type:char(30);uniqueIndex:email_deleted;not null"`

	DeletedAt gormx.DeletedAt `gorm:"column:deleted_at;default:0;uniqueIndex:email_deleted;not null"`
}

var EmptyUser = User{Model: gormx.Model{ID: 0}}

func (User) TableName() string {
	return "user"
}

func (u User) IsEmpty() bool {
	return u.ID == 0
}

func (u User) GetUID() string {
	return strconv.Itoa(int(u.ID))
}

type IUserUsecase interface {
	GetByName(ctx context.Context, name string) (User, error)

	Get(ctx context.Context, id int64) (User, error)
	Create(ctx context.Context, user *User) (User, error)
	Update(ctx context.Context, id int64, user *User) error
	Delete(ctx context.Context, id int64) error
}

type IUserRepository interface {
	GetByName(ctx context.Context, name string) (User, error)

	Get(ctx context.Context, id int64) (User, error)
	Create(ctx context.Context, user *User) (User, error)
	Update(ctx context.Context, id int64, user *User) error
	Delete(ctx context.Context, id int64) error
}
