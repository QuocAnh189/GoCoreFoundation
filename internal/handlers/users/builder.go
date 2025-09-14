package users

func BuildCreateUserDTO(req *CreateUserRequest) *CreateUserDTO {
	role := RoleUser
	if req.Role != "" {
		role = req.Role
	}

	return &CreateUserDTO{
		FirstName:  req.FirstName,
		MiddleName: req.MiddleName,
		LastName:   req.LastName,
		Email:      req.Email,
		Phone:      req.Phone,
		Role:       role,
	}
}

func BuildUpdateUserDTO(req *UpdateUserRequest) *UpdateUserDTO {
	return &UpdateUserDTO{
		ID:         req.UID,
		FirstName:  req.FirstName,
		MiddleName: req.MiddleName,
		LastName:   req.LastName,
		Email:      req.Email,
		Phone:      req.Phone,
		Role:       req.Role,
		Status:     req.Status,
	}
}
