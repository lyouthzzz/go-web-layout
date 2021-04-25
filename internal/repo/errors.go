package repo

import "github.com/pkg/errors"

var (
	UserNotFoundError      = errors.New("user: not found")
	UserAlreadyExistsError = errors.New("user: already exists")
	UserUnkonwError        = errors.New("user: unknown error")
)
