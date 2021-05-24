package data

import (
	"github.com/jinzhu/gorm"
	"github.com/lyouthzzz/go-web-layout/internal/data/model"
)

type Data struct {
	db *gorm.DB
}

func NewData(db *gorm.DB) (*Data, error) {
	user := model.User{}
	err := db.AutoMigrate(&user).Error
	if err != nil {
		return nil, err
	}
	return &Data{db: db}, nil
}
