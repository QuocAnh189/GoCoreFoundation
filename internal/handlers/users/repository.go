package users

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/enum"
	"github.com/QuocAnh189/GoCoreFoundation/internal/db"
	"github.com/QuocAnh189/GoCoreFoundation/internal/utils/pagination"
)

// IRepository defines the user repository interface.
type IRepository interface {
	CreateUserWithAssociations(ctx context.Context, handler db.HanderlerWithTx, uid string) (*User, error)
	DeleteUserWithAssociations(ctx context.Context, handler db.HanderlerWithTx) error
	ForceDeleteUserWithAssociations(ctx context.Context, handler db.HanderlerWithTx) error

	GetUserByLoginName(ctx context.Context, loginName string) (*User, error)

	// users
	List(ctx context.Context, req *ListUserRequest) (*ListUserResponse, error)
	FindByID(ctx context.Context, id string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, tx *sql.Tx, dto *CreateUserDTO) (int64, error) // Add tx parameter
	Update(ctx context.Context, dto *UpdateUserDTO) (int64, error)
	Delete(ctx context.Context, id string) error
	ForceDelete(ctx context.Context, tx *sql.Tx, id string) error

	// aliases
	StoreUserAlias(ctx context.Context, tx *sql.Tx, dto *CreateAliasDTO) error // Add tx parameter
	DeleteUserAlias(ctx context.Context, uid string) error
	ForceDeleteUserAlias(ctx context.Context, tx *sql.Tx, uid string) error

	// logins
	StoreLogin(ctx context.Context, tx *sql.Tx, dto *CreateLoginDTO) error // Add tx parameter
	DeleteLogin(ctx context.Context, uid string) error
	ForceDeleteLogin(ctx context.Context, tx *sql.Tx, uid string) error
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
	ID           string
	FirstName    sql.NullString
	MiddleName   sql.NullString
	LastName     sql.NullString
	Phone        sql.NullString
	Email        sql.NullString
	Role         sql.NullString
	Status       sql.NullString
	HashPassword string
	CreateID     sql.NullInt64
	CreateDT     sql.NullTime
	ModifyID     sql.NullInt64
	ModifyDT     sql.NullTime
}

// CreateWithAliases creates a user and their aliases in a single transaction.
func (r *Repository) CreateUserWithAssociations(ctx context.Context, handler db.HanderlerWithTx, uid string) (*User, error) {
	err := r.db.WithTransaction(handler)

	if err != nil {
		return nil, err
	}

	result, err := r.FindByID(ctx, uid)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteUserWithAssociations deletes a user and their associated records in a single transaction.
func (r *Repository) DeleteUserWithAssociations(ctx context.Context, handler db.HanderlerWithTx) error {
	err := r.db.WithTransaction(handler)
	if err != nil {
		return err
	}

	return nil
}

// ForceDeleteUserWithAssociations permanently deletes a user and their associated records in a single transaction.
func (r *Repository) ForceDeleteUserWithAssociations(ctx context.Context, handler db.HanderlerWithTx) error {
	err := r.db.WithTransaction(handler)
	if err != nil {
		return err
	}

	return nil
}

// GetUserByLoginName retrieves a user by their login name (email or phone).
func (r *Repository) GetUserByLoginName(ctx context.Context, loginName string) (*User, error) {
	query := `
		SELECT u.id, u.first_name, u.middle_name, u.last_name, u.phone, u.email, 
		u.role, u.status, l.hash_pass, u.create_id, u.create_dt, u.modify_id, u.modify_dt
		FROM users u
		JOIN aliases a ON u.id = a.uid
		JOIN logins l ON u.id = l.uid
		WHERE a.aka = ? AND u.deleted_dt IS NULL AND a.deleted_dt IS NULL AND l.deleted_dt IS NULL
	`
	result := r.db.QueryRow(ctx, nil, query, loginName)

	var su sqlUser
	err := result.Scan(
		&su.ID, &su.FirstName, &su.MiddleName, &su.LastName, &su.Phone, &su.Email,
		&su.Role, &su.Status, &su.HashPassword, &su.CreateID, &su.CreateDT,
		&su.ModifyID, &su.ModifyDT,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("scan error: %v", err)
	}

	user := MapSchemaToUser(&su)

	return user, nil
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
	} else {
		countQuery += ` WHERE deleted_dt IS NULL`
	}
	var total int64
	countRow := r.db.QueryRow(ctx, nil, countQuery, args...)
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
	rows, err := r.db.Query(ctx, nil, queryBuilder.String(), args...)
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

		users = append(users, MapSchemaToUser(&su))
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

	result := r.db.QueryRow(ctx, nil, query, id)

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

	user := MapSchemaToUser(&su)

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
	result := r.db.QueryRow(ctx, nil, query, email)

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

	user := MapSchemaToUser(&su)

	return user, nil
}

