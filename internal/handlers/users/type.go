package users

import "github.com/QuocAnh189/GoCoreFoundation/internal/utils/pagination"

type User struct {
	ID         int64   `json:"id" db:"id"`
	FirstName  string  `json:"first_name" db:"first_name"`
	MiddleName *string `json:"middle_name" db:"middle_name"`
	LastName   string  `json:"last_name" db:"last_name"`
	Phone      string  `json:"phone" db:"phone"`
	Email      string  `json:"email" db:"email"`
	Role       Role    `json:"role" db:"role"`
	Status     string  `json:"status" db:"status"`
	CreateID   *int64  `json:"create_id" db:"create_id"`
	CreateDT   *string `json:"create_dt" db:"create_dt"`
	ModifyID   *int64  `json:"modify_id" db:"modify_id"`
	ModifyDT   *string `json:"modify_dt" db:"modify_dt"`
}

type ListUserRequest struct {
	Search    string `json:"search,omitempty" form:"search"`
	Page      int64  `json:"-" form:"page"`
	Limit     int64  `json:"-" form:"size"`
	OrderBy   string `json:"-" form:"order_by"`
	OrderDesc bool   `json:"-" form:"order_desc"`
	TakeAll   bool   `json:"-" form:"take_all"`
}

type ListUserResponse struct {
	Users      []*User                `json:"items"`
	Pagination *pagination.Pagination `json:"metadata"`
}

type GetUserResponse struct {
	User *User `json:"result"`
}

type CreateUserRequest struct {
	FirstName  string `json:"first_name" validate:"required"`
	MiddleName string `json:"middle_name,omitempty"`
	LastName   string `json:"last_name" validate:"required"`
	Phone      string `json:"phone,omitempty"`
	Email      string `json:"email" validate:"required,email"`
	Role       Role   `json:"role,omitempty"`
	Status     string `json:"status,omitempty"`
}

type UpdateUserRequest struct {
	ID         int64  `json:"-"`
	FirstName  string `json:"first_name" validate:"required"`
	MiddleName string `json:"middle_name,omitempty"`
	LastName   string `json:"last_name" validate:"required"`
	Phone      string `json:"phone,omitempty"`
	Email      string `json:"email" validate:"required,email"`
	Role       Role   `json:"role,omitempty"`
	Status     string `json:"status,omitempty"`
}
