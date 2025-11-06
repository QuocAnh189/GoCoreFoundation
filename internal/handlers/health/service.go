package health

import (
	"time"

	"github.com/QuocAnh189/GoCoreFoundation/internal/services/sms"
)

type Service struct {
	smsSvc *sms.Service
}

func NewService(smsSvc *sms.Service) *Service {
	return &Service{
		smsSvc: smsSvc,
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
	err := s.smsSvc.SendSmsToPhone("+84905636640", "Ping test message from service ID")
	if err != nil {
		// res.SmsMessage = "can not send sms: " + err.Error()
		return "can not send sms: " + err.Error()
	}

	return "SMS sent successfully at " + time.Now().Format(time.RFC3339)
}
