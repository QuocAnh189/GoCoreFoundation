package login

import (
	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/enum"
	"github.com/QuocAnh189/GoCoreFoundation/internal/utils/uuid"
)

func BuildCreateLoginLogDTO(uid, ipAddress, deviceUUID, token string, loginStatus enum.ELoginStatus) *CreateLoginLogDTO {
	uuid, _ := uuid.GenerateUUIDV7()

	return &CreateLoginLogDTO{
		ID:         uuid,
		UID:        uid,
		IpAddress:  ipAddress,
		DeviceUUID: deviceUUID,
		Token:      token,
		Status:     loginStatus,
	}
}

func BuildUpdateLoginLogDTO(id, uid, ipAddress, deviceUUID, token string, loginStatus enum.ELoginStatus) *UpdateLoginLogDTO {
	return &UpdateLoginLogDTO{
		ID:         id,
		UID:        uid,
		IpAddress:  ipAddress,
		DeviceUUID: deviceUUID,
		Token:      token,
		Status:     loginStatus,
	}
}

func MapSchemaToLoginLog(sl *sqlLoginLog) *LoginLog {
	return &LoginLog{
		ID:         sl.ID,
		UID:        sl.UID,
		IpAddress:  sl.IpAddress,
		DeviceUUID: sl.DeviceUUID,
		Token:      sl.Token,
	}
}
