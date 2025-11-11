package otp

import (
	"time"

	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/enum"
)

type CreateOTPDTO struct {
	ID             string
	Purpose        enum.EOTPPurpose
	UID            *string
	Identifier     string
	DeviceUUID     string
	DeviceName     string
	OTPCode        string
	OTPCreateDt    time.Time
	OTPExpireDt    time.Time
	GenOTPCount    int
	VerifyOTPCount int
	Status         enum.EOTPStatus
}

type UpdateOTPDTO struct {
	ID             string
	VerifyOTPCount int
	Status         enum.EOTPStatus
}
