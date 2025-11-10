package users

import (
	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/enum"
	"github.com/QuocAnh189/GoCoreFoundation/internal/utils/uuid"
)

func BuildCreateUserDTO(req *CreateUserRequest) *CreateUserDTO {
	role := enum.RoleUser
	if req.Role != "" {
		role = req.Role
	}

	uuid, _ := uuid.GenerateUUIDV7()

	return &CreateUserDTO{
		ID:         uuid,
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

func MapSchemaToUser(su *sqlUser) *User {
	if su == nil {
		return nil
	}

	return &User{
		ID:         su.ID,
		FirstName:  su.FirstName.String,
		MiddleName: &su.MiddleName.String,
		LastName:   su.LastName.String,
		Email:      su.Email.String,
		Phone:      su.Phone.String,
		Status:     su.Status.String,
		Role:       enum.ERole(su.Role.String),
		CreateID:   &su.CreateID.Int64,
		CreateDT:   su.CreateDT.Time,
		ModifyID:   &su.ModifyID.Int64,
		ModifyDT:   su.ModifyDT.Time,
	}
}

func BuildAliasDTO(uid, aliasValue string) *CreateAliasDTO {
	uuid, _ := uuid.GenerateUUIDV7()

	return &CreateAliasDTO{
		ID:        uuid,
		UID:       uid,
		AliasName: aliasValue,
	}
}

func BuildLoginDTO(uid, password string) *CreateLoginDTO {
	uuid, _ := uuid.GenerateUUIDV7()

	return &CreateLoginDTO{
		ID:       uuid,
		UID:      uid,
		HassPass: password,
	}
}
