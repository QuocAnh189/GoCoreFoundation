package configs

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

type ServerConfig struct {
	Port        string
	LogFilePath string
}

type HostConfig struct {
	ServerHost    string
	ServerPort    string
	HttpsCertFile *string
	HttpsKeyFile  *string
}

type TwilioConfig struct {
	AccountSID          *string
	AuthToken           *string
	FromPhoneNumber     *string
	MessagingServiceSID *string
}
