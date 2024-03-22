package common

import (
	"errors"

	"github.com/jackc/pgx/v5"
)

var (
	ErrDataNotFound         = errors.New("data not found")
	ErrLoginFail            = errors.New("username or password invalid")
	ErrEmailIsAlreadyExists = errors.New("this email is already exists")
	ErrAccountIsInvalid     = errors.New("accounts id is invalid")
	ErrDBNoRows             = pgx.ErrNoRows
)
