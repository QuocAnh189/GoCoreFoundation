package misc

import (
	"net/http"

	"github.com/QuocAnh189/GoCoreFoundation/internal/app/resource"
	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/status"
	"github.com/QuocAnh189/GoCoreFoundation/internal/sessions"
	"github.com/QuocAnh189/GoCoreFoundation/internal/utils/response"
)

type Controller struct {
	appResources *resource.AppResource
}

func NewController(res *resource.AppResource) *Controller {
	return &Controller{
		appResources: res,
	}
}

func (c *Controller) HandleSessionDump(w http.ResponseWriter, r *http.Request) {
	dumpedSession := sessions.Dump(c.appResources.SessionManager)
	response.WriteJson(w, r.Context(), dumpedSession, nil, status.SUCCESS)
}
