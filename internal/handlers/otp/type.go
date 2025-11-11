package otp

import "github.com/QuocAnh189/GoCoreFoundation/internal/constants/enum"

type SendOTPReq struct {
	Purpose        enum.EOTPPurpose `json:"purpose"`
	IdentifierName string           `json:"identifier_name"`
	DeviceUUID     string           `json:"-"`
	DeviceName     string           `json:"-"`
}

type SendOtpRes struct {
	Status string `json:"status"`
}

type VerifyOTPReq struct {
	Purpose        enum.EOTPPurpose `json:"purpose"`
	IdentifierName string           `json:"identifier_name"`
	OTPCode        string           `json:"otp_code"`
	DeviceUUID     string           `json:"-"`
	DeviceName     string           `json:"-"`
}

type VerifyOTPRes struct {
	Status string `json:"status"`
}
