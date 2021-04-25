package repo

import (
	"context"
	"fmt"
	"strconv"

	"github.com/lyouthzzz/framework/pkg/gormx"
	"github.com/lyouthzzz/go-web-layout/internal/domain"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.IUserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) GetByName(ctx context.Context, name string) (user domain.User, err error) {
	tx := u.db.WithContext(ctx)
	if err = tx.Where("username = ?", name).First(&user).Error; err != nil {
		return
	}
	return
}

func (u *UserRepository) Get(ctx context.Context, id int64) (user domain.User, err error) {
	tx := u.db.WithContext(ctx)
	if err = tx.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.Wrap(UserNotFoundError, fmt.Sprintf("get user failed. id: %d. \n err: %s \n", id, err.Error()))
		} else {
			err = errors.Wrap(UserUnkonwError, fmt.Sprintf("get user failed. id: %d. \n err: %s \n", id, err.Error()))
		}
		return
	}
	return
}

func (u *UserRepository) Create(ctx context.Context, user *domain.User) (domain.User, error) {
	tx := u.db.WithContext(ctx)
	if err := tx.Create(&user).Error; err != nil {
		if gormx.IsDuplicateError(err) {
			err = UserAlreadyExistsError
		}
		err = errors.Wrap(err, "create user failed")
		return domain.User{}, err
	}
	return *user, nil
}

func (u *UserRepository) Update(ctx context.Context, id int64, user *domain.User) (err error) {
	tx := u.db.WithContext(ctx)
	if err = tx.Where("id = ?", id).Omit("created_at", "updated_at").Updates(&user).Error; err != nil {
		return errors.Wrap(err, "update user failed. id: "+strconv.FormatInt(id, 10))
	}
	if tx.RowsAffected == 0 {
		return errors.Wrap(UserNotFoundError, "update user failed. id: "+strconv.FormatInt(id, 10))
	}
	return
}

func (u *UserRepository) Delete(ctx context.Context, id int64) (err error) {
	tx := u.db.WithContext(ctx)
	if err = tx.Where("id = ?", id).Delete(&domain.User{}).Error; err != nil {
		return errors.Wrap(err, "delete user failed. id: "+strconv.FormatInt(id, 10))
	}
	if tx.RowsAffected == 0 {
		return errors.Wrap(UserNotFoundError, "delete user failed. id: "+strconv.FormatInt(id, 10))
	}
	return
}
