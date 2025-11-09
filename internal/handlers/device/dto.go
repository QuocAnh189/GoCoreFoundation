package device

type CreateDeviceDTO struct {
	ID              string
	UID             *string
	DeviceUUID      string
	DeviceName      string
	DevicePushToken *string
	IsVerified      bool
}

type UpdateDeviceDTO struct {
	ID              string
	UID             *string
	DeviceUUID      *string
	DeviceName      *string
	DevicePushToken *string
	IsVerified      *bool
}
