package lingos

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/QuocAnh189/GoCoreFoundation/internal/db"
	"github.com/QuocAnh189/GoCoreFoundation/internal/utils/uuid"
)

type IRepository interface {
	Insert(ctx context.Context, l Lingo) (*Lingo, error)
	GetByLangAndKey(ctx context.Context, lang Lang, key string) (*Lingo, error)
	DeleteByLangAndKey(ctx context.Context, lang Lang, key string) error
}

type LingoRepository struct {
	db db.IDatabase
}

func NewLingoRepository(db db.IDatabase) *LingoRepository {
	return &LingoRepository{
		db: db,
	}
}

type sqlLingo struct {
	ID       string
	Lang     sql.NullString
	Key      sql.NullString
	Val      sql.NullString
	Status   sql.NullString
	CreateDT sql.NullTime
	CreateID sql.NullInt64
	ModifyDT sql.NullTime
	ModifyID sql.NullInt64
}

func (r *LingoRepository) Insert(ctx context.Context, l Lingo) (*Lingo, error) {
	query := `
		INSERT INTO lingos (id, lang, lkey, lval, status, create_id, create_dt, modify_id, modify_dt)
		VALUES (?, ?, ?, ?, ?, ?, NOW(), ?, NOW())
	`
	if l.ID == "" {
		uuid, err := uuid.GenerateUUIDV7()
		if uuid == "" {
			return nil, fmt.Errorf("failed to generate UUIDv7: %v", err)
		}
		l.ID = uuid
	}

	_, err := r.db.Exec(ctx, query,
		l.ID,
		l.Lang,
		l.Key,
		l.Val,
		l.Status,
		l.CreateID,
		l.ModifyID,
	)
	if err != nil {
		return nil, err
	}

	return r.GetByLangAndKey(ctx, l.Lang, l.Key)
}

func (r *LingoRepository) GetByLangAndKey(ctx context.Context, lang Lang, key string) (*Lingo, error) {
	query := `
		SELECT id, lang, lkey, lval, status, create_dt, modify_dt
		FROM lingos
		WHERE lang = ? AND lkey = ?
	`
	var sl sqlLingo

	result := r.db.QueryRow(ctx, query, lang, key)

	err := result.Scan(
		&sl.ID, &sl.Lang, &sl.Key, &sl.Val, &sl.Status,
		&sl.CreateDT, &sl.ModifyDT,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("scan error: %v", err)
	}

	lingo := &Lingo{
		ID:       sl.ID,
		Lang:     Lang(sl.Lang.String),
		Key:      sl.Key.String,
		Val:      sl.Val.String,
		Status:   sl.Status.String,
		CreateID: sl.CreateID.Int64,
		CreateDT: sl.CreateDT.Time,
		ModifyID: sl.ModifyID.Int64,
		ModifyDT: sl.ModifyDT.Time,
	}

	return lingo, nil
}

func (r *LingoRepository) DeleteByLangAndKey(ctx context.Context, lang Lang, key string) error {
	query := `
		DELETE FROM lingos
		WHERE lang = ? AND lkey = ?
	`
	_, err := r.db.Exec(ctx, query, lang, key)
	return err
}
