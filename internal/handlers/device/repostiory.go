package device

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/enum"
	"github.com/QuocAnh189/GoCoreFoundation/internal/db"
)

type IRepository interface {
	GetDeviceByDeviceUUID(ctx context.Context, deviceUUID string) (*Device, error)
	StoreDevice(ctx context.Context, tx *sql.Tx, dto *CreateDeviceDTO) error
	UpdateDevice(ctx context.Context, dto *UpdateDeviceDTO) error
	DeleteDeviceByUID(ctx context.Context, tx *sql.Tx, uid string) error
}

type Repository struct {
	db db.IDatabase
}

// NewRepository creates a new instance of Repository.
func NewRepository(db db.IDatabase) *Repository {
	return &Repository{
		db: db,
	}
}

type sqlDevice struct {
	ID              string
	UID             sql.NullString
	DeviceUuid      string
	DeviceName      string
	DevicePushToken string
	IsVerified      bool
	Status          sql.NullString
	CreateID        sql.NullInt64
	CreateDT        sql.NullTime
	ModifyID        sql.NullInt64
	ModifyDT        sql.NullTime
}

func (r *Repository) GetDeviceByDeviceUUID(ctx context.Context, deviceUUID string) (*Device, error) {
	query := `
		SELECT id, uid, device_uuid, device_name, device_push_token, is_verified
		FROM devices
		WHERE device_uuid = ? AND status = ?
	`

	result := r.db.QueryRow(ctx, nil, query, deviceUUID, enum.StatusActive)
	var sqlDev sqlDevice
	err := result.Scan(&sqlDev.ID, &sqlDev.UID, &sqlDev.DeviceUuid, &sqlDev.DeviceName, &sqlDev.DevicePushToken, &sqlDev.IsVerified)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	device := MapSchemaToDevice(&sqlDev)

	return device, nil
}

func (r *Repository) StoreDevice(ctx context.Context, tx *sql.Tx, dto *CreateDeviceDTO) error {
	query := `
		INSERT INTO devices (id, uid, device_uuid, device_name, device_push_token, is_verified, status)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	_, err := r.db.Exec(ctx, tx, query,
		dto.ID,
		dto.UID,
		dto.DeviceUUID,
		dto.DeviceName,
		dto.DevicePushToken,
		dto.IsVerified,
		enum.StatusActive,
	)

	return err
}

func (r *Repository) UpdateDevice(ctx context.Context, dto *UpdateDeviceDTO) error {
	query := `
		UPDATE devices
		SET uid = COALESCE(?, uid),
			device_uuid = COALESCE(?, device_uuid),
			device_name = COALESCE(?, device_name),
			device_push_token = COALESCE(?, device_push_token),
			is_verified = COALESCE(?, is_verified)
		WHERE id = ? AND status = ?
	`

	_, err := r.db.Exec(ctx, nil, query,
		dto.UID,
		dto.DeviceUUID,
		dto.DeviceName,
		dto.DevicePushToken,
		dto.IsVerified,
		dto.ID,
		enum.StatusActive,
	)

	return err
}

func (r *Repository) DeleteDeviceByUID(ctx context.Context, tx *sql.Tx, uid string) error {
	query := `
		UPDATE devices
		SET deleted_dt = ?
		WHERE uid = ? AND deleted_dt IS NULL
	`
	_, err := r.db.Exec(ctx, tx, query, time.Now().UTC(), uid)
	if err != nil {
		return fmt.Errorf("failed to delete user logins: %v", err)
	}
	return nil
}
