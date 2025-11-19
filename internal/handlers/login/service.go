package login

import (
	"context"
	"errors"

	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/enum"
	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/status"
	"github.com/QuocAnh189/GoCoreFoundation/internal/handlers/device"
	"github.com/QuocAnh189/GoCoreFoundation/internal/handlers/user"
	"github.com/QuocAnh189/GoCoreFoundation/internal/sessions"
	"github.com/QuocAnh189/GoCoreFoundation/internal/utils"
	appctx "github.com/QuocAnh189/GoCoreFoundation/internal/utils/context"
)

type Service struct {
	repo       IRepository
	userRepo   user.IRepository
	deviceRepo device.IRepository
}

func NewService(repo IRepository, userRepo user.IRepository, deviceRepo device.IRepository) *Service {
	return &Service{
		repo:       repo,
		userRepo:   userRepo,
		deviceRepo: deviceRepo,
	}
}

func (s *Service) Login(ctx context.Context, sess *sessions.AppSession, req *LoginReq) (status.Code, *LoginRes, error) {
	logger := appctx.GetLogger(ctx)

	var (
		isSecure    bool              = true
		needs2FA    bool              = false
		loginStatus enum.ELoginStatus = enum.LoginStatusActive
	)

	statusCode, err := ValidateLoginReq(req)
	if err != nil {
		return statusCode, nil, err
	}

	// get user by login name
	user, err := s.userRepo.GetUserByLoginName(ctx, req.LoginName)
	if err != nil {
		return status.INTERNAL, nil, err
	}
	if user == nil {
		logger.Info("User not found with login name: %s", req.LoginName)
		return status.LOGIN_WRONG_CREDENTIALS, nil, ErrInvalidCredentials
	}

	// compare password
	err = utils.DefaultHasher.Compare(req.RawPassword, user.Password)
	if err != nil {
		logger.Info("Error comparing password: %v", err)
		return status.LOGIN_WRONG_CREDENTIALS, nil, ErrInvalidCredentials
	}

	isTrustedDevice, err := s.deviceRepo.CheckTrustedDeviceByUID(ctx, user.ID, req.DeviceUUID)
	if err != nil {
		return status.INTERNAL, nil, err
	}

	if !isTrustedDevice {
		isSecure = false
		needs2FA = true
		loginStatus = enum.LoginStatusTwoFactorRequired
	}

	// Get auth token
	authTokenRaw, ok := sess.Get("token")
	if !ok {
		return status.FAIL, nil, errors.New("session token not found")
	}
	authToken := authTokenRaw.(string)

	// Update session
	logger.Info("Login successful, updating session data...")
	sess.Init(sessions.InitData{
		Source:    "login",
		IsSecure:  isTrustedDevice,
		UID:       user.ID,
		Email:     user.Email,
		LoginName: req.LoginName,
	})

	loginLog, err := s.repo.GetLoginLogByUIDAndDeviceUUID(ctx, user.ID, req.DeviceUUID)
	if err != nil {
		logger.Info("Error getting login log: %v", err)
		return status.INTERNAL, nil, err
	}

	if loginLog != nil {
		updateLoginLogDTO := BuildUpdateLoginLogDTO(loginLog.ID, user.ID, req.IpAddress, req.DeviceUUID, authToken, loginStatus)
		err = s.repo.UpdateLoginLog(ctx, updateLoginLogDTO)
		if err != nil {
			logger.Info("Error updating login log: %v", err)
			return status.INTERNAL, nil, err
		}
	} else {
		createLoginLogDTO := BuildCreateLoginLogDTO(user.ID, req.IpAddress, req.DeviceUUID, authToken, loginStatus)
		err = s.repo.StoreLoginLog(ctx, createLoginLogDTO)
		if err != nil {
			logger.Info("Error storing login log: %v", err)
			return status.INTERNAL, nil, err
		}
	}

	res := &LoginRes{
		IsSecure:    isSecure,
		Needs2FA:    needs2FA,
		User:        user,
		LoginStatus: loginStatus,
	}

	return status.SUCCESS, res, nil
}
