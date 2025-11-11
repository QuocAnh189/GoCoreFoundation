package otp

import (
	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/status"
)

func GetArgsByStatatus(statusCode status.Code) []interface{} {
	switch statusCode {
	case status.OTP_STILL_ACTIVE:
		return []interface{}{"retry_after", 60}
	case status.OTP_EXCEED_MAX_SEND:
		return []interface{}{"retry_after", 15 * 60}
	case status.OTP_EXCEED_MAX_VERIFY:
		return []interface{}{"retry_after", 60}
	default:
		return nil
	}
}
