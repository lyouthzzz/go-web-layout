package repo

import (
	"context"
	"errors"
	"github.com/lyouthzzz/go-web-layout/internal/domain"
	"github.com/lyouthzzz/go-web-layout/pkg/gormx"
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
		return
	}
	return
}

func (u *UserRepository) Create(ctx context.Context, user *domain.User) (domain.User, error) {
	tx := u.db.WithContext(ctx)
	if err := tx.Create(&user).Error; err != nil {
		if gormx.IsDuplicateError(err) {
			// todo
		}
		return domain.User{}, err
	}
	return *user, nil
}

func (u *UserRepository) Update(ctx context.Context, id int64, user *domain.User) (err error) {
	tx := u.db.WithContext(ctx)
	if err = tx.Where("id = ?", id).Omit("created_at", "updated_at").Updates(&user).Error; err != nil {
		return
	}
	if tx.RowsAffected == 0 {
		return errors.New("not found")
	}
	return
}

func (u *UserRepository) Delete(ctx context.Context, id int64) (err error) {
	tx := u.db.WithContext(ctx)
	if err = tx.Where("id = ?", id).Delete(&domain.User{}).Error; err != nil {
		return
	}
	if tx.RowsAffected == 0 {
		return errors.New("not found")
	}
	return
}
