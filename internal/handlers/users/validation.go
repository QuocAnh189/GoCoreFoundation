package users

import (
	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/status"
	"github.com/QuocAnh189/GoCoreFoundation/internal/utils/validate"
)

func ValidateCreateUserRequest(req *CreateUserRequest) (status.Code, error) {
	if req.FirstName == "" {
		return status.USER_MISSING_FIRST_NAME, ErrMissingFirstName
	}

	if req.LastName == "" {
		return status.USER_MISSING_LAST_NAME, ErrMissingLastName
	}

	if req.Phone == "" {
		return status.USER_MISSING_PHONE, ErrMissingPhone

	}
	if req.Email == "" {
		return status.USER_MISSING_EMAIL, ErrMissingEmail
	}

	if req.Password == "" {
		return status.USER_MISSING_PASSWORD, ErrMissingPassword
	}

	if !validate.IsValidEmail(req.Email) {
		return status.USER_INVALID_EMAIL, ErrInvalidEmail
	}

	if !validate.ValidatePhoneNumber(req.Phone) {
		return status.USER_INVALID_PHONE, ErrInvalidPhone
	}

	if req.Role != "" && !validate.IsValidRole(req.Role) {
		return status.USER_INVALID_ROLE, ErrInvalidRole
	}

	return status.SUCCESS, nil
}

func ValidateUpdateUserRequest(req *UpdateUserRequest) (status.Code, error) {
	if req.UID == "" {
		return status.USER_INVALID_ID, ErrInvalidUserID
	}

	if req.Email != nil && !validate.IsValidEmail(*req.Email) {
		return status.USER_INVALID_EMAIL, ErrInvalidEmail
	}

	if req.Role != nil && !validate.IsValidRole(*req.Role) {
		return status.USER_INVALID_ROLE, ErrInvalidRole
	}

	if req.Status != nil && !validate.IsValidStatus(*req.Status) {
		return status.USER_INVALID_STATUS, ErrInvalidStatus
	}

	return status.SUCCESS, nil
}
