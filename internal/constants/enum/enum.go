package enum

const (
	RoleAdmin ERole = "admin"
	RoleUser  ERole = "user"
	RoleGuest ERole = "guest"

	StatusActive   EStatus = "ACTIVE"
	StatusInactive EStatus = "INACTIVE"
	StatusBanned   EStatus = "BANNED"

	LoginStatusActive            ELoginStatus = "ACTIVE"
	LoginStatusTwoFactorRequired ELoginStatus = "2FA_REQUIRED"

	BlockTypeDevice EBlockType = "DEVICE"
	BlockTypeEmail  EBlockType = "EMAIL"
	BlockTypePhone  EBlockType = "PHONE"
	BlockTypeIP     EBlockType = "IP"

	OTPPurposeLogin2FA  EOTPPurpose = "LOGIN2FA"
	OTPPurposeSignUp    EOTPPurpose = "SIGNUP"
	OTPPurposeResetPass EOTPPurpose = "RESET_PASSWORD"

	OTPStatusActive   EOTPStatus = "ACTIVE"
	OTPStatusInactive EOTPStatus = "INACTIVE"
	OTPStatusUsed     EOTPStatus = "USED"

	DefaultLang = "en"
)
