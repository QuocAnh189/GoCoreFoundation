package users

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/QuocAnh189/GoCoreFoundation/internal/constants"
	"github.com/QuocAnh189/GoCoreFoundation/internal/db"
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
	ctx := r.Context()
	role := r.URL.Query().Get("role")
	var opts []db.FindOption
	if role != "" {
		opts = append(opts, db.WithQuery(db.NewQuery("role = ?", constants.Role(role))))
	}

	users, err := u.service.GetListUser(ctx, opts...)
	if err != nil {
		response.WriteJson(w, nil, err)
		return
	}
	response.WriteJson(w, users, nil)
}

func (u *UserController) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.WriteJson(w, nil, errors.New("invalid user ID"))
		return
	}

	user, err := u.service.GetUserByID(ctx, id)
	if err != nil {
		response.WriteJson(w, nil, err)
		return
	}
	response.WriteJson(w, user, nil)
}

func (u *UserController) HandleGetProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value("user_id").(int64)
	if !ok || userID == 0 {
		response.WriteJson(w, nil, errors.New("unauthorized: user ID not found in context"))
		return
	}

	user, err := u.service.GetUserByID(ctx, userID)
	if err != nil {
		response.WriteJson(w, nil, err)
		return
	}
	response.WriteJson(w, user, nil)
}

func (u *UserController) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response.WriteJson(w, nil, errors.New("invalid JSON body"))
		return
	}
	user.ID = 0 // Ensure ID is not set (auto-increment)

	userID, ok := ctx.Value("user_id").(int)
	if ok {
		user.CreateID = &userID
		user.ModifyID = &userID
	}

	if err := u.service.CreateUser(ctx, &user); err != nil {
		response.WriteJson(w, nil, err)
		return
	}
	response.WriteJson(w, user, nil)
}

func (u *UserController) HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.WriteJson(w, nil, errors.New("invalid user ID"))
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response.WriteJson(w, nil, errors.New("invalid JSON body"))
		return
	}
	user.ID = id

	userID, ok := ctx.Value("user_id").(int)
	if ok {
		user.ModifyID = &userID
	}

	if err := u.service.UpdateUser(ctx, &user); err != nil {
		response.WriteJson(w, nil, err)
		return
	}
	response.WriteJson(w, user, nil)
}

func (u *UserController) HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.WriteJson(w, nil, errors.New("invalid user ID"))
		return
	}

	if err := u.service.DeleteUser(ctx, id); err != nil {
		response.WriteJson(w, nil, err)
		return
	}
	response.WriteJson(w, map[string]string{"message": "user deleted"}, nil)
}
