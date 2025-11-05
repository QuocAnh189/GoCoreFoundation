package sessionprovider

import (
	"net/http"
)

type SessionProvider interface {
	GetSessionFromRequest(r *http.Request) (*SessionResult, error)
}
