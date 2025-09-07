package users

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/QuocAnh189/GoCoreFoundation/internal/db"
	"github.com/QuocAnh189/GoCoreFoundation/internal/utils/pagination"
)

// IRepository defines the user repository interface.
type IRepository interface {
	List(ctx context.Context, req *ListUserRequest) (*ListUserResponse, error)
	FindByID(ctx context.Context, id int64) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, dto *CreateUserDTO) (*User, error)
	Update(ctx context.Context, dto *UpdateUserDTO) error
	Delete(ctx context.Context, id int64) error
}

// UserRepository implements the IRepository interface.
type UserRepository struct {
	db db.IDatabase
}

// NewUserRepository creates a new UserRepository.
func NewUserRepository(db db.IDatabase) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

type sqlUser struct {
	ID         int64
	FirstName  sql.NullString
	MiddleName sql.NullString
	LastName   sql.NullString
	Phone      sql.NullString
	Email      sql.NullString
	Role       sql.NullString
	Status     sql.NullString
	CreateID   sql.NullInt64
	CreateDT   sql.NullString
	ModifyID   sql.NullInt64
	ModifyDT   sql.NullString
}

// List retrieves a paginated list of users with optional search and sorting.
func (r *UserRepository) List(ctx context.Context, req *ListUserRequest) (*ListUserResponse, error) {
	var queryBuilder strings.Builder
	args := []interface{}{}

	// Base query
	queryBuilder.WriteString(`
		SELECT id, first_name, middle_name, last_name, phone, email, 
		role, status, create_id, create_dt, modify_id, modify_dt
		FROM users
	`)

	// Add search condition
	if req.Search != "" {
		queryBuilder.WriteString(` WHERE first_name LIKE ? OR last_name LIKE ? OR email LIKE ?`)
		searchTerm := "%" + req.Search + "%"
		args = append(args, searchTerm, searchTerm, searchTerm)
	}

	// Count total records for pagination
	countQuery := "SELECT COUNT(*) FROM users"
	if req.Search != "" {
		countQuery += ` WHERE first_name LIKE ? OR last_name LIKE ? OR email LIKE ?`
	}
	var total int64
	countRow := r.db.QueryRow(ctx, countQuery, args...)
	if err := countRow.Scan(&total); err != nil {
		return nil, fmt.Errorf("failed to count users: %v", err)
	}

	// Initialize pagination
	pagination := pagination.NewPagination(req.Page, req.Limit, total)
	if req.TakeAll {
		pagination.Size = total
		pagination.Skip = 0
	}

	// Add sorting
	if req.OrderBy != "" {
		queryBuilder.WriteString(fmt.Sprintf(" ORDER BY %s", req.OrderBy))
		if req.OrderDesc {
			queryBuilder.WriteString(" DESC")
		}
	}

	// Add pagination
	if !req.TakeAll {
		queryBuilder.WriteString(` LIMIT ? OFFSET ?`)
		args = append(args, pagination.Size, pagination.Skip)
	}

	// Execute query
	rows, err := r.db.Query(ctx, queryBuilder.String(), args...)
	if err != nil {
		return nil, fmt.Errorf("failed to list users: %v", err)
	}
	defer rows.Close()

	// Scan results
	var users []*User
	for rows.Next() {
		var su sqlUser
		if err := rows.Scan(
			&su.ID, &su.FirstName, &su.MiddleName, &su.LastName, &su.Phone, &su.Email,
			&su.Role, &su.Status, &su.CreateID, &su.CreateDT,
			&su.ModifyID, &su.ModifyDT,
		); err != nil {
			return nil, fmt.Errorf("scan error: %v", err)
		}

		user := &User{
			ID:         su.ID,
			FirstName:  su.FirstName.String,
			MiddleName: &su.MiddleName.String,
			LastName:   su.LastName.String,
			Email:      su.Email.String,
			Phone:      su.Phone.String,
			Status:     su.Status.String,
			Role:       Role(su.Role.String),
			CreateID:   &su.CreateID.Int64,
			CreateDT:   &su.CreateDT.String,
			ModifyID:   &su.ModifyID.Int64,
			ModifyDT:   &su.ModifyDT.String,
		}
		users = append(users, user)
	}

	return &ListUserResponse{
		Users:      users,
		Pagination: pagination,
	}, nil
}

