package users

import (
	"encoding/json"
	"net/http"

	"github.com/QuocAnh189/GoCoreFoundation/internal/app/resource"
	"github.com/QuocAnh189/GoCoreFoundation/internal/utils/bind"
	ctx "github.com/QuocAnh189/GoCoreFoundation/internal/utils/context"
	"github.com/QuocAnh189/GoCoreFoundation/internal/utils/response"
)

type Controller struct {
	appResources *resource.AppResource
	service      *Service
}

func NewController(appResources *resource.AppResource, service *Service) *Controller {
	return &Controller{
		appResources: appResources,
		service:      service,
	}
}

// Get - /users/list
func (u *Controller) HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	var req ListUserRequest

	if err := bind.ParseQuery(r, &req); err != nil {
		response.WriteJson(w, nil, response.ErrInvalidParams())
		return
	}

	users, pagination, err := u.service.ListUsers(r.Context(), &req)
	if err != nil {
		response.WriteJson(w, nil, err)
		return
	}

	res := &ListUserResponse{
		Users:      users,
		Pagination: pagination,
	}

	response.WriteJson(w, res, nil)
}

// Get - /users/{id}
func (u *Controller) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("id")

	language := ctx.GetLocale(r.Context())

	println("language", language)

	user, err := u.service.GetUserByID(r.Context(), userID)
	if err != nil {
		var appErr response.AppError
		appErr.BaseError = err
		appErr.Status = DetermineErrStatus(err)
		appErr.Message = GetMessageFromKey(language, DetermineErrKey(err))

		response.WriteJson(w, nil, &appErr)
		return
	}

	res := &GetUserResponse{
		User: user,
	}

	response.WriteJson(w, res, nil)
}

// Get - /users/profile
func (u *Controller) HandleGetProfile(w http.ResponseWriter, r *http.Request) {
	var userID string

	language := ctx.GetLocale(r.Context())

	user, err := u.service.GetUserByID(r.Context(), userID)
	if err != nil {
		var appErr response.AppError
		appErr.BaseError = err
		appErr.Status = DetermineErrStatus(err)
		appErr.Message = GetMessageFromKey(language, DetermineErrKey(err))

		response.WriteJson(w, nil, &appErr)
		return
	}

	response.WriteJson(w, user, nil)
}

// POST - /users/create
func (u *Controller) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJson(w, nil, response.ErrInvalidParams())
		return
	}

	language := ctx.GetLocale(r.Context())

	user, err := u.service.CreateUser(r.Context(), &req)
	if err != nil {
		var appErr response.AppError
		appErr.BaseError = err
		appErr.Status = DetermineErrStatus(err)
		appErr.Message = GetMessageFromKey(language, DetermineErrKey(err))

		response.WriteJson(w, nil, &appErr)
		return
	}

	res := &CreateUserResponse{
		User: user,
	}

	response.WriteJson(w, res, nil)
}

// POST - /users/update
func (u *Controller) HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
	var req UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		println("UpdateUser request:", req.UID)
		response.WriteJson(w, nil, response.ErrInvalidParams())
		return
	}

	language := ctx.GetLocale(r.Context())

	user, err := u.service.UpdateUser(r.Context(), &req)
	if err != nil {
		var appErr response.AppError
		appErr.BaseError = err
		appErr.Status = DetermineErrStatus(err)
		appErr.Message = GetMessageFromKey(language, DetermineErrKey(err))

		response.WriteJson(w, nil, &appErr)
		return
	}

	res := &UpdateUserResponse{
		User: user,
	}

	response.WriteJson(w, res, nil)
}

// POST - /users/delete
func (u *Controller) HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	var req DeleteUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJson(w, nil, response.ErrInvalidParams())
		return
	}

	language := ctx.GetLocale(r.Context())

	err := u.service.DeleteUser(r.Context(), req.UserID)
	if err != nil {
		var appErr response.AppError
		appErr.BaseError = err
		appErr.Status = DetermineErrStatus(err)
		appErr.Message = GetMessageFromKey(language, DetermineErrKey(err))

		response.WriteJson(w, nil, &appErr)
		return
	}

	response.WriteJson(w, "Delete user success", nil)
}
