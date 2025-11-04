package configs

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
