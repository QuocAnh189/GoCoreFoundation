package otp

import (
	"time"

	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/enum"
	"github.com/QuocAnh189/GoCoreFoundation/internal/utils/uuid"
)

func BuildCreateOTPDTO(purpose enum.EOTPPurpose, uid *string, identifier, deviceUUID, deviceName, otpCode string, duration time.Duration, otpStatus enum.EOTPStatus) *CreateOTPDTO {
	uuid, _ := uuid.GenerateUUIDV7()
	return &CreateOTPDTO{
		ID:          uuid,
		UID:         uid,
		Purpose:     purpose,
		Identifier:  identifier,
		DeviceUUID:  deviceUUID,
		DeviceName:  deviceName,
		OTPCode:     otpCode,
		OTPCreateDt: time.Now(),
		OTPExpireDt: time.Now().Add(duration),
		GenOTPCount: 1,
		Status:      otpStatus,
	}
}

func BuildUpdateOTPDTO(id string, verifyCount int, status enum.EOTPStatus) *UpdateOTPDTO {
	return &UpdateOTPDTO{
		ID:             id,
		VerifyOTPCount: verifyCount,
		Status:         status,
	}
}

func MapSchemaToOTP(sqlOTP *sqlOTP) *OTP {
	return &OTP{
		ID:             sqlOTP.ID,
		Purpose:        enum.EOTPPurpose(sqlOTP.Purpose),
		Identifier:     sqlOTP.Identifier,
		DeviceUUID:     sqlOTP.DeviceUUID,
		DeviceName:     sqlOTP.DeviceName,
		OTPCode:        sqlOTP.OTPCode,
		OTPCreateDt:    sqlOTP.OTPCreateDt,
		OTPExpireDt:    sqlOTP.OTPExpireDt,
		GenOTPCount:    sqlOTP.GenOTPCount,
		VerifyOTPCount: sqlOTP.VerifyOTPCount,
		Status:         enum.EOTPStatus(sqlOTP.Status),
	}
}
