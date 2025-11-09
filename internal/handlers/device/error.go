package device

import "errors"

var (
	ErrMissingDeviceUUID = errors.New("device UUID is required")
	ErrMissingDeviceName = errors.New("device name is required")
)
