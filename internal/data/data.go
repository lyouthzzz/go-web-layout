package data

import (
	"github.com/google/wire"
	"github.com/lyouthzzz/go-web-layout/internal/conf"
	"github.com/lyouthzzz/go-web-layout/internal/data/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(NewData, NewUserRepo)

type Data struct {
	db *gorm.DB
}

func NewData(conf *conf.Data) (*Data, func(), error) {
	db, err := gorm.Open(mysql.Open(conf.Database.Source))
	if err != nil {
		panic(err)
	}
	if err = db.AutoMigrate(&model.User{}); err != nil {
		panic(err)
	}
	return &Data{db: db}, func() {}, nil
}
