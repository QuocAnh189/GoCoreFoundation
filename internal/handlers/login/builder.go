package login

import (
	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/enum"
	"github.com/QuocAnh189/GoCoreFoundation/internal/utils/uuid"
)

func BuildCreateLoginLogDTO(uid, ipAddress, deviceUUID, token string, loginStatus enum.ELoginStatus) *LoginLogDTO {
	uuid, _ := uuid.GenerateUUIDV7()

	return &LoginLogDTO{
		ID:         uuid,
		UID:        uid,
		IpAddress:  ipAddress,
		DeviceUUID: deviceUUID,
		Token:      token,
		Status:     loginStatus,
	}
}
