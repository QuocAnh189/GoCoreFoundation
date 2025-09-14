package lingos

import (
	"time"
)

type Lingo struct {
	ID       string    `json:"id"`
	Lang     Lang      `json:"lang"`
	Key      string    `json:"key"`
	Val      string    `json:"val"`
	Status   string    `json:"status"`
	CreateDT time.Time `json:"create_dt"`
	CreateID int64     `json:"create_id"`
	ModifyDT time.Time `json:"modify_dt"`
	ModifyID int64     `json:"modify_id"`
}

type GetLingoRequest struct {
	Lang string `json:"lang" form:"lang"`
	Key  string `json:"key" form:"key"`
}

type GetLingoResponse struct {
	Lingo *Lingo `json:"lingo"`
}

type CreateLingoRequest struct {
	Lang Lang   `json:"lang"`
	Key  string `json:"key"`
}
