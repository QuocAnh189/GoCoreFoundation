package users

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/QuocAnh189/GoCoreFoundation/internal/db"
	"github.com/QuocAnh189/GoCoreFoundation/internal/utils/pagination"
	"github.com/QuocAnh189/GoCoreFoundation/internal/utils/uuid"
)

// IRepository defines the user repository interface.
type IRepository interface {
	List(ctx context.Context, req *ListUserRequest) (*ListUserResponse, error)
	FindByID(ctx context.Context, id string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, dto *CreateUserDTO) (*User, error)
	Update(ctx context.Context, dto *UpdateUserDTO) (*User, error)
	Delete(ctx context.Context, id string) error
}

// UserRepository implements the IRepository interface.
type Repository struct {
	db db.IDatabase
}

// NewUserRepository creates a new UserRepository.
func NewRepository(db db.IDatabase) *Repository {
	return &Repository{
		db: db,
	}
}

type sqlUser struct {
	ID         string
	FirstName  sql.NullString
	MiddleName sql.NullString
	LastName   sql.NullString
	Phone      sql.NullString
	Email      sql.NullString
	Role       sql.NullString
	Status     sql.NullString
	CreateID   sql.NullInt64
	CreateDT   sql.NullTime
	ModifyID   sql.NullInt64
	ModifyDT   sql.NullTime
}

// List retrieves a paginated list of users with optional search and sorting.
func (r *Repository) List(ctx context.Context, req *ListUserRequest) (*ListUserResponse, error) {
	var queryBuilder strings.Builder
	args := []interface{}{}

	// Base query
	queryBuilder.WriteString(`
		SELECT id, first_name, middle_name, last_name, phone, email, 
		role, status, create_id, create_dt, modify_id, modify_dt
		FROM users WHERE deleted_dt IS NULL
	`)

	// Add search condition
	if req.Search != "" {
		queryBuilder.WriteString(` AND first_name LIKE ? OR last_name LIKE ? OR email LIKE ?`)
		searchTerm := "%" + req.Search + "%"
		args = append(args, searchTerm, searchTerm, searchTerm)
	}

	// Count total records for pagination
	countQuery := "SELECT COUNT(*) FROM users"
	if req.Search != "" {
		countQuery += ` WHERE first_name LIKE ? OR last_name LIKE ? OR email LIKE ? AND deleted_dt IS NULL`
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
		pagination.Page = 1
		pagination.TotalPages = 1
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
			CreateDT:   su.CreateDT.Time,
			ModifyID:   &su.ModifyID.Int64,
			ModifyDT:   su.ModifyDT.Time,
		}
		users = append(users, user)
	}

	return &ListUserResponse{
		Users:      users,
		Pagination: pagination,
	}, nil
}

// FindByID retrieves a user by ID.
func (r *Repository) FindByID(ctx context.Context, id string) (*User, error) {
	query := `
		SELECT id, first_name, middle_name, last_name, phone, email, 
		role, status, create_id, create_dt, modify_id, modify_dt
		FROM users
		WHERE id = ? AND deleted_dt IS NULL
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
		CreateDT:   su.CreateDT.Time,
		ModifyID:   &su.ModifyID.Int64,
		ModifyDT:   su.ModifyDT.Time,
	}

	return user, nil
}

// FindByEmail retrieves a user by email.
func (r *Repository) FindByEmail(ctx context.Context, email string) (*User, error) {
	query := `
		SELECT id, first_name, middle_name, last_name, phone, email, 
		role, status, create_id, create_dt, modify_id, modify_dt
		FROM users
		WHERE email = ? AND deleted_dt IS NULL
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
		CreateDT:   su.CreateDT.Time,
		ModifyID:   &su.ModifyID.Int64,
		ModifyDT:   su.ModifyDT.Time,
	}

	return user, nil
}

// Create inserts a new user into the database.
func (r *Repository) Create(ctx context.Context, dto *CreateUserDTO) (*User, error) {
	uuid, err := uuid.GenerateUUIDV7()
	if uuid == "" {
		return nil, fmt.Errorf("failed to generate UUIDv7: %v", err)
	}

	query := `
		INSERT INTO users (id, first_name, middle_name, last_name, phone, email, role, status)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err = r.db.Exec(ctx, query,
		uuid,
		dto.FirstName,
		dto.MiddleName,
		dto.LastName,
		dto.Phone,
		dto.Email,
		dto.Role,
		StatusActive,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	return r.FindByID(ctx, uuid)
}

// Update updates an existing user.
func (r *Repository) Update(ctx context.Context, dto *UpdateUserDTO) (*User, error) {
	query := `
		UPDATE users
		SET first_name = COALESCE(?, first_name),
			middle_name = COALESCE(?, middle_name),
			last_name = COALESCE(?, last_name),
			phone = COALESCE(?, phone),
			email = COALESCE(?, email),
			role = COALESCE(?, role),
			status = COALESCE(?, status)
		WHERE id = ? AND deleted_dt IS NULL
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
		return nil, fmt.Errorf("failed to update user: %v", err)
	}

	return r.FindByID(ctx, dto.ID)
}

// Delete removes a user by ID.
func (r *Repository) Delete(ctx context.Context, id string) error {
	query := `
		UPDATE users
		SET deleted_dt = ?
		WHERE id = ?
	`
	_, err := r.db.Exec(ctx, query, time.Now().UTC(), id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}
	return nil
}
