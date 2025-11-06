package resource

import (
	"fmt"
	"net/http"

	"github.com/QuocAnh189/GoCoreFoundation/internal/configs"
	"github.com/QuocAnh189/GoCoreFoundation/internal/db"
	"github.com/QuocAnh189/GoCoreFoundation/internal/sessions"
	"github.com/QuocAnh189/GoCoreFoundation/root"
)

type AppResource struct {
	Env            *configs.Env
	HostConfig     root.HostConfig
	Db             db.IDatabase
	SessionManager *sessions.SessionManager
}

func (a *AppResource) GetRequestSession(r *http.Request) (*sessions.AppSession, error) {
	sess := sessions.GetRequestSession(r)
	if sess == nil {
		return nil, fmt.Errorf("session not found")
	}

	return sess, nil
}

func (a *AppResource) GetRequestUID(r *http.Request) (int64, error) {
	sess, err := a.GetRequestSession(r)
	if err != nil {
		return 0, err
	}

	id, ok := sess.UID()
	if !ok {
		return 0, fmt.Errorf("uid missing from session (did you forget to send the Authorization header?)")
	}

	return id, nil
}
