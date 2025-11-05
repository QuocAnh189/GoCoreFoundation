package users

import (
	"regexp"

	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/status"
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

	if !isValidEmail(req.Email) {
		return status.USER_INVALID_EMAIL, ErrInvalidEmail
	}

	if req.Role != "" && !isValidRole(req.Role) {
		return status.USER_INVALID_ROLE, ErrInvalidRole
	}

	return status.SUCCESS, nil
}

// validateUser performs validation on user data.
func ValidateUpdateUserRequest(req *UpdateUserRequest) (status.Code, error) {
	if req.UID == "" {
		return status.USER_INVALID_ID, ErrInvalidUserID
	}

	if req.Email != nil && !isValidEmail(*req.Email) {
		return status.USER_INVALID_EMAIL, ErrInvalidEmail
	}

	if req.Role != nil && !isValidRole(*req.Role) {
		return status.USER_INVALID_ROLE, ErrInvalidRole
	}

	if req.Status != nil && !isValidStatus(*req.Status) {
		return status.USER_INVALID_STATUS, ErrInvalidStatus
	}

	return status.SUCCESS, nil
}

// isValidEmail checks if the email format is valid.
func isValidEmail(email string) bool {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

func isValidRole(role Role) bool {
	validRoles := []Role{RoleAdmin, RoleUser, RoleGuest}
	for _, validRole := range validRoles {
		if role == validRole {
			return true
		}
	}
	return false
}

func isValidStatus(status Status) bool {
	validStatuses := []Status{StatusActive, StatusInactive, StatusBanned}
	for _, validStatus := range validStatuses {
		if status == validStatus {
			return true
		}
	}
	return false
}
