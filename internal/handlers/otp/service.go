package otp

import (
	"context"

	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/status"
)

type Service struct {
	repo IRepository
}

func NewService(repo IRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) SendOTP(ctx context.Context, req *SendOTPReq) (status.Code, *SendOtpRes, error) {
	if statusCode, err := ValidateSendOTPReq(req); err != nil {
		return statusCode, nil, err
	}
	return status.SUCCESS, nil, nil
}

func (s *Service) VerifyOTP(ctx context.Context, req *VerifyOTPReq) (status.Code, *VerifyOTPRes, error) {
	return status.SUCCESS, nil, nil
}
