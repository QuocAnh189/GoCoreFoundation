package login

import "github.com/QuocAnh189/GoCoreFoundation/internal/handlers/users"

type LoginReq struct {
	LoginName   string `json:"login_name"`
	RawPassword string `json:"password"`

	DeviceUUID string `json:"device_uuid,omitempty"`
	DeviceName string `json:"device_name,omitempty"`
}

type LoginRes struct {
	User     *users.User
	Needs2FA bool
	IsSecure bool
}
