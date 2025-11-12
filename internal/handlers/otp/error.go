package otp

import "errors"

var (
	ErrMissingOTPPurpose = errors.New("missing otp purpose")
	ErrInvalidOTPPurpose = errors.New("invalid otp purpose")
	ErrMissingIdentifier = errors.New("missing identifier")
	ErrMissingOTPCode    = errors.New("missing otp code")

	// send otp
	ErrOTPStillActive = errors.New("a verification code was already sent. Please check your email or wait to request a new one")
	ErrExceedMaxSend  = errors.New("you have exceeded the maximum number of verification code requests. Please try again later")

	// verify otp
	ErrOTPInValid      = errors.New("the verification code you entered is incorrect, please try again or request a new one")
	ErrOTPExpired      = errors.New("your code has expired, please request a new one to continue")
	ErrExceedMaxVerify = errors.New("you have exceeded the maximum number of verification attempts. Please request a new code to continue")
	ErrUserBlocked     = errors.New("you will be blocked")

	// valid
	ErrOTPNotAllowed = errors.New("otp action not allowed")

	// block
	ErrDeviceBlocked      = errors.New("device is blocked")
	ErrDevicePhoneBlocked = errors.New("device and phone number are blocked")
	ErrDeviceEmailBlocked = errors.New("device and email are blocked")
)
