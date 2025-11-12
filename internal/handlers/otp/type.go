package otp

import (
	"time"

	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/enum"
)

type OTP struct {
	ID             string           `json:"id"`
	Purpose        enum.EOTPPurpose `json:"purpose"`
	UID            string           `json:"uid"`
	Identifier     string           `json:"identifier"`
	DeviceUUID     string           `json:"device_uuid"`
	DeviceName     string           `json:"device_name"`
	GenOTPCount    int              `json:"gen_otp_cnt"`
	VerifyOTPCount int              `json:"verify_otp_cnt"`
	OTPCode        string           `json:"otp_code"`
	OTPCreateDt    time.Time        `json:"otp_create_dt"`
	OTPExpireDt    time.Time        `json:"otp_expire_dt"`
	Status         enum.EOTPStatus  `json:"status"`
}

type SendOTPReq struct {
	Purpose        enum.EOTPPurpose `json:"purpose"`
	IdentifierName string           `json:"identifier_name"`
	DeviceUUID     string           `json:"-"`
	DeviceName     string           `json:"-"`
}

type SendOtpRes struct {
	Status        string `json:"status"`
	RemainingTime int64  `json:"remaining_time,omitempty"` // in seconds
	BlockDuration int64  `json:"block_duration,omitempty"` // in seconds
}

type VerifyOTPReq struct {
	Purpose        enum.EOTPPurpose `json:"purpose"`
	IdentifierName string           `json:"identifier_name"`
	OTPCode        string           `json:"otp_code"`
	DeviceUUID     string           `json:"-"`
	DeviceName     string           `json:"-"`
}

type VerifyOTPRes struct {
	Status        string `json:"status"`
	RemainingTime int64  `json:"remaining_time"` // in seconds
}
