package otp

import (
	"context"
	"database/sql"
	"time"

	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/enum"
	"github.com/QuocAnh189/GoCoreFoundation/internal/db"
)

type IRepository interface {
	CreateOTP(ctx context.Context, dto *CreateOTPDTO) error
	GetLatestOTP(ctx context.Context, purpose enum.EOTPPurpose, identifier, deviceUUID string) (*OTP, error)
	UpdateOTP(ctx context.Context, dto *UpdateOTPDTO) error
	InvalidateOldSession(ctx context.Context, purpose enum.EOTPPurpose, identifier, deviceUUID string) error
	CountOTPsInSession(ctx context.Context, purpose enum.EOTPPurpose, identifier, deviceUUID string, sessionStart time.Time) (int, error)
	UpdateOTPForSession(ctx context.Context, dto *CreateOTPDTO) error
	ForceDeleteOTPByStatus(ctx context.Context, status enum.EOTPStatus) error
}

type Repository struct {
	db db.IDatabase
}

func NewRepository(db db.IDatabase) IRepository {
	return &Repository{
		db: db,
	}
}

type sqlOTP struct {
	ID             string
	UID            sql.NullString
	Purpose        string
	Identifier     string
	DeviceUUID     string
	DeviceName     string
	OTPCode        string
	OTPCreateDt    time.Time
	OTPExpireDt    time.Time
	GenOTPCount    int
	VerifyOTPCount int
	Status         string
}

func (r *Repository) CreateOTP(ctx context.Context, dto *CreateOTPDTO) error {
	query := `
		INSERT INTO otp_codes (id, purpose, uid, identifier, device_uuid, device_name, otp_code, otp_create_dt, otp_expire_dt, gen_otp_cnt, verify_otp_cnt, status)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.Exec(ctx, nil, query, dto.ID, dto.Purpose, dto.UID, dto.Identifier, dto.DeviceUUID, dto.DeviceName,
		dto.OTPCode, dto.OTPCreateDt, dto.OTPExpireDt, dto.GenOTPCount, dto.VerifyOTPCount, dto.Status)
	return err
}

func (r *Repository) GetLatestOTP(ctx context.Context, purpose enum.EOTPPurpose, identifier, deviceUUID string) (*OTP, error) {
	query := `
		SELECT id, uid, purpose, identifier, device_uuid, device_name, otp_code, otp_create_dt, otp_expire_dt, gen_otp_cnt, verify_otp_cnt, status
		FROM otp_codes
		WHERE purpose = ? AND identifier = ? AND device_uuid = ? AND status = ?
		ORDER BY otp_create_dt DESC
		LIMIT 1`
	var sqlOTP sqlOTP
	row := r.db.QueryRow(ctx, nil, query, purpose, identifier, deviceUUID, enum.OTPStatusActive)
	err := row.Scan(&sqlOTP.ID, &sqlOTP.UID, &sqlOTP.Purpose, &sqlOTP.Identifier, &sqlOTP.DeviceUUID, &sqlOTP.DeviceName,
		&sqlOTP.OTPCode, &sqlOTP.OTPCreateDt, &sqlOTP.OTPExpireDt, &sqlOTP.GenOTPCount, &sqlOTP.VerifyOTPCount, &sqlOTP.Status)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return MapSchemaToOTP(&sqlOTP), nil
}

func (r *Repository) UpdateOTP(ctx context.Context, dto *UpdateOTPDTO) error {
	query := `
		UPDATE otp_codes
		SET verify_otp_cnt = ?, status = ?
		WHERE id = ?`
	_, err := r.db.Exec(ctx, nil, query, dto.VerifyOTPCount, dto.Status, dto.ID)
	return err
}

func (r *Repository) InvalidateOldSession(ctx context.Context, purpose enum.EOTPPurpose, identifier, deviceUUID string) error {
	query := `
		UPDATE otp_codes
		SET status = ?
		WHERE purpose = ? AND identifier = ? AND device_uuid = ? AND status = ?`
	_, err := r.db.Exec(ctx, nil, query, enum.OTPStatusInactive, purpose, identifier, deviceUUID, enum.OTPStatusActive)
	return err
}

func (r *Repository) CountOTPsInSession(ctx context.Context, purpose enum.EOTPPurpose, identifier, deviceUUID string, sessionStart time.Time) (int, error) {
	query := `
		SELECT gen_otp_cnt
		FROM otp_codes
		WHERE purpose = ? AND identifier = ? AND device_uuid = ? AND otp_create_dt >= ? AND status = ?`
	var count int
	row := r.db.QueryRow(ctx, nil, query, purpose, identifier, deviceUUID, sessionStart, enum.OTPStatusActive)
	err := row.Scan(&count)
	if err == sql.ErrNoRows {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *Repository) UpdateOTPForSession(ctx context.Context, dto *CreateOTPDTO) error {
	query := `
		UPDATE otp_codes
		SET otp_code = ?, otp_create_dt = ?, otp_expire_dt = ?, gen_otp_cnt = ?, verify_otp_cnt = ?, status = ?
		WHERE id = ?`
	_, err := r.db.Exec(ctx, nil, query, dto.OTPCode, dto.OTPCreateDt, dto.OTPExpireDt, dto.GenOTPCount,
		dto.VerifyOTPCount, dto.Status, dto.ID)
	return err
}

func (r *Repository) ForceDeleteOTPByStatus(ctx context.Context, status enum.EOTPStatus) error {
	query := `
		DELETE FROM otp_codes
		WHERE status = ?`
	_, err := r.db.Exec(ctx, nil, query, status)
	return err
}
