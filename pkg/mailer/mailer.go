package mailer

import (
	"gopkg.in/gomail.v2"
)

type Client struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

func NewClient(host string, port int, username, password, from string) *Client {
	return &Client{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		From:     from,
	}
}

func (m *Client) SendEmail(to string, subject string, body string, isHTML bool) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", m.From)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)

	if isHTML {
		msg.SetBody("text/html", body)
	} else {
		msg.SetBody("text/plain", body)
	}

	dialer := gomail.NewDialer(m.Host, m.Port, m.Username, m.Password)

	if err := dialer.DialAndSend(msg); err != nil {
		return err
	}

	return nil
}
