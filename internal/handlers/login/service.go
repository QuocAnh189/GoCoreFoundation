package login

import (
	"context"
	"errors"

	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/status"
	"github.com/QuocAnh189/GoCoreFoundation/internal/sessions"
	appctx "github.com/QuocAnh189/GoCoreFoundation/internal/utils/context"
)

type Service struct {
	Repository IRepository
}

func NewService(repo IRepository) *Service {
	return &Service{
		Repository: repo,
	}
}

func (s *Service) Login(ctx context.Context, sess *sessions.AppSession, req *LoginReq) (status.Code, string, error) {
	logger := appctx.GetLogger(ctx)

	// Get auth token
	authTokenRaw, ok := sess.Get("token")
	if !ok {
		return status.INTERNAL, "", errors.New("session token not found")
	}
	authToken := authTokenRaw.(string)

	logger.Info("Login successful, updating session data...")
	sess.Init(sessions.InitData{
		Source:    "login",
		IsSecure:  true,
		UID:       123,
		Email:     "test@gmail.com",
		LoginName: req.LoginName,
	})

	return status.SUCCESS, authToken, nil
}