// Create inserts a new user into the database.
func (r *Repository) Create(ctx context.Context, tx *sql.Tx, dto *CreateUserDTO) (int64, error) {
	query := `
		INSERT INTO users (id, first_name, middle_name, last_name, phone, email, role, status)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`
	result, err := r.db.Exec(ctx, tx, query,
		dto.ID,
		dto.FirstName,
		dto.MiddleName,
		dto.LastName,
		dto.Phone,
		dto.Email,
		dto.Role,
		enum.StatusActive,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to create user: %v", err)
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to retrieve last insert ID: %v", err)
	}

	return insertedID, nil
}

// Update updates an existing user.
func (r *Repository) Update(ctx context.Context, dto *UpdateUserDTO) (int64, error) {
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
	result, err := r.db.Exec(ctx, nil, query,
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
		return 0, fmt.Errorf("failed to update user: %v", err)
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("failed to retrieve affected rows: %v", err)
	}

	return affectedRows, nil
}

// Delete removes a user by ID.
func (r *Repository) Delete(ctx context.Context, uid string) error {
	query := `
		UPDATE users
		SET deleted_dt = ?
		WHERE id = ?
	`
	_, err := r.db.Exec(ctx, nil, query, time.Now().UTC(), uid)
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}
	return nil
}

// ForceDelete removes a user by ID permanently.
func (r *Repository) ForceDelete(ctx context.Context, tx *sql.Tx, uid string) error {
	query := `
		DELETE FROM users
		WHERE id = ?
	`
	_, err := r.db.Exec(ctx, tx, query, uid)
	if err != nil {
		return fmt.Errorf("failed to force delete user: %v", err)
	}
	return nil
}

// StoreUserAlias stores a user alias in the database.
func (r *Repository) StoreUserAlias(ctx context.Context, tx *sql.Tx, dto *CreateAliasDTO) error {
	query := `
		INSERT INTO aliases (id, uid, aka, status)
		VALUES (?, ?, ?, ?)
	`
	_, err := r.db.Exec(ctx, tx, query,
		dto.ID,
		dto.UID,
		dto.AliasName,
		enum.StatusActive,
	)
	if err != nil {
		return fmt.Errorf("failed to store user alias: %v", err)
	}
	return nil
}

// DeleteUserAlias deletes user aliases by user ID.
func (r *Repository) DeleteUserAlias(ctx context.Context, uid string) error {
	query := `
		UPDATE aliases
		SET deleted_dt = ?
		WHERE uid = ? AND deleted_dt IS NULL
	`
	_, err := r.db.Exec(ctx, nil, query, time.Now().UTC(), uid)
	if err != nil {
		return fmt.Errorf("failed to delete user aliases: %v", err)
	}
	return nil
}

// ForceDeleteUserAlias permanently deletes user aliases by user ID.
func (r *Repository) ForceDeleteUserAlias(ctx context.Context, tx *sql.Tx, uid string) error {
	query := `
		DELETE FROM aliases
		WHERE uid = ?
	`
	_, err := r.db.Exec(ctx, tx, query, uid)
	if err != nil {
		return fmt.Errorf("failed to force delete user aliases: %v", err)
	}
	return nil
}

// StoreLogin stores a user login record in the database.
func (r *Repository) StoreLogin(ctx context.Context, tx *sql.Tx, dto *CreateLoginDTO) error {
	query := `
		INSERT INTO logins (id, uid, hash_pass, status)
		VALUES (?, ?, ?, ?)
	`
	_, err := r.db.Exec(ctx, tx, query,
		dto.ID,
		dto.UID,
		dto.HassPass,
		enum.StatusActive,
	)
	if err != nil {
		return fmt.Errorf("failed to store user login: %v", err)
	}
	return nil
}

// DeleteLogin deletes user logins by user ID.
func (r *Repository) DeleteLogin(ctx context.Context, uid string) error {
	query := `
		UPDATE logins
		SET deleted_dt = ?
		WHERE uid = ? AND deleted_dt IS NULL
	`
	_, err := r.db.Exec(ctx, nil, query, time.Now().UTC(), uid)
	if err != nil {
		return fmt.Errorf("failed to delete user logins: %v", err)
	}
	return nil
}

// ForceDeleteLogin permanently deletes user logins by user ID.
func (r *Repository) ForceDeleteLogin(ctx context.Context, tx *sql.Tx, uid string) error {
	query := `
		DELETE FROM logins
		WHERE uid = ?
	`
	_, err := r.db.Exec(ctx, tx, query, uid)
	if err != nil {
		return fmt.Errorf("failed to force delete user logins: %v", err)
	}
	return nil
}