// FindByID retrieves a user by ID.
func (r *UserRepository) FindByID(ctx context.Context, id int64) (*User, error) {
	query := `
		SELECT id, first_name, middle_name, last_name, phone, email, 
		role, status, create_id, create_dt, modify_id, modify_dt
		FROM users
		WHERE id = ?
	`

	result := r.db.QueryRow(ctx, query, id)

	var su sqlUser
	err := result.Scan(
		&su.ID, &su.FirstName, &su.MiddleName, &su.LastName, &su.Phone, &su.Email,
		&su.Role, &su.Status, &su.CreateID, &su.CreateDT, &su.ModifyID, &su.ModifyDT,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("scan error: %v", err)
	}

	user := &User{
		ID:         su.ID,
		FirstName:  su.FirstName.String,
		MiddleName: &su.MiddleName.String,
		LastName:   su.LastName.String,
		Email:      su.Email.String,
		Phone:      su.Phone.String,
		Status:     su.Status.String,
		Role:       Role(su.Role.String),
		CreateID:   &su.CreateID.Int64,
		CreateDT:   &su.CreateDT.String,
		ModifyID:   &su.ModifyID.Int64,
		ModifyDT:   &su.ModifyDT.String,
	}

	return user, nil
}

// FindByEmail retrieves a user by email.
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*User, error) {
	query := `
		SELECT id, first_name, middle_name, last_name, phone, email, 
		role, status, create_id, create_dt, modify_id, modify_dt
		FROM users
		WHERE email = ?
	`
	result := r.db.QueryRow(ctx, query, email)

	var su sqlUser
	err := result.Scan(
		&su.ID, &su.FirstName, &su.MiddleName, &su.LastName, &su.Phone, &su.Email,
		&su.Role, &su.Status, &su.CreateID, &su.CreateDT,
		&su.ModifyID, &su.ModifyDT,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("scan error: %v", err)
	}

	user := &User{
		ID:         su.ID,
		FirstName:  su.FirstName.String,
		MiddleName: &su.MiddleName.String,
		LastName:   su.LastName.String,
		Email:      su.Email.String,
		Phone:      su.Phone.String,
		Status:     su.Status.String,
		Role:       Role(su.Role.String),
		CreateID:   &su.CreateID.Int64,
		CreateDT:   &su.CreateDT.String,
		ModifyID:   &su.ModifyID.Int64,
		ModifyDT:   &su.ModifyDT.String,
	}

	return user, nil
}

// Create inserts a new user into the database.
func (r *UserRepository) Create(ctx context.Context, dto *CreateUserDTO) (*User, error) {
	query := `
		INSERT INTO users (first_name, middle_name, last_name, phone, email, role)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	result, err := r.db.Exec(ctx, query,
		dto.FirstName,
		dto.MiddleName,
		dto.LastName,
		dto.Phone,
		dto.Email,
		dto.Role,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve last insert ID: %v", err)
	}

	return r.FindByID(ctx, id)
}

// Update updates an existing user.
func (r *UserRepository) Update(ctx context.Context, dto *UpdateUserDTO) error {
	query := `
		UPDATE users
		SET first_name = COALESCE(?, first_name),
			middle_name = COALESCE(?, middle_name),
			last_name = COALESCE(?, last_name),
			phone = COALESCE(?, phone),
			email = COALESCE(?, email),
			role = COALESCE(?, role),
			status = COALESCE(?, status),
		WHERE id = ?
	`
	_, err := r.db.Exec(ctx, query,
		dto.FirstName,
		dto.MiddleName,
		dto.LastName,
		dto.Phone,
		dto.Email,
		dto.Role,
		dto.Status,
		dto.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}

	return nil
}

// Delete removes a user by ID.
func (r *UserRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM users WHERE id = ?`
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}
	return nil
}
