package user

import "github.com/QuocAnh189/GoCoreFoundation/internal/constants/enum"

type CreateUserDTO struct {
	ID         string
	FirstName  string
	MiddleName string
	LastName   string
	Email      string
	Phone      string
	Password   string
	Role       enum.ERole
}

type UpdateUserDTO struct {
	ID         string
	FirstName  *string
	MiddleName *string
	LastName   *string
	Email      *string
	Phone      *string
	Role       *enum.ERole
	Status     *enum.EStatus
}

type CreateAliasDTO struct {
	ID        string
	UID       string
	AliasName string
}

type CreateLoginDTO struct {
	ID       string
	UID      string
	HassPass string
}
