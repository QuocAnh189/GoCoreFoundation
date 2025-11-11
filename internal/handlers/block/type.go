package block

import (
	"time"

	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/enum"
	"github.com/QuocAnh189/GoCoreFoundation/internal/utils/pagination"
)

type Block struct {
	ID             string          `json:"id"`
	Type           enum.EBlockType `json:"type"`
	Value          string          `json:"value"`
	Reason         string          `json:"reason"`
	BlockedDt      time.Time       `json:"blocked_dt"`
	BlockedUntilDt *time.Time      `json:"blocked_until_dt"`
}

type CreateBlockReq struct {
	Type     enum.EBlockType `json:"type"`
	Value    string          `json:"value"`
	Duration time.Duration   `json:"duration"` // in minutes
	Reason   string          `json:"reason"`
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
