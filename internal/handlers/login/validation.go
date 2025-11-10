package login

import "github.com/QuocAnh189/GoCoreFoundation/internal/constants/status"

func ValidateLoginReq(req *LoginReq) (status.Code, error) {
	if req.LoginName == "" || req.RawPassword == "" {
		return status.LOGIN_MISSING_PARAMETERS, ErrMissingParameters
	}

	if req.DeviceUUID == "" {
		return status.DEVICE_MISSING_UUID, ErrMissingDeviceUUID
	}
	return status.SUCCESS, nil
}
