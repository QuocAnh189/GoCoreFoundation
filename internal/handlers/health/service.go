package health

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

// Check Stripe

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
func (s *Service) CheckTwilio() bool {
	return true
}
