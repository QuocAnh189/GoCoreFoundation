package login

import (
	"context"
	"database/sql"
	"time"

	"github.com/QuocAnh189/GoCoreFoundation/internal/db"
)

type IRepository interface {
	GetLoginLogByUIDAndDeviceUUID(ctx context.Context, uid string, deviceUUID string) (*LoginLog, error)
	StoreLoginLog(ctx context.Context, dto *CreateLoginLogDTO) error
	UpdateLoginLog(ctx context.Context, dto *UpdateLoginLogDTO) error
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

func (r *Repository) GetLoginLogByUIDAndDeviceUUID(ctx context.Context, uid string, deviceUUID string) (*LoginLog, error) {
	query := `
		SELECT id, uid, ip_address, device_uuid, token
		FROM login_logs
		WHERE uid = ? AND device_uuid = ?
	`

	var sqlLog sqlLoginLog
	result := r.db.QueryRow(ctx, nil, query, uid, deviceUUID)
	err := result.Scan(&sqlLog.ID, &sqlLog.UID, &sqlLog.IpAddress, &sqlLog.DeviceUUID, &sqlLog.Token)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	loginLog := MapSchemaToLoginLog(&sqlLog)

	return loginLog, nil
}

func (r *Repository) StoreLoginLog(ctx context.Context, dto *CreateLoginLogDTO) error {
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

func (r *Repository) UpdateLoginLog(ctx context.Context, dto *UpdateLoginLogDTO) error {
	query := `
		UPDATE login_logs
		SET uid = COALESCE(?, uid),
			ip_address = COALESCE(?, ip_address),
			device_uuid = COALESCE(?, device_uuid),
			token = COALESCE(?, token),
			status = COALESCE(?, status),
			modify_dt = ?
		WHERE id = ?
	`

	_, err := r.db.Exec(ctx, nil, query,
		dto.UID,
		dto.IpAddress,
		dto.DeviceUUID,
		dto.Token,
		dto.Status,
		time.Now().UTC(),
		dto.ID,
	)

	return err
}
