package users

import "errors"

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

var (
	ErrInvalidParameter = errors.New("invalid parameter")
	ErrInvalidUserID    = errors.New("invalid user ID")
	ErrUserNotFound     = errors.New("user not found")
)
