package sms

type Service struct {
	provider Provider
}

func NewService(provider Provider) *Service {
	return &Service{provider: provider}
}

func (s *Service) SendSmsToPhone(to, body string) error {
	return s.provider.SendSMS(to, body)
}
