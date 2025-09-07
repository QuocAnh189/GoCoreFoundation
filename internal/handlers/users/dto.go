package users

type CreateUserDTO struct {
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Password   string `json:"password"`
	Role       Role   `json:"role"`
}

type UpdateUserDTO struct {
	ID         int64   `json:"id"`
	FirstName  *string `json:"first_name,omitempty"`
	MiddleName *string `json:"middle_name,omitempty"`
	LastName   *string `json:"last_name,omitempty"`
	Email      *string `json:"email,omitempty"`
	Phone      *string `json:"phone,omitempty"`
	Role       *Role   `json:"role,omitempty"`
	Status     *string `json:"status,omitempty"`
}
