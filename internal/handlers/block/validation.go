package block

import (
	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/status"
	"github.com/QuocAnh189/GoCoreFoundation/internal/utils/validate"
)

func ValidateCreateBlockReq(req *CreateBlockReq) (status.Code, error) {
	if req.Type == "" {
		return status.BLOCK_MISSING_TYPE, ErrMissingBlockType
	}
	if req.Value == "" {
		return status.BLOCK_MISSING_VALUE, ErrMissingBlockValue
	}

	if !validate.IsValidBlockType(req.Type) {
		return status.BLOCK_INVALID_TYPE, ErrInvalidBlockType
	}

	return status.SUCCESS, nil
}
