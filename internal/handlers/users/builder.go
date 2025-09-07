package users

func BuildCreateUserDTO(req *CreateUserRequest) *CreateUserDTO {
	return &CreateUserDTO{
		FirstName:  req.FirstName,
		MiddleName: req.MiddleName,
		LastName:   req.LastName,
		Email:      req.Email,
		Phone:      req.Phone,
		Role:       req.Role,
	}
}

func BuildUpdateUserDTO(req *UpdateUserRequest) *UpdateUserDTO {
	return &UpdateUserDTO{
		ID:         req.ID,
		FirstName:  req.FirstName,
		MiddleName: req.MiddleName,
		LastName:   req.LastName,
		Email:      req.Email,
		Phone:      req.Phone,
		Role:       req.Role,
		Status:     req.Status,
	}
}
