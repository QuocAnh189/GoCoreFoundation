package health

type PingRes struct {
	ServerPing   string `json:"server_ping"`
	DatabasePing string `json:"database_ping"`
}
