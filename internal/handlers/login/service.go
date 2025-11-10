package login

import (
	"context"
	"log"

	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/status"
	"github.com/QuocAnh189/GoCoreFoundation/internal/handlers/device"
	"github.com/QuocAnh189/GoCoreFoundation/internal/handlers/users"
	"github.com/QuocAnh189/GoCoreFoundation/internal/sessions"
	appctx "github.com/QuocAnh189/GoCoreFoundation/internal/utils/context"
)

type Service struct {
	repo       IRepository
	userRepo   users.IRepository
	deviceRepo device.IRepository
}

func NewService(repo IRepository, userRepo users.IRepository, deviceRepo device.IRepository) *Service {
	return &Service{
		repo:       repo,
		userRepo:   userRepo,
		deviceRepo: deviceRepo,
	}
}

func (s *Service) Login(ctx context.Context, sess *sessions.AppSession, req *LoginReq) (status.Code, *LoginRes, error) {
	logger := appctx.GetLogger(ctx)

	var (
		isSecure bool
		needs2FA bool
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

	// check password
	// originalPassword := "test"
	// newHash, err := utils.DefaultHasher.Hash(originalPassword)
	// if err != nil {
	// 	logger.Error("Failed to generate hash: %v", err)
	// 	return status.LOGIN_WRONG_CREDENTIALS, nil, ErrInvalidCredentials
	// }
	// err = utils.DefaultHasher.Compare(originalPassword, newHash)
	// if err != nil {
	// 	logger.Info("Comparison with stored hash failed: %v", err)
	// }

	println("raw", req.RawPassword)
	println("hash", user.Password)
	// err = utils.DefaultHasher.Compare(req.RawPassword, user.Password)
	// if err != nil {
	// 	logger.Info("Error comparing password: %v", err)
	// 	return status.LOGIN_WRONG_CREDENTIALS, nil, ErrInvalidCredentials
	// }

	// // Get auth token
	// authTokenRaw, ok := sess.Get("token")
	// if !ok {
	// 	return status.INTERNAL, "", errors.New("session token not found")
	// }
	// authToken := authTokenRaw.(string)

	// logger.Info("Login successful, updating session data...")
	// sess.Init(sessions.InitData{
	// 	Source:    "login",
	// 	IsSecure:  true,
	// 	UID:       123,
	// 	Email:     "test@gmail.com",
	// 	LoginName: req.LoginName,
	// })

	device, err := s.deviceRepo.GetDeviceByUIDAnDeviceUUID(ctx, user.ID, req.DeviceUUID)
	if err != nil {
		log.Fatalf("Failed to get device by UUID: %v", err)
		return status.INTERNAL, nil, err
	}
	if device != nil && device.IsVerified {
		isSecure = true
		needs2FA = false
	} else {
		isSecure = false
		needs2FA = true
	}

	res := &LoginRes{
		IsSecure: isSecure,
		Needs2FA: needs2FA,
		User:     user,
	}

	return status.SUCCESS, res, nil
}
