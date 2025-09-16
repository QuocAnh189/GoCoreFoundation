package users

import (
	"errors"

	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/status"
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

func DetermineErrKey(err error) string {
	switch err {
	case ErrInvalidParameter:
		return "user.invalid_parameter"
	case ErrInvalidUserID:
		return "user.invalid_user_id"
	case ErrUserNotFound:
		return "user.not_found"
	case ErrMissingFirstName:
		return "user.first_name_required"
	case ErrMissingLastName:
		return "user.last_name_required"
	case ErrMissingPhone:
		return "user.phone_required"
	case ErrMissingEmail:
		return "user.email_required"
	case ErrInvalidEmail:
		return "user.invalid_email_format"
	case ErrInvalidRole:
		return "user.invalid_role"
	case ErrInvalidStatus:
		return "user.invalid_status"
	default:
		return "user.unknown_error"
	}
}

func DetermineErrStatus(err error) int {
	switch err {
	case ErrInvalidParameter, ErrInvalidUserID, ErrMissingFirstName,
		ErrMissingLastName, ErrMissingPhone, ErrMissingEmail,
		ErrInvalidEmail, ErrInvalidRole, ErrInvalidStatus:
		return status.BAD_REQUEST
	case ErrUserNotFound:
		return status.NOT_FOUND
	default:
		return status.INTERNAL
	}
}
