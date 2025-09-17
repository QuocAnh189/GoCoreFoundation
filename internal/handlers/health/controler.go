package health

import (
	"net/http"
	"time"

	"github.com/QuocAnh189/GoCoreFoundation/internal/app/resource"
	"github.com/QuocAnh189/GoCoreFoundation/internal/utils/response"
)

type Controller struct {
	appResource *resource.AppResource
	service     *Service
}

func NewController(appResource *resource.AppResource, service *Service) *Controller {
	return &Controller{
		appResource: appResource,
		service:     service,
	}
}

// GET - /health/ping
func (c *Controller) HandlePing(w http.ResponseWriter, r *http.Request) {
	res := PingRes{}
	res.ServerPing = "Go sever live " + time.Now().Format(time.RFC3339)

	err := c.appResource.Db.PingContext(r.Context())
	if err != nil {
		res.DatabasePing = "can not connect database: " + err.Error()
	} else {
		res.DatabasePing = "Database live " + time.Now().Format(time.RFC3339)
	}

	response.WriteJson(w, res, nil)
}
