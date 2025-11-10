package login

import (
	"context"

	"github.com/QuocAnh189/GoCoreFoundation/internal/db"
)

type IRepository interface {
	StoreLoginLog(ctx context.Context, dto *LoginLogDTO) error
}

type Repository struct {
	db db.IDatabase
}

func NewRepository(db db.IDatabase) *Repository {
	return &Repository{
		db: db,
	}
}

type sqlLoginLog struct {
	ID         string `db:"id"`
	UID        string `db:"uid"`
	IpAddress  string `db:"ip_address"`
	DeviceUUID string `db:"device_uuid"`
	Token      string `db:"token"`
}

func (r *Repository) StoreLoginLog(ctx context.Context, dto *LoginLogDTO) error {
	query := `
		INSERT INTO login_logs (id, uid, ip_address, device_uuid, token, status)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	_, err := r.db.Exec(ctx, nil, query,
		dto.ID,
		dto.UID,
		dto.IpAddress,
		dto.DeviceUUID,
		dto.Token,
		dto.Status,
	)

	return err
}
