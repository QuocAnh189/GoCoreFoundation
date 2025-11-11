package block

import "github.com/QuocAnh189/GoCoreFoundation/internal/constants/enum"

type CreateBlockDTO struct {
	ID       string
	Type     enum.EBlockType
	Value    string
	Reason   string
	Duration int64
}
