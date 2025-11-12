package otp

import (
	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/enum"
	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/status"
)

func GetArgsByStatatus(statusCode status.Code, sendOTPRes *SendOtpRes, verifyOTPRes *VerifyOTPRes) []interface{} {
	switch statusCode {
	case status.OTP_INVALID_PURPOSE:
		return []interface{}{enum.OTPPurposeLogin2FA, enum.OTPPurposeSignUp, enum.OTPPurposeResetPass}
	case status.OTP_STILL_ACTIVE:
		return []interface{}{sendOTPRes.RemainingTime}
	case status.OTP_BLOCK_DEVICE, status.OTP_BLOCK_DEVICE_EMAIL, status.OTP_BLOCK_DEVICE_PHONE:
		return []interface{}{10}
	case status.OTP_EXCEED_MAX_SEND:
		return []interface{}{"retry_after", 15 * 60}
	case status.OTP_EXCEED_MAX_VERIFY:
		return []interface{}{verifyOTPRes.RemainingTime}
	default:
		return nil
	}
}
