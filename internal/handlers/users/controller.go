package users

import (
	"encoding/json"
	"net/http"

	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/status"
	"github.com/QuocAnh189/GoCoreFoundation/internal/utils/bind"
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

// Get - /users/list
func (u *UserController) HandleGetUsers(w http.ResponseWriter, r *http.Request) {
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
func (u *UserController) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("id")

	user, err := u.service.GetUserByID(r.Context(), userID)
	if err != nil {
		var appErr response.AppError
		appErr.BaseError = err
		switch err {
		case ErrUserNotFound:
			appErr.Message = "User not found"
			appErr.Status = status.NOT_FOUND
		default:
			appErr.Message = "Something went wrong"
		}
		response.WriteJson(w, nil, &appErr)
		return
	}

	res := &GetUserResponse{
		User: user,
	}

	response.WriteJson(w, res, nil)
}

// Get - /users/profile
func (u *UserController) HandleGetProfile(w http.ResponseWriter, r *http.Request) {
	var userID string

	user, err := u.service.GetUserByID(r.Context(), userID)
	if err != nil {
		response.WriteJson(w, nil, err)
		return
	}

	response.WriteJson(w, user, nil)
}

// POST - /users/create
func (u *UserController) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJson(w, nil, response.ErrInvalidParams())
		return
	}

	user, err := u.service.CreateUser(r.Context(), &req)
	if err != nil {
		var appErr response.AppError
		appErr.BaseError = err
		switch err {
		case ErrMissingFirstName, ErrMissingLastName, ErrMissingPhone, ErrMissingEmail, ErrInvalidEmail, ErrInvalidRole:
			appErr.Message = "Invalid params"
			appErr.Status = status.BAD_REQUEST
		default:
			appErr.Message = "Something went wrong"
		}
		response.WriteJson(w, nil, &appErr)
		return
	}

	res := &CreateUserResponse{
		User: user,
	}

	response.WriteJson(w, res, nil)
}

// POST - /users/update
func (u *UserController) HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
	var req UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJson(w, nil, response.ErrInvalidParams())
		return
	}

	user, err := u.service.UpdateUser(r.Context(), &req)
	if err != nil {
		var appErr response.AppError
		appErr.BaseError = err
		switch err {
		case ErrInvalidEmail, ErrInvalidRole:
			appErr.Message = "Invalid params"
			appErr.Status = status.BAD_REQUEST
		case ErrInvalidUserID:
			appErr.Message = "Invalid user ID"
			appErr.Status = status.BAD_REQUEST
		default:
			appErr.Message = "Something went wrong"
		}
		response.WriteJson(w, nil, &appErr)
		return
	}

	res := &UpdateUserResponse{
		User: user,
	}

	response.WriteJson(w, res, nil)
}

// POST - /users/delete
func (u *UserController) HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	var req DeleteUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJson(w, nil, response.ErrInvalidParams())
		return
	}

	err := u.service.DeleteUser(r.Context(), req.UserID)
	if err != nil {
		var appErr response.AppError
		appErr.BaseError = err
		switch err {
		case ErrInvalidUserID:
			appErr.Message = "Invalid user ID"
			appErr.Status = status.BAD_REQUEST
		default:
			appErr.Message = "Something went wrong"
		}
		response.WriteJson(w, nil, &appErr)
		return
	}

	response.WriteJson(w, "Delete user success", nil)
}
