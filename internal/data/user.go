package data

import (
	"context"

	"github.com/douyu/jupiter/pkg/xlog"
	"github.com/jinzhu/gorm"
	v1 "github.com/lyouthzzz/go-web-layout/api/user/v1"
	"github.com/lyouthzzz/go-web-layout/internal/data/model"
	"github.com/lyouthzzz/go-web-layout/internal/domain"
	"github.com/lyouthzzz/go-web-layout/pkg/ecode"
)

type UserRepo interface {
	GetUser(ctx context.Context, userId uint64) (*domain.User, error)
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	DeleteUser(ctx context.Context, userId uint64) error
	ListUser(ctx context.Context, pageNum, pageSize uint64) ([]*domain.User, error)
}

type userRepo struct {
	data *Data
	log  *xlog.Logger
}

func NewUserRepo(data *Data, logger *xlog.Logger) UserRepo {
	return &userRepo{data: data, log: logger.With(xlog.String("pkg", "data/user"))}
}

func (r *userRepo) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	return nil, ecode.NotFound(v1.ErrorReason_USER_ALREADY_EXISTS.String(), "")
}

func (r *userRepo) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	return nil, ecode.NotFound(v1.ErrorReason_INTERNALE_ERROR.String(), "")
}

func (r *userRepo) GetUser(ctx context.Context, id uint64) (*domain.User, error) {
	var user model.User
	tx := r.data.db.Where("id = ?", id).First(&user)
	if err := tx.Error; err != nil {
		if gorm.IsRecordNotFoundError(tx.Error) {
			return nil, ecode.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), err.Error())
		}
		return nil, ecode.Internal(v1.ErrorReason_INTERNALE_ERROR.String(), err.Error())
	}
	return &domain.User{Id: user.ID, Username: user.Username, Email: user.Email}, nil
}

func (r *userRepo) DeleteUser(ctx context.Context, userId uint64) error {
	return ecode.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "")
}

func (r *userRepo) ListUser(ctx context.Context, pageNum, pageSize uint64) ([]*domain.User, error) {
	return nil, ecode.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "")
}
