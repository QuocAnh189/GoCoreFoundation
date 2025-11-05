package users

import (
	"regexp"
)

func ValidateCreateUserRequest(req *CreateUserRequest) error {
	if req.FirstName == "" {
		return ErrMissingFirstName
	}

	if req.LastName == "" {
		return ErrMissingLastName
	}

	if req.Phone == "" {
		return ErrMissingPhone
	}
	if req.Email == "" {
		return ErrMissingEmail
	}

	if !isValidEmail(req.Email) {
		return ErrInvalidEmail
	}

	if req.Role != "" && !isValidRole(req.Role) {
		return ErrInvalidRole
	}

	return nil
}

// validateUser performs validation on user data.
func ValidateUpdateUserRequest(req *UpdateUserRequest) error {
	if req.UID == "" {
		return ErrInvalidUserID
	}

	if req.Email != nil && !isValidEmail(*req.Email) {
		return ErrInvalidEmail
	}

	if req.Role != nil && !isValidRole(*req.Role) {
		return ErrInvalidRole
	}

	if req.Status != nil && !isValidStatus(*req.Status) {
		return ErrInvalidStatus
	}

	return nil
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
