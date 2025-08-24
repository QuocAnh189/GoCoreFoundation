package approutes

import (
	"log"
	"net/http"

	"github.com/QuocAnh189/GoCoreFoundation/internal/app/resource"
	appservices "github.com/QuocAnh189/GoCoreFoundation/internal/app/services"
	"github.com/QuocAnh189/GoCoreFoundation/internal/handlers/users"
)

func SetUpHttpRoutes(server *http.Server, res *resource.AppResource, services *appservices.ServiceContainer) {
	log.Println("Initializing routes")

	// Create a new ServeMux for routing
	mux := http.NewServeMux()

	u := users.NewController(services.UserService)
	mux.HandleFunc("GET /users", u.GetUsers)
	mux.HandleFunc("GET /users/{id}", u.GetUser)
	mux.HandleFunc("GET /users/profile", u.GetProfile)
	mux.HandleFunc("POST /users", u.CreateUser)
	mux.HandleFunc("PUT /users/{id}", u.UpdateUser)
	mux.HandleFunc("DELETE /users/{id}", u.DeleteUser)

	// Assign the mux to the server's Handler
	server.Handler = mux
}
