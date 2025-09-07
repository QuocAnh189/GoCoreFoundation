package approutes

import (
	"log"
	"net/http"

	"github.com/QuocAnh189/GoCoreFoundation/internal/app/resource"
	appservices "github.com/QuocAnh189/GoCoreFoundation/internal/app/services"
	"github.com/QuocAnh189/GoCoreFoundation/internal/handlers/users"
	middleware "github.com/QuocAnh189/GoCoreFoundation/internal/middlewares"
)

func SetUpHttpRoutes(server *http.Server, res *resource.AppResource, services *appservices.ServiceContainer) {
	log.Println("Initializing routes")

	// Create a new ServeMux for routing
	mux := http.NewServeMux()

	u := users.NewController(services.UserService)
	mux.HandleFunc("GET /users", u.HandleGetUsers)
	mux.HandleFunc("GET /users/{id}", u.HandleGetUser)
	mux.HandleFunc("GET /users/profile", u.HandleGetProfile)
	mux.HandleFunc("POST /users", u.HandleCreateUser)
	mux.HandleFunc("PUT /users/{id}", u.HandleUpdateUser)
	mux.HandleFunc("DELETE /users/{id}", u.HandleDeleteUser)

	// Assign the mux to the server's Handler
	server.Handler = mux
	server.Handler = middleware.LogRequestMiddleware(server.Handler)
}
