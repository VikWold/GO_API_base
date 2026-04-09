package data

import (
	"database/sql"
	"errors"
	"time"
)

var (
	ErrRecordNotFound       = errors.New("recorder not found")
	ErrConstraintViolation  = errors.New("constraint violation")
	ErrUniqueIndexViolation = errors.New("unique index violation")
	ErrUniqueKeyViolation   = errors.New("unique key violation")
)

type Models struct {
	User UserModel
}

func NewModels(db *sql.DB, timeout *time.Duration) Models {
	return Models{
		User: UserModel{DB: db, Timeout: timeout},
	}
}
