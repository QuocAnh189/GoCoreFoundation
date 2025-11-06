package sms

import (
	"errors"

	"github.com/QuocAnh189/GoCoreFoundation/pkg/twilio"
)

type TwilioProvider struct {
	client *twilio.Client
}

func NewTwilioSmsProvider(accountSid, authToken, twilioNumber, twilioServiceId string) *TwilioProvider {
	client := twilio.NewClient(
		accountSid,
		authToken,
		twilioNumber,
		twilioServiceId,
	)
	return &TwilioProvider{client: client}
}

func (c *TwilioProvider) SendSMS(to, body string) error {
	if c.client == nil {
		return errors.New("twilio client is not configured")
	}

	return c.client.SendSMSByMessagingService(to, body)
}
