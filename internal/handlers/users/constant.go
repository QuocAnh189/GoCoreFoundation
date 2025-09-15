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

var UserErrKeyMap = map[error]string{
	ErrInvalidParameter: "user.invalid_parameter",
	ErrInvalidUserID:    "user.invalid_user_id",
	ErrUserNotFound:     "user.not_found",

	ErrMissingFirstName: "user.first_name_required",
	ErrMissingLastName:  "usre_.last_name_required",
	ErrMissingPhone:     "user.phone_required",
	ErrMissingEmail:     "user.email_required",
	ErrInvalidEmail:     "user.invalid_email_format",
	ErrInvalidRole:      "user.invalid_role",
	ErrInvalidStatus:    "user.invalid_status",
}

var UserErrStatusMap = map[error]status.AppStatusCode{
	ErrInvalidParameter: status.BAD_REQUEST,
	ErrInvalidUserID:    status.BAD_REQUEST,
	ErrUserNotFound:     status.NOT_FOUND,

	ErrMissingFirstName: status.BAD_REQUEST,
	ErrMissingLastName:  status.BAD_REQUEST,
	ErrMissingPhone:     status.BAD_REQUEST,
	ErrMissingEmail:     status.BAD_REQUEST,
	ErrInvalidEmail:     status.BAD_REQUEST,
	ErrInvalidRole:      status.BAD_REQUEST,
	ErrInvalidStatus:    status.BAD_REQUEST,
}
