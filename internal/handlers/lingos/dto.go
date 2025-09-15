package lingos

type LingoUpdateDTO struct {
	ID     string  `json:"id"`
	Lang   *string `json:"lang,omitempty"`
	Key    *string `json:"key,omitempty"`
	Val    *string `json:"val,omitempty"`
	Status *string `json:"status,omitempty"`
}
