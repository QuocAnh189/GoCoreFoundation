package otp

import (
	"context"
	"fmt"
	"time"

	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/enum"
	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/status"
	"github.com/QuocAnh189/GoCoreFoundation/internal/handlers/block"
	"github.com/QuocAnh189/GoCoreFoundation/internal/handlers/device"
	"github.com/QuocAnh189/GoCoreFoundation/internal/handlers/users"
	"github.com/QuocAnh189/GoCoreFoundation/internal/services/mail"
	"github.com/QuocAnh189/GoCoreFoundation/internal/services/sms"
	"github.com/QuocAnh189/GoCoreFoundation/internal/utils/random"
	"github.com/QuocAnh189/GoCoreFoundation/internal/utils/validate"
)

const (
	MaxOTPsPerSession   = 3
	MaxVerifyAttempts   = 3
	OTPSessionDuration  = 15 * time.Minute
	OTPValidityDuration = 1 * time.Minute
)

type Service struct {
	repo      IRepository
	userSvc   *users.Service
	blockSvc  *block.Service
	deviceSvc *device.Service
	mailSvc   *mail.Service
	smsSvc    *sms.Service
}

func NewService(
	repo IRepository,
	userSvc *users.Service,
	blockSvc *block.Service,
	deviceSvc *device.Service,
	mailSvc *mail.Service,
	smsSvc *sms.Service,
) *Service {
	return &Service{
		repo:      repo,
		userSvc:   userSvc,
		blockSvc:  blockSvc,
		deviceSvc: deviceSvc,
		mailSvc:   mailSvc,
		smsSvc:    smsSvc,
	}
}

func (s *Service) SendOTP(ctx context.Context, req *SendOTPReq) (status.Code, *SendOtpRes, error) {
	var (
		UID *string
	)
	// Validate request
	if statusCode, err := ValidateSendOTPReq(req); err != nil {
		return statusCode, nil, err
	}

	// get user by identifier
	statusCode, user, err := s.userSvc.GetUserByLoginName(ctx, req.IdentifierName)
	if err != nil {
		return statusCode, nil, err
	}
	if user != nil {
		UID = &user.ID
	}

	// Check if sending OTP is allowed
	if !IsAllowedSendOTP(req.Purpose, UID) {
		return status.OTP_NOT_ALLOWED, nil, ErrOTPNotAllowed
	}

	// Check session limits
	sessionStart := time.Now().Add(-OTPSessionDuration)
	count, err := s.repo.CountOTPsInSession(ctx, req.Purpose, req.IdentifierName, req.DeviceUUID, sessionStart)
	if err != nil {
		return status.INTERNAL, nil, err
	}
	if count >= MaxOTPsPerSession {
		// Invalidate the session by setting status to INACTIVE
		latestOTP, err := s.repo.GetLatestOTP(ctx, req.Purpose, req.IdentifierName, req.DeviceUUID)
		if err != nil {
			return status.INTERNAL, nil, err
		}
		if latestOTP != nil {
			updateDTO := BuildUpdateOTPDTO(latestOTP.ID, latestOTP.VerifyOTPCount, enum.OTPStatusInactive)
			if err := s.repo.UpdateOTP(ctx, updateDTO); err != nil {
				return status.INTERNAL, nil, err
			}
		}

		statusCode, err := s.CreateBlock(ctx, latestOTP, "exceed max send otp attempts", time.Minute*10)
		if err != nil {
			return statusCode, nil, err
		}

		return status.OTP_EXCEED_MAX_SEND, nil, ErrExceedMaxSend
	}

	// Check for active OTP
	latestOTP, err := s.repo.GetLatestOTP(ctx, req.Purpose, req.IdentifierName, req.DeviceUUID)
	if err != nil {
		return status.INTERNAL, nil, err
	}
	if latestOTP != nil && time.Now().Before(latestOTP.OTPExpireDt) {
		return status.OTP_STILL_ACTIVE, nil, ErrOTPStillActive
	}

	// Generate OTP code
	otpCode := random.GenerateOTPWithLength(4)

	// send OTP to user
	if validate.IsValidEmail(req.IdentifierName) {
		subject := "Your OTP Code"
		body := fmt.Sprintf("Your OTP code is: %s. It is valid for %d minutes.", otpCode, int(OTPValidityDuration.Minutes()))
		if err := s.mailSvc.Send(req.IdentifierName, subject, body, false); err != nil {
			return status.INTERNAL, nil, err
		}
	} else if validate.IsValidPhoneNumber(req.IdentifierName) {
		body := fmt.Sprintf("Your OTP code is: %s. It is valid for %d minutes.", otpCode, int(OTPValidityDuration.Minutes()))
		if err := s.smsSvc.SendSmsToPhone(req.IdentifierName, body); err != nil {
			return status.INTERNAL, nil, err
		}
	}

	// Handle session row
	if latestOTP != nil && time.Now().Before(latestOTP.OTPCreateDt.Add(OTPSessionDuration)) {
		// Update existing session row
		dto := BuildCreateOTPDTO(req.Purpose, UID, req.IdentifierName, req.DeviceUUID, req.DeviceName, otpCode, OTPValidityDuration, enum.OTPStatusActive)
		dto.ID = latestOTP.ID
		dto.GenOTPCount = latestOTP.GenOTPCount + 1
		dto.VerifyOTPCount = 0 // Reset verification count for new OTP
		if err := s.repo.UpdateOTPForSession(ctx, dto); err != nil {
			return status.INTERNAL, nil, err
		}
	} else {
		dto := BuildCreateOTPDTO(req.Purpose, UID, req.IdentifierName, req.DeviceUUID, req.DeviceName, otpCode, OTPValidityDuration, enum.OTPStatusActive)
		if err := s.repo.CreateOTP(ctx, dto); err != nil {
			return status.INTERNAL, nil, err
		}
	}

	// TODO: Send OTP to user (e.g., via email or SMS)
	return status.SUCCESS, &SendOtpRes{Status: "sent otp successfully"}, nil
}

