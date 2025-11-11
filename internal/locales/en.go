package locales

import (
	"fmt"

	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/status"
)

var (
	EN LanguageType = "en"
)

func GetMessageENFromStatus(statusCode status.Code, args ...any) string {
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
	case status.USER_MISSING_PASSWORD:
		return "Password is required"
	case status.USER_INVALID_EMAIL:
		return "Invalid email format"
	case status.USER_EMAIL_ALREADY_EXISTS:
		return "Email already exists"
	case status.USER_MISSING_PHONE:
		return "Phone is required"
	case status.USER_INVALID_PHONE:
		return "Invalid phone format"
	case status.USER_PHONE_ALREADY_EXISTS:
		return "Phone already exists"
	case status.USER_INVALID_ROLE:
		return fmt.Sprintf("Invalid role. Valid roles are: %v", args)
	case status.USER_INVALID_STATUS:
		return fmt.Sprintf("Invalid status. Valid statuses are: %v", args)
	case status.DEVICE_INVALID_PARAMS:
		return "Invalid device parameters"
	case status.DEVICE_MISSING_UUID:
		return "Device UUID is required"
	case status.DEVICE_MISSING_NAME:
		return "Device name is required"
	case status.LOGIN_MISSING_PARAMETERS:
		return "Missing required parameters"
	case status.LOGIN_WRONG_CREDENTIALS:
		return "Wrong login credentials"
	case status.BLOCK_MISSING_TYPE:
		return "Block type is required"
	case status.BLOCK_INVALID_TYPE:
		return fmt.Sprintf("Invalid block type. Valid statuses are: %v", args)
	case status.BLOCK_MISSING_VALUE:
		return "Block value is required"
	case status.SUCCESS:
		return "Success"
	default:
		return "Unknown"
	}
}
