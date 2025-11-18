package session_provider

import (
	"errors"

	"github.com/QuocAnh189/GoCoreFoundation/root/session"
)

var (
	ErrMalformedJwt = errors.New("invalid or malformed JWT")
)

type SessionFactory func() session.SessionStorer

// SessionResult contains the session and metadata about the session retrieval
type SessionResult struct {
	Session        session.SessionStorer
	DidAutoRefresh bool
	AuthToken      string
}
