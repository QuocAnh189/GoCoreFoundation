package sms

// Provider represents a specific SMS sending method (Twilio, VN sms provider, etc.)
// This is used internally by the SMS service and allows the
// dependency injection to specify how SMS are sent.
//
// NOTEDo not directly use this interface in other services! Use the SMS service instead.

type Provider interface {
	SendSMS(to, body string) error
}
