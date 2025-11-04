package configs

import (
	"fmt"
	"os"
	"strconv"

	"github.com/QuocAnh189/GoCoreFoundation/internal/db"
	"github.com/joho/godotenv"
)

type Env struct {
	DBEnv      *db.Config
	ServerEnv  *ServerConfig
	HostConfig HostConfig
}

func NewEnv(envpath string) (*Env, error) {
	err := godotenv.Load(envpath)
	if err != nil {
		return nil, fmt.Errorf("failed to load env file: %v", err)
	}

	result := &Env{
		DBEnv: &db.Config{
			Host:     getConfig("DB_HOST"),
			Port:     getConfig("DB_PORT"),
			User:     getConfig("DB_USER"),
			Password: getConfig("DB_PASSWORD"),
			Name:     getConfig("DB_NAME"),
			SSLMode:  getConfig("DB_SSL_MODE"),
		},
		ServerEnv: &ServerConfig{
			Port:        getConfig("SERVER_PORT"),
			LogFilePath: getConfig("LOG_FILE_PATH"),
		},
		HostConfig: HostConfig{
			ServerHost:    getConfig("SERVER_HOST"),
			ServerPort:    getConfig("SERVER_PORT"),
			HttpsCertFile: getConfigOptional("HTTPS_CERT_FILE"),
			HttpsKeyFile:  getConfigOptional("HTTPS_KEY_FILE"),
		},
	}
	return result, nil
}

func getFloatConfigOptional(key string) *float64 {
	if os.Getenv(key) == "" {
		return nil
	}
	val := os.Getenv(key)
	floatVal, _ := strconv.ParseFloat(val, 64)
	return &floatVal
}

func getConfigOptional(key string) *string {
	if os.Getenv(key) == "" {
		return nil
	}
	val := os.Getenv(key)
	return &val
}

func getConfig(key string) string {
	val := getConfigOptional(key)
	if val == nil {
		return ""
	}
	return *val
}

func getBoolConfig(key string) bool {
	val := getConfigOptional(key)
	if val == nil {
		return false
	}
	return *val == "true"
}

func loadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}
