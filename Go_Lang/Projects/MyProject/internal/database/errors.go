package database

import (
	"errors"

	"github.com/go-sql-driver/mysql"
)

func IsDuplicateError(err error) bool {
	if err == nil {
		return false
	}

	var mysqlError *mysql.MySQLError
	errors.As(err, &mysqlError)
	if mysqlError.Number == 1062 {
		return true
	}
	return false
}