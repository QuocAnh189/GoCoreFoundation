package mail

import (
	"errors"

	"github.com/QuocAnh189/GoCoreFoundation/pkg/mailer"
)

type Service struct {
	client *mailer.Client
}

func NewService(client *mailer.Client) *Service {
	return &Service{
		client: client,
	}
}

func (s *Service) Send(to string, subject string, body string, isHTML bool) error {
	if s.client == nil {
		return errors.New("mailer client is not configured")
	}
	return s.client.SendEmail(to, subject, body, isHTML)
}
