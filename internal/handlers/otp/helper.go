package otp

import (
	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/enum"
	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/status"
)

func GetArgsByStatatus(statusCode status.Code) []any {
	switch statusCode {
	case status.OTP_INVALID_PURPOSE:
		return []any{enum.OTPPurposeLogin2FA, enum.OTPPurposeRegistration, enum.OTPPurposeResetPass}
	default:
		return []any{}
	}
}
