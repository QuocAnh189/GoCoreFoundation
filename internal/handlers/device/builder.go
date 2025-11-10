package device

import "github.com/QuocAnh189/GoCoreFoundation/internal/utils/uuid"

func BuildCreateDeviceDTO(req *CreateDeviceReq) *CreateDeviceDTO {
	uuid, _ := uuid.GenerateUUIDV7()

	return &CreateDeviceDTO{
		ID:              uuid,
		UID:             req.UID,
		DeviceUUID:      req.DeviceUUID,
		DeviceName:      req.DeviceName,
		DevicePushToken: req.DevicePushToken,
		IsVerified:      req.IsVerified,
	}
}

func BuildUpdateDeviceDTO(req *UpdateDeviceReq) *UpdateDeviceDTO {
	return &UpdateDeviceDTO{
		ID:              req.ID,
		UID:             req.UID,
		DeviceUUID:      req.DeviceUUID,
		DeviceName:      req.DeviceName,
		DevicePushToken: req.DevicePushToken,
		IsVerified:      req.IsVerified,
	}
}

func MapSchemaToDevice(sd *sqlDevice) *Device {
	if sd == nil {
		return nil
	}

	return &Device{
		ID:              sd.ID,
		UID:             sd.UID.String,
		DeviceUuid:      sd.DeviceUuid,
		DeviceName:      sd.DeviceName,
		DevicePushToken: sd.DevicePushToken,
		IsVerified:      sd.IsVerified,
	}
}
