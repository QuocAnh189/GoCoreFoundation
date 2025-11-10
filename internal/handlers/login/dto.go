package login

import "github.com/QuocAnh189/GoCoreFoundation/internal/constants/enum"

type LoginLogDTO struct {
	ID         string            `json:"id"`
	UID        string            `json:"uid"`
	IpAddress  string            `json:"ip_address"`
	DeviceUUID string            `json:"device_uuid"`
	Token      string            `json:"token"`
	Status     enum.ELoginStatus `json:"status"`
}
