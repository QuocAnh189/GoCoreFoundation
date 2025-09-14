package lingos

import (
	"net/http"

	"github.com/QuocAnh189/GoCoreFoundation/internal/utils/bind"
	"github.com/QuocAnh189/GoCoreFoundation/internal/utils/response"
)

type Controller struct {
	service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{
		service: service,
	}
}

// POST - /lingos/create
func (c *Controller) HandleCreateLingo(w http.ResponseWriter, r *http.Request) {
	// return c.service.CreateLingo(ctx, l)
	response.WriteJson(w, nil, nil)
}

// POST - /lingos/list
func (c *Controller) HandleGetListLingos(w http.ResponseWriter, r *http.Request) {
	// return c.service.CreateLingo(ctx, l)
	response.WriteJson(w, nil, nil)
}

// Get - /lingos?lang=en&key=welcome
func (c *Controller) HandleGetLingo(w http.ResponseWriter, r *http.Request) {
	var req GetLingoRequest

	if err := bind.ParseQuery(r, &req); err != nil {
		response.WriteJson(w, nil, response.ErrInvalidParams())
		return
	}

	lingo, err := c.service.GetLingo(r.Context(), Lang(req.Lang), req.Key)
	if err != nil {
		response.WriteJson(w, nil, err)
		return
	}

	res := &GetLingoResponse{
		Lingo: lingo,
	}

	response.WriteJson(w, res, nil)
}

// POST - /lingos/update
func (c *Controller) HandleUpdateLingo(w http.ResponseWriter, r *http.Request) {
	// return c.service.UpdateLingo(ctx, l)
	response.WriteJson(w, nil, nil)
}

// POST - /lingos/delete
func (c *Controller) HandleDeleteLingo(w http.ResponseWriter, r *http.Request) {
	// return c.service.DeleteLingo(ctx, id)
	response.WriteJson(w, nil, nil)
}
