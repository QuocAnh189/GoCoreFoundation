package otp

import "errors"

var (
	ErrMissingOTPPurpose = errors.New("missing otp purpose")
	ErrInvalidOTPPurpose = errors.New("invalid otp purpose")
	ErrMissingIdentifier = errors.New("missing identifier")
	ErrMissingOTPCode    = errors.New("missing otp code")
)
