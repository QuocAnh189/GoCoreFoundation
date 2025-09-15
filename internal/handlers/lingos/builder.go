package lingos

func BuildUpdateLingoDTO(req *UpdateLingoRequest) *LingoUpdateDTO {
	return &LingoUpdateDTO{
		ID:     req.ID,
		Lang:   req.Lang,
		Key:    req.Key,
		Val:    req.Val,
		Status: req.Status,
	}
}
