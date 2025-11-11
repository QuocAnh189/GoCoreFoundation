package block

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

type IRepository interface {
	ListBlocks(ctx context.Context, req *ListBlockRequest) (*ListBlockResponse, error)
	StoreBlock(ctx context.Context, tx *sql.Tx, dto *CreateBlockDTO) error
	StoreMultipleBlocks(ctx context.Context, handler db.HanderlerWithTx) error
}

type Repository struct {
	db db.IDatabase
}

func NewRepository(db db.IDatabase) IRepository {
	return &Repository{
		db: db,
	}
}

type sqlBlock struct {
	ID             string
	Type           string
	Value          string
	Reason         string
	BlockedDt      time.Time
	BlockedUntilDt *time.Time
}

func (r *Repository) ListBlocks(ctx context.Context, req *ListBlockRequest) (*ListBlockResponse, error) {
	var queryBuilder strings.Builder
	args := []interface{}{}

	// Base query
	queryBuilder.WriteString(`
		SELECT id, type, value, reason, blocked_dt, blocked_until_dt
		FROM blocks WHERE deleted_dt IS NULL
	`)

	// Add search condition
	if req.Search != "" {
		queryBuilder.WriteString(` AND value LIKE ?`)
		searchTerm := "%" + req.Search + "%"
		args = append(args, searchTerm)
	}

	// Count total records for pagination
	countQuery := "SELECT COUNT(*) FROM blocks"
	if req.Search != "" {
		countQuery += ` WHERE value LIKE ? AND deleted_dt IS NULL`
	} else {
		countQuery += ` WHERE deleted_dt IS NULL`
	}
	var total int64
	countRow := r.db.QueryRow(ctx, nil, countQuery, args...)
	if err := countRow.Scan(&total); err != nil {
		return nil, fmt.Errorf("failed to count blocks: %v", err)
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
		return nil, fmt.Errorf("failed to list blocks: %v", err)
	}
	defer rows.Close()

	// Scan results
	var items []*Block
	for rows.Next() {
		var sb sqlBlock
		if err := rows.Scan(
			&sb.ID, &sb.Type, &sb.Value, &sb.Reason, &sb.BlockedDt, &sb.BlockedUntilDt,
		); err != nil {
			return nil, fmt.Errorf("scan error: %v", err)
		}

		items = append(items, MapSchemaToBlock(&sb))
	}

	return &ListBlockResponse{
		Items:      items,
		Pagination: pagination,
	}, nil
}

func (r *Repository) StoreBlock(ctx context.Context, tx *sql.Tx, dto *CreateBlockDTO) error {
	query := `
		INSERT INTO blocks (id, type, value, reason, blocked_dt, blocked_until_dt, status)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`
	_, err := r.db.Exec(ctx, tx, query,
		dto.ID,
		dto.Type,
		dto.Value,
		dto.Reason,
		time.Now().UTC(),
		time.Now().Add(dto.Duration).UTC(),
		enum.StatusActive,
	)
	if err != nil {
		return fmt.Errorf("failed to store block: %v", err)
	}
	return nil
}

func (r *Repository) StoreMultipleBlocks(ctx context.Context, handler db.HanderlerWithTx) error {
	err := r.db.WithTransaction(handler)

	if err != nil {
		return err
	}
	return nil
}
