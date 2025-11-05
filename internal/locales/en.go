package locales

import "github.com/QuocAnh189/GoCoreFoundation/internal/constants/status"

var (
	EN LanguageType = "en"
)

func GetMessageENFromStatus(statusCode status.Code) string {
	switch statusCode {
	case status.USER_INVALID_PARAMS:
		return "Invalid parameters"
	case status.USER_INVALID_ID:
		return "Invalid user ID"
	case status.USER_NOT_FOUND:
		return "User not found"
	case status.USER_MISSING_FIRST_NAME:
		return "First name is required"
	case status.USER_MISSING_LAST_NAME:
		return "Last name is required"
	case status.USER_MISSING_EMAIL:
		return "Email is required"
	case status.USER_INVALID_EMAIL:
		return "Invalid email format"
	case status.USER_MISSING_PHONE:
		return "Phone is required"
	case status.USER_INVALID_ROLE:
		return "Invalid role"
	case status.USER_INVALID_STATUS:
		return "Invalid status"
	case status.SUCCESS:
		return "Success"
	default:
		return "Unknown error"
	}
}
