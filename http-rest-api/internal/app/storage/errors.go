package storage

import "errors"

var (
	ErrRecordNotFound         = errors.New("record not found")
	ErrRecordAlreadyExist     = errors.New("record already exist")
	ErrEmailOrPasswordInvalid = errors.New("email or password is invalid")
	ErrUserUnauthorized       = errors.New("user is unauthorized")
)
