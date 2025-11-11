package block

import (
	"time"

	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/enum"
	"github.com/QuocAnh189/GoCoreFoundation/internal/utils/pagination"
)

type Block struct {
	ID             string
	Type           enum.EBlockType
	Value          string
	Reason         string
	BlockedDt      time.Time
	BlockedUntilDt *time.Time
}

type CreateBlockReq struct {
	Type     enum.EBlockType
	Value    string
	Duration time.Duration
	Reason   string
}

type CreateBlockByValueReq struct {
	Items []CreateBlockReq
}

type ListBlockRequest struct {
	Search    string `json:"search,omitempty" form:"search"`
	Page      int64  `json:"page" form:"page"`
	Limit     int64  `json:"size" form:"size"`
	OrderBy   string `json:"order_by" form:"order_by"`
	OrderDesc bool   `json:"order_desc" form:"order_desc"`
	TakeAll   bool   `json:"take_all" form:"take_all"`
}

type ListBlockResponse struct {
	Items      []*Block               `json:"result"`
	Pagination *pagination.Pagination `json:"metadata"`
}
