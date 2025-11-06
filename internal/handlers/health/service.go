package health

import (
	"time"

	"github.com/QuocAnh189/GoCoreFoundation/internal/services/mail"
	"github.com/QuocAnh189/GoCoreFoundation/internal/services/sms"
)

type Service struct {
	smsSvc  *sms.Service
	mailSvc *mail.Service
}

func NewService(smsSvc *sms.Service, mailSvc *mail.Service) *Service {
	return &Service{
		smsSvc:  smsSvc,
		mailSvc: mailSvc,
	}
}

// Check Stripe
func (s *Service) CheckStripe() bool {
	return true
}

// Check Goat
func (s *Service) CheckGoat() bool {
	return true
}

// Check Cybersource
func (s *Service) CheckCybersource() bool {
	return true
}

// Check Plaid
func (s *Service) CheckPlaid() bool {
	return true
}

// Check Twilio
func (s *Service) CheckTwilio() string {
	if s.smsSvc == nil {
		return "SMS service is not configured"
	}

	err := s.smsSvc.SendSmsToPhone("+84905636640", "Ping test message from service ID")
	if err != nil {
		// res.SmsMessage = "can not send sms: " + err.Error()
		return "can not send sms: " + err.Error()
	}

	return "SMS sent successfully at " + time.Now().Format(time.RFC3339)
}

// Check Mailer
func (s *Service) CheckMailer() string {
	if s.mailSvc == nil {
		return "Mailer service is not configured"
	}

	if err := s.mailSvc.Send("binbin18092003@gmail.com", "Hello!", "<h1>Congratulations</h1><p>Your account has been successfully created</p>", true); err != nil {
		return "can not send email: " + err.Error()
	}

	return "Mail sent successfully at " + time.Now().Format(time.RFC3339)
}
