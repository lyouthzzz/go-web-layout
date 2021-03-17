package tests

import (
	"github.com/lyouthzzz/framework/pkg/gormx"
	"github.com/lyouthzzz/go-web-layout/internal/domain"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error

	if db, err = gormx.Connect(gormx.MysqlDns("root", "root", "172.81.212.232", 3306, "web_layout", ""), true); err != nil {
		panic(err)
	}

	if err = db.Migrator().AutoMigrate(domain.User{}); err != nil {
		panic(err)
	}
}
