package users

import (
	"context"
	"errors"
	"regexp"

	"github.com/QuocAnh189/GoCoreFoundation/internal/constants"
)

// validateUser performs validation on user data.
func (s *UserService) validateUser(ctx context.Context, user *User, isCreate bool) error {
	// Check required fields
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

	// Validate email format
	if !isValidEmail(user.Email) {
		return errors.New("invalid email format: " + user.Email)
	}

	// Validate role (assuming constants.Role is a string enum)
	if !isValidRole(user.Role) {
		return errors.New("invalid role: " + string(user.Role))
	}

	// Add more validation from validation.go if available
	// Example: if err := ValidateUser(user); err != nil { return err }
	return nil
}

// isValidEmail checks if the email format is valid.
func isValidEmail(email string) bool {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

// isValidRole checks if the role is valid (adjust based on constants.Role).
func isValidRole(role constants.Role) bool {
	// Example: Replace with actual valid roles from constants.Role
	validRoles := []constants.Role{"admin", "user", "guest"} // Adjust based on constants.Role
	for _, validRole := range validRoles {
		if role == validRole {
			return true
		}
	}
	return false
}

func isValidStatus(status string) bool {
	validStatuses := []string{"active", "inactive"} // Adjust based on constants/status.go
	for _, validStatus := range validStatuses {
		if status == validStatus {
			return true
		}
	}
	return false
}
