package users

import (
	"context"
	"errors"
	"regexp"
)

// validateUser performs validation on user data.
func (s *UserService) validateUser(ctx context.Context, user *User, isCreate bool) error {
	if isCreate && user.ID == 0 {
		return errors.New("ID is required")
	}
	if user.FirstName == "" {
		return errors.New("first_name is required")
	}
	if user.Phone == "" {
		return errors.New("phone is required")
	}
	if user.Email == "" {
		return errors.New("email is required")
	}

	if !isValidEmail(user.Email) {
		return errors.New("invalid email format: " + user.Email)
	}

	if !isValidRole(user.Role) {
		return errors.New("invalid role: " + string(user.Role))
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
	validRoles := []Role{"admin", "user", "guest"}
	for _, validRole := range validRoles {
		if role == validRole {
			return true
		}
	}
	return false
}

func isValidStatus(status string) bool {
	validStatuses := []string{"active", "inactive"}
	for _, validStatus := range validStatuses {
		if status == validStatus {
			return true
		}
	}
	return false
}