func (s *Service) VerifyOTP(ctx context.Context, req *VerifyOTPReq) (status.Code, *VerifyOTPRes, error) {
	// Validate request
	if statusCode, err := ValidateVerifyOTPReq(req); err != nil {
		return statusCode, nil, err
	}

	// Get latest OTP
	otp, err := s.repo.GetLatestOTP(ctx, req.Purpose, req.IdentifierName, req.DeviceUUID)
	if err != nil {
		return status.INTERNAL, nil, err
	}
	if otp == nil {
		return status.OTP_INVALID_CODE, nil, ErrOTPInValid
	}

	// Check OTP expiration
	if time.Now().After(otp.OTPExpireDt) {
		return status.OTP_EXPIRED, nil, ErrOTPExpired
	}

	// Check verification attempts
	if otp.VerifyOTPCount >= MaxVerifyAttempts {
		// Check if this is the last OTP in the session
		sessionStart := time.Now().Add(-OTPSessionDuration)
		count, err := s.repo.CountOTPsInSession(ctx, req.Purpose, req.IdentifierName, req.DeviceUUID, sessionStart)
		if err != nil {
			return status.INTERNAL, nil, err
		}
		if count >= MaxOTPsPerSession-1 {
			updateDTO := BuildUpdateOTPDTO(otp.ID, otp.VerifyOTPCount, enum.OTPStatusInactive)
			if err := s.repo.UpdateOTP(ctx, updateDTO); err != nil {
				return status.INTERNAL, nil, err
			}

			statusCode, err := s.CreateBlock(ctx, otp, "exceed max verify otp attempts", time.Minute*10)
			if err != nil {
				return statusCode, nil, err
			}

		}
		return status.OTP_EXCEED_MAX_VERIFY, nil, ErrExceedMaxVerify
	}

	// Verify OTP code
	if otp.OTPCode != req.OTPCode {
		updateDTO := BuildUpdateOTPDTO(otp.ID, otp.VerifyOTPCount+1, otp.Status)
		if err := s.repo.UpdateOTP(ctx, updateDTO); err != nil {
			return status.INTERNAL, nil, err
		}
		return status.OTP_INVALID_CODE, nil, ErrOTPInValid
	}

	// OTP verified successfully
	updateDTO := BuildUpdateOTPDTO(otp.ID, otp.VerifyOTPCount+1, enum.OTPStatusUsed)
	if err := s.repo.UpdateOTP(ctx, updateDTO); err != nil {
		return status.INTERNAL, nil, err
	}

	// update device as verified
	if req.Purpose == enum.OTPPurposeLogin2FA {
		statusCode, err := s.deviceSvc.MarkVerifiedDevice(ctx, otp.UID, otp.DeviceUUID)
		if err != nil {
			return statusCode, nil, err
		}
	}

	return status.SUCCESS, &VerifyOTPRes{Status: "verified otp successfully"}, nil
}

func (s *Service) CreateBlock(ctx context.Context, otp *OTP, reason string, duration time.Duration) (status.Code, error) {
	listBlockReq := block.CreateBlockByValueReq{}
	if otp.UID == "" {
		listBlockReq.Items = append(listBlockReq.Items, block.CreateBlockReq{
			Value:    otp.Identifier,
			Type:     enum.BlockTypePhone,
			Duration: duration,
			Reason:   reason,
		})
	}
	listBlockReq.Items = append(listBlockReq.Items, block.CreateBlockReq{
		Value:    otp.DeviceUUID,
		Type:     enum.BlockTypeDevice,
		Duration: duration,
		Reason:   reason,
	})

	statusCode, err := s.blockSvc.CreateMutilpleBlock(ctx, &listBlockReq)
	if err != nil {
		return statusCode, err
	}

	return status.SUCCESS, nil
}

func (s *Service) DeleteOTPByStatus(ctx context.Context, status enum.EOTPStatus) error {
	return s.repo.ForceDeleteOTPByStatus(ctx, status)
}
