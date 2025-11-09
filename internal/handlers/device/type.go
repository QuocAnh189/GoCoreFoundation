package device

type Device struct {
	ID              string
	UID             string
	DeviceUuid      string
	DeviceName      string
	DevicePushToken string
	IsVerified      bool
}

type CreateDeviceReq struct {
	UID             *string `json:"uid"`
	DeviceUUID      string  `json:"device_uuid"`
	DeviceName      string  `json:"device_name"`
	DevicePushToken *string `json:"device_push_token"`
	IsVerified      bool    `json:"is_verified"`
}

type UpdateDeviceReq struct {
	ID              string  `json:"id"`
	UID             *string `json:"uid"`
	DeviceUUID      *string `json:"device_uuid"`
	DeviceName      *string `json:"device_name"`
	DevicePushToken *string `json:"device_push_token"`
	IsVerified      *bool   `json:"is_verified"`
}
