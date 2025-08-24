package users

import (
	"github.com/QuocAnh189/GoCoreFoundation/internal/constants"
)

// User represents the user model, mapping to the users table.
type User struct {
	ID         int64          `json:"id" db:"id"`
	FirstName  string         `json:"first_name" db:"first_name"`
	MiddleName *string        `json:"middle_name" db:"middle_name"`
	LastName   string         `json:"last_name" db:"last_name"`
	Phone      string         `json:"phone" db:"phone"`
	Email      string         `json:"email" db:"email"`
	Role       constants.Role `json:"role" db:"role"`
	Status     string         `json:"status" db:"status"`
	CreateID   *int           `json:"create_id" db:"create_id"`
	CreateDt   *string        `json:"create_dt" db:"create_dt"`
	ModifyID   *int           `json:"modify_id" db:"modify_id"`
	ModifyDt   *string        `json:"modify_dt" db:"modify_dt"`
} // table:"users"
