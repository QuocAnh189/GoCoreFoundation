package sms

import "errors"

type Service struct {
	provider Provider
}

func NewService(provider Provider) *Service {
	return &Service{provider: provider}
}

func (s *Service) SendSmsToPhone(to, body string) error {
	if s.provider == nil {
		return errors.New("twilio client is not configured")
	}

	return s.provider.SendSMS(to, body)
}
