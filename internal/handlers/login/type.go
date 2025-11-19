package login

import (
	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/enum"
	"github.com/QuocAnh189/GoCoreFoundation/internal/handlers/user"
)

type LoginLog struct {
	ID         string `json:"id"`
	UID        string `json:"uid"`
	IpAddress  string `json:"ip_address"`
	DeviceUUID string `json:"device_uuid"`
	Token      string `json:"token"`
}

type LoginReq struct {
	LoginName   string `json:"login_name"`
	RawPassword string `json:"password"`

	DeviceUUID string `json:"device_uuid,omitempty"`
	DeviceName string `json:"device_name,omitempty"`
	IpAddress  string
}

type LoginRes struct {
	User        *user.User        `json:"user"`
	Needs2FA    bool              `json:"needs_2fa"`
	IsSecure    bool              `json:"is_secure"`
	LoginStatus enum.ELoginStatus `json:"login_status"`
}
