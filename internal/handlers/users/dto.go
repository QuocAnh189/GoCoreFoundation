package users

import "github.com/QuocAnh189/GoCoreFoundation/internal/constants/enum"

type CreateUserDTO struct {
	FirstName  string     `json:"first_name"`
	MiddleName string     `json:"middle_name"`
	LastName   string     `json:"last_name"`
	Email      string     `json:"email"`
	Phone      string     `json:"phone"`
	Password   string     `json:"password"`
	Role       enum.ERole `json:"role"`
}

type UpdateUserDTO struct {
	ID         string        `json:"id"`
	FirstName  *string       `json:"first_name,omitempty"`
	MiddleName *string       `json:"middle_name,omitempty"`
	LastName   *string       `json:"last_name,omitempty"`
	Email      *string       `json:"email,omitempty"`
	Phone      *string       `json:"phone,omitempty"`
	Role       *enum.ERole   `json:"role,omitempty"`
	Status     *enum.EStatus `json:"status,omitempty"`
}
