package gormx

import (
	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

func IsDuplicateError(err error) bool {
	var mysqlErr mysql.MySQLError
	return errors.As(err, mysqlErr) && mysqlErr.Number == 1062
}
