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

func IsValidPhoneNumber(phone string) bool {
	r, e := regexp.Compile(`^\+?[\d\s()-]{7,20}$`)
	if e != nil {
		return false
	}

	return r.MatchString(phone)
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

func IsValidBlockType(blockType enum.EBlockType) bool {
	validTypes := []enum.EBlockType{enum.BlockTypeIP, enum.BlockTypeEmail, enum.BlockTypePhone, enum.BlockTypeIP}
	for _, validType := range validTypes {
		if blockType == validType {
			return true
		}
	}
	return false
}

func IsValidOTPPurpose(purpose enum.EOTPPurpose) bool {
	validPurposes := []enum.EOTPPurpose{enum.OTPPurposeLogin2FA, enum.OTPPurposeSignUp, enum.OTPPurposeResetPass}
	for _, validPurpose := range validPurposes {
		if purpose == validPurpose {
			return true
		}
	}
	return false
}
