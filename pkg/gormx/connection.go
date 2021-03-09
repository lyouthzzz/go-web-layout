package gormx

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlDns(user, pwd, host string, port int, dbName string, extra string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&%s",
		user, pwd, host, port, dbName, extra)
}

func connectMysql(mysqlDns string, config *gorm.Config) (db *gorm.DB, err error) {
	return gorm.Open(mysql.Open(mysqlDns), config)
}

func Connect(mysqlDns string, debug bool) (*gorm.DB, error) {
	db, err := connectMysql(mysqlDns, &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if debug {
		db = db.Debug()
	}
	return db, nil
}
