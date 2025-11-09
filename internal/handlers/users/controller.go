package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/QuocAnh189/GoCoreFoundation/internal/app/resource"
	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/status"
	"github.com/QuocAnh189/GoCoreFoundation/internal/utils/bind"
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
		response.WriteJson(w, r.Context(), nil, fmt.Errorf("invalid parameters"), status.BAD_REQUEST)
		return
	}

	statusCode, users, pagination, err := u.service.ListUsers(r.Context(), &req)
	if err != nil {
		response.WriteJson(w, r.Context(), nil, err, statusCode)
		return
	}

	res := &ListUserResponse{
		Users:      users,
		Pagination: pagination,
	}

	response.WriteJson(w, r.Context(), res, nil, statusCode)
}

// Get - /users/{id}
func (u *Controller) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("id")

	statusCode, user, err := u.service.GetUserByID(r.Context(), userID)
	if err != nil {
		response.WriteJson(w, r.Context(), nil, err, statusCode)
		return
	}

	res := &GetUserResponse{
		User: user,
	}

	response.WriteJson(w, r.Context(), res, nil, statusCode)
}

// Get - /users/profile
func (u *Controller) HandleGetProfile(w http.ResponseWriter, r *http.Request) {
	var userID string

	statusCode, res, err := u.service.GetUserByID(r.Context(), userID)
	if err != nil {
		response.WriteJson(w, r.Context(), nil, err, statusCode)
		return
	}

	response.WriteJson(w, r.Context(), res, nil, statusCode)
}

// POST - /users/create
func (u *Controller) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJson(w, r.Context(), nil, fmt.Errorf("invalid parameters"), status.BAD_REQUEST)
		return
	}

	statusCode, user, err := u.service.CreateUser(r.Context(), &req)
	if err != nil {
		response.WriteJson(w, r.Context(), nil, err, statusCode, GetArgsByStatatus(statusCode)...)
		return
	}

	res := &CreateUserResponse{
		User: user,
	}

	response.WriteJson(w, r.Context(), res, nil, statusCode)
}

// POST - /users/update
func (u *Controller) HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
	var req UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJson(w, r.Context(), nil, fmt.Errorf("invalid parameters"), status.BAD_REQUEST)
		return
	}

	statusCode, user, err := u.service.UpdateUser(r.Context(), &req)
	if err != nil {
		response.WriteJson(w, r.Context(), nil, err, statusCode, GetArgsByStatatus(statusCode)...)
		return
	}

	res := &UpdateUserResponse{
		User: user,
	}

	response.WriteJson(w, r.Context(), res, nil, statusCode)
}

// POST - /users/delete
func (u *Controller) HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	var req DeleteUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJson(w, r.Context(), nil, fmt.Errorf("invalid parameters"), status.BAD_REQUEST)
		return
	}

	statusCode, err := u.service.DeleteUser(r.Context(), req.UserID)
	if err != nil {
		response.WriteJson(w, r.Context(), nil, err, statusCode)
		return
	}

	response.WriteJson(w, r.Context(), "Delete successfully", nil, statusCode)
}

// POST - /users/force-delete
func (u *Controller) HandleForceDeleteUser(w http.ResponseWriter, r *http.Request) {
	var req DeleteUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJson(w, r.Context(), nil, fmt.Errorf("invalid parameters"), status.BAD_REQUEST)
		return
	}

	statusCode, err := u.service.ForceDeleteUser(r.Context(), req.UserID)
	if err != nil {
		response.WriteJson(w, r.Context(), nil, err, statusCode)
		return
	}

	response.WriteJson(w, r.Context(), "Force delete successfully", nil, statusCode)
}
