package block

import (
	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/enum"
	"github.com/QuocAnh189/GoCoreFoundation/internal/utils/uuid"
)

func BuildCeateBlockDTO(req *CreateBlockReq) *CreateBlockDTO {
	id, _ := uuid.GenerateUUIDV7()

	return &CreateBlockDTO{
		ID:     id,
		Type:   req.Type,
		Value:  req.Value,
		Reason: req.Reason,
	}
}

func BuildCreateBlockByValueDTO(blockType enum.EBlockType, value string) *CreateBlockDTO {
	id, _ := uuid.GenerateUUIDV7()

	return &CreateBlockDTO{
		ID:    id,
		Type:  blockType,
		Value: value,
	}
}

func MapSchemaToBlock(sb *sqlBlock) *Block {
	return &Block{
		ID:             sb.ID,
		Type:           enum.EBlockType(sb.Type),
		Value:          sb.Value,
		Reason:         sb.Reason,
		BlockedDt:      sb.BlockedDt,
		BlockedUntilDt: sb.BlockedUntilDt,
	}
}
