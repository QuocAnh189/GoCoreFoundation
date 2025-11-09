package device

import "github.com/QuocAnh189/GoCoreFoundation/internal/constants/status"

func ValidationCreateDeviceReq(req *CreateDeviceReq) (status.Code, error) {
	if req.DeviceUUID == "" {
		return status.DEVICE_MISSING_UUID, ErrMissingDeviceUUID
	}

	if req.DeviceName == "" {
		return status.DEVICE_MISSING_NAME, ErrMissingDeviceName
	}

	return status.SUCCESS, nil
}

func ValidationUpdateDeviceReq(req *UpdateDeviceReq) (status.Code, error) {
	if req.DeviceUUID == nil {
		return status.DEVICE_MISSING_UUID, ErrMissingDeviceUUID
	}

	if req.DeviceName == nil {
		return status.DEVICE_MISSING_NAME, ErrMissingDeviceName
	}

	return status.SUCCESS, nil
}
