package block

import "errors"

var (
	ErrMissingBlockType  = errors.New("block type is requred")
	ErrInvalidBlockType  = errors.New("invalid block type")
	ErrMissingBlockValue = errors.New("block value is required")
	ErrDeviceBlocked     = errors.New("device is blocked")
)
