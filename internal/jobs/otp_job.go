package jobs

import (
	"context"

	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/enum"
	"github.com/QuocAnh189/GoCoreFoundation/internal/handlers/otp"
)

type DeleteOTPJob struct {
	name   string
	id     int
	otpSvc *otp.Service
}

func NewOTPJob(otpSvc *otp.Service) *DeleteOTPJob {
	return &DeleteOTPJob{
		name:   "otp-job",
		id:     3,
		otpSvc: otpSvc,
	}
}

func (j *DeleteOTPJob) Name() string {
	return j.name
}

func (j *DeleteOTPJob) TickInterval() int {
	return 1
}

func (j *DeleteOTPJob) Run(ctx context.Context) error {
	return j.otpSvc.DeleteOTPByStatus(ctx, enum.OTPStatusInactive)
}
