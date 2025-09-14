package users

import (
	"time"

	"github.com/QuocAnh189/GoCoreFoundation/internal/utils/pagination"
)

type User struct {
	ID         string    `json:"id"`
	FirstName  string    `json:"first_name"`
	MiddleName *string   `json:"middle_name"`
	LastName   string    `json:"last_name"`
	Phone      string    `json:"phone"`
	Email      string    `json:"email"`
	Role       Role      `json:"role"`
	Status     string    `json:"status"`
	CreateID   *int64    `json:"create_id"`
	CreateDT   time.Time `json:"create_dt"`
	ModifyID   *int64    `json:"modify_id"`
	ModifyDT   time.Time `json:"modify_dt"`
}

type ListUserRequest struct {
	Search    string `json:"search,omitempty" form:"search"`
	Page      int64  `json:"page" form:"page"`
	Limit     int64  `json:"size" form:"size"`
	OrderBy   string `json:"order_by" form:"order_by"`
	OrderDesc bool   `json:"order_desc" form:"order_desc"`
	TakeAll   bool   `json:"take_all" form:"take_all"`
}

type ListUserResponse struct {
	Users      []*User                `json:"items"`
	Pagination *pagination.Pagination `json:"metadata"`
}

type GetUserResponse struct {
	User *User `json:"result"`
}

type CreateUserRequest struct {
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name,omitempty"`
	LastName   string `json:"last_name"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Role       Role   `json:"role,omitempty"`
}

type CreateUserResponse struct {
	User *User `json:"result"`
}

type UpdateUserRequest struct {
	UID        string  `json:"uid"`
	FirstName  *string `json:"first_name",omitempty"`
	MiddleName *string `json:"middle_name",omitempty"`
	LastName   *string `json:"last_name",omitempty"`
	Phone      *string `json:"phone",omitempty"`
	Email      *string `json:"email",omitempty"`
	Role       *Role   `json:"role",omitempty"`
	Status     *string `json:"status",omitempty"`
}

type UpdateUserResponse struct {
	User *User `json:"result"`
}

type DeleteUserRequest struct {
	UserID string `json:"uid"`
}
