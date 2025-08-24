package users

import (
	"net/http"

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

func (u *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	response.WriteJson(w, []byte("Get user"), nil)
}

func (u *UserController) GetProfile(w http.ResponseWriter, r *http.Request) {
	response.WriteJson(w, []byte("Get profile"), nil)
}

func (u *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	response.WriteJson(w, []byte("Get users"), nil)
}

func (u *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	response.WriteJson(w, []byte("Create user"), nil)
}

func (u *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	response.WriteJson(w, []byte("Update user"), nil)
}

func (u *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	response.WriteJson(w, []byte("Delete user"), nil)
}
