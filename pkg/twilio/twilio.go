package twilio

import (
	"log"

	"github.com/QuocAnh189/GoCoreFoundation/internal/utils/phone"
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

type credentials struct {
	accountSID string
	authToken  string
}

type Client struct {
	TwilioNumber    string
	TwilioServiceID string

	credentials credentials
	client      *twilio.RestClient
}

func NewClient(accountSid, authToken, twilioNumber, twilioServiceId string) *Client {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	return &Client{
		TwilioNumber:    twilioNumber,
		TwilioServiceID: twilioServiceId,
		credentials: credentials{
			accountSID: accountSid,
			authToken:  authToken,
		},
		client: client,
	}
}

func (c *Client) SendSMSByPhoneNumber(to, body string) error {
	to, err := phone.NormalizePhone(to)
	if err != nil {
		return err
	}

	log.Printf("SendSMSByPhoneNumber: Sending SMS to %s with from=%s\n", to, c.TwilioNumber)

	params := buildParams(to, body)
	params.SetFrom(c.TwilioNumber)
	if _, err := c.client.Api.CreateMessage(params); err != nil {
		return err
	}

	return nil
}

func (c *Client) SendSMSByMessagingService(to, body string) error {
	to, err := phone.NormalizePhone(to)
	if err != nil {
		return err
	}

	log.Printf("SendSMSByMessagingService: Sending SMS to %s with serviceId=%s\n", to, c.TwilioServiceID)

	params := buildParams(to, body)
	params.SetMessagingServiceSid(c.TwilioServiceID)
	if _, err := c.client.Api.CreateMessage(params); err != nil {
		return err
	}

	return nil
}

func buildParams(to, body string) *api.CreateMessageParams {
	params := &api.CreateMessageParams{}
	params.SetTo(to)
	params.SetBody(body)
	return params
}
