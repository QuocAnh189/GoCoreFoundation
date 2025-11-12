package otp

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/QuocAnh189/GoCoreFoundation/internal/app/resource"
	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/status"
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

// POST - /otp/send
func (c *Controller) HandleSendOTP(w http.ResponseWriter, r *http.Request) {
	var req SendOTPReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.WriteJson(w, r.Context(), nil, fmt.Errorf("invalid parameters"), status.BAD_REQUEST)
		return
	}
	defer r.Body.Close()
	req.DeviceUUID = r.Header.Get("Device-UUID")
	req.DeviceName = r.Header.Get("Device-Name")

	ctx := r.Context()
	statusCode, res, err := c.service.SendOTP(ctx, &req)
	if err != nil {
		response.WriteJson(w, ctx, nil, err, statusCode, GetArgsByStatatus(statusCode, res, nil)...)
		return
	}

	response.WriteJson(w, ctx, res, nil, 200)
}

// POST - /otp/verify
func (c *Controller) HandleVerifyOTP(w http.ResponseWriter, r *http.Request) {
	var req VerifyOTPReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.WriteJson(w, r.Context(), nil, fmt.Errorf("invalid parameters"), status.BAD_REQUEST)
		return
	}
	defer r.Body.Close()
	req.DeviceUUID = r.Header.Get("Device-UUID")
	req.DeviceName = r.Header.Get("Device-Name")

	ctx := r.Context()
	statusCode, res, err := c.service.VerifyOTP(ctx, &req)
	if err != nil {
		response.WriteJson(w, ctx, nil, err, statusCode, GetArgsByStatatus(statusCode, nil, res)...)
		return
	}

	response.WriteJson(w, ctx, res, nil, 200)
}
