package users

import "net/http"

type UserController struct {
	service *UserService
}

func NewController(service *UserService) *UserController {
	return &UserController{
		service: service,
	}
}

func (u *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get users"))
}
