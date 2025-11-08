package validate

import (
	"regexp"

	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/enum"
)

func IsValidEmail(email string) bool {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

func IsValidRole(role enum.ERole) bool {
	validRoles := []enum.ERole{enum.RoleAdmin, enum.RoleUser, enum.RoleGuest}
	for _, validRole := range validRoles {
		if role == validRole {
			return true
		}
	}
	return false
}

func IsValidStatus(status enum.EStatus) bool {
	validStatuses := []enum.EStatus{enum.StatusActive, enum.StatusInactive, enum.StatusBanned}
	for _, validStatus := range validStatuses {
		if status == validStatus {
			return true
		}
	}
	return false
}
