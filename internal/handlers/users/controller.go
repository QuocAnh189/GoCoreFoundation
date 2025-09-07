package users

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/status"
	"github.com/QuocAnh189/GoCoreFoundation/internal/utils/response"
)

type UserController struct {
	service *UserService
}

func NewController(service *UserService) *UserController {
	return &UserController{
		service: service,
	}
}

func (u *UserController) HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	var req ListUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJson(w, nil, ErrInvalidParameter)
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

func (u *UserController) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.WriteJson(w, nil, ErrInvalidUserID)
		return
	}

	user, err := u.service.GetUserByID(r.Context(), userID)
	if err != nil {
		var appErr response.AppError
		switch err {
		case ErrUserNotFound:
			appErr.Message = "User not found"
			appErr.Debug = appErr.Error()
			appErr.Status = status.NOT_FOUND
			response.WriteJson(w, nil, &appErr)
			return
		default:
			response.WriteJson(w, nil, err)
			return
		}
	}

	res := &GetUserResponse{
		User: user,
	}

	response.WriteJson(w, res, nil)
}

func (u *UserController) HandleGetProfile(w http.ResponseWriter, r *http.Request) {
	var userID int64

	user, err := u.service.GetUserByID(r.Context(), userID)
	if err != nil {
		response.WriteJson(w, nil, err)
		return
	}

	response.WriteJson(w, user, nil)
}

func (u *UserController) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJson(w, nil, ErrInvalidParameter)
		return
	}

	user, err := u.service.CreateUser(r.Context(), &req)
	if err != nil {
		response.WriteJson(w, nil, err)
		return
	}
	response.WriteJson(w, user, nil)
}

func (u *UserController) HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
	var req UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJson(w, nil, ErrInvalidParameter)
		return
	}

	userID, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.WriteJson(w, nil, ErrInvalidUserID)
		return
	}
	req.ID = userID

	if err := u.service.UpdateUser(r.Context(), &req); err != nil {
		response.WriteJson(w, nil, err)
		return
	}
	response.WriteJson(w, nil, nil)
}

func (u *UserController) HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.WriteJson(w, nil, ErrInvalidUserID)
		return
	}

	err = u.service.DeleteUser(r.Context(), userID)
	if err != nil {
		response.WriteJson(w, nil, err)
		return
	}

	response.WriteJson(w, "Delete user success", nil)
}
