package users

import (
	"errors"
)

type Role string
type Status string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
	RoleGuest Role = "guest"

	StatusActive   Status = "ACTIVE"
	StatusInactive Status = "INACTIVE"
	StatusBanned   Status = "BANNED"

	DefaultLang = "vn"
)

var (
	ErrInvalidParameter = errors.New("invalid parameter")
	ErrInvalidUserID    = errors.New("invalid user ID")
	ErrUserNotFound     = errors.New("user not found")

	ErrMissingFirstName = errors.New("first name is required")
	ErrMissingLastName  = errors.New("last name is required")
	ErrMissingPhone     = errors.New("phone is required")
	ErrMissingEmail     = errors.New("email is required")
	ErrInvalidEmail     = errors.New("invalid email format")
	ErrInvalidRole      = errors.New("invalid role")
	ErrInvalidStatus    = errors.New("invalid status")
)
