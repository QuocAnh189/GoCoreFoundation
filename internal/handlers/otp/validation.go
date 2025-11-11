package otp

import (
	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/status"
	"github.com/QuocAnh189/GoCoreFoundation/internal/utils/validate"
)

func ValidateSendOTPReq(req *SendOTPReq) (status.Code, error) {
	if req.Purpose == "" {
		return status.OTP_MISSING_PURPOSE, ErrMissingOTPPurpose
	}

	if req.IdentifierName == "" {
		return status.OTP_MISSING_IDENTIFIER, ErrMissingIdentifier
	}

	if !validate.IsValidOTPPurpose(req.Purpose) {
		return status.OTP_INVALID_PURPOSE, ErrInvalidOTPPurpose
	}

	return status.SUCCESS, nil
}

func ValidateVerifyOTPReq(req *VerifyOTPReq) (status.Code, error) {
	if req.Purpose == "" {
		return status.OTP_MISSING_PURPOSE, ErrMissingOTPPurpose
	}

	if req.IdentifierName == "" {
		return status.OTP_MISSING_IDENTIFIER, ErrMissingIdentifier
	}

	if !validate.IsValidOTPPurpose(req.Purpose) {
		return status.OTP_INVALID_PURPOSE, ErrInvalidOTPPurpose
	}

	if req.OTPCode == "" {
		return status.OTP_MISSING_OTP_CODE, ErrMissingOTPCode
	}

	return status.SUCCESS, nil
}
