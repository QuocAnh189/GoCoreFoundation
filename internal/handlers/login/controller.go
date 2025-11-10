package login

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/QuocAnh189/GoCoreFoundation/internal/app/resource"
	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/status"
	"github.com/QuocAnh189/GoCoreFoundation/internal/utils/response"
)

type Controller struct {
	appResources *resource.AppResource
	service      *Service
}

func NewController(
	appResources *resource.AppResource,
	service *Service,
) *Controller {
	return &Controller{
		appResources: appResources,
		service:      service,
	}
}

// POST - /login
func (c *Controller) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var req LoginReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.WriteJson(w, r.Context(), nil, fmt.Errorf("invalid parameters"), status.BAD_REQUEST)
		return
	}
	defer r.Body.Close()
	req.DeviceUUID = r.Header.Get("Device-UUID")
	req.DeviceName = r.Header.Get("Device-Name")

	ctx := r.Context()

	sess, err := c.appResources.GetRequestSession(r)
	if err != nil {
		response.WriteJson(w, ctx, nil, fmt.Errorf("failed to get session %w", err), status.INTERNAL)
		return
	}

	// Perform login
	statusCode, res, err := c.service.Login(ctx, sess, &req)
	if err != nil {
		response.WriteJson(w, ctx, nil, err, statusCode)
		return
	}

	response.WriteJson(w, ctx, res, nil, statusCode)
}
