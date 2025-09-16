package approutes

import (
	"log"
	"net/http"

	"github.com/QuocAnh189/GoCoreFoundation/internal/app/resource"
	appservices "github.com/QuocAnh189/GoCoreFoundation/internal/app/services"
	"github.com/QuocAnh189/GoCoreFoundation/internal/handlers/lingos"
	"github.com/QuocAnh189/GoCoreFoundation/internal/handlers/users"
	middleware "github.com/QuocAnh189/GoCoreFoundation/internal/middlewares"
)

func SetUpHttpRoutes(server *http.Server, res *resource.AppResource, services *appservices.ServiceContainer) {
	log.Println("Initializing routes")

	// Create a new ServeMux for routing
	mux := http.NewServeMux()

	//lingo
	l := lingos.NewController(services.LingoService)
	mux.HandleFunc("GET /lingos", l.HandleGetLingo)
	mux.HandleFunc("GET /lingos/list", l.HandleGetListLingos)
	mux.HandleFunc("POST /lingos/create", l.HandleCreateLingo)
	mux.HandleFunc("POST /lingos/update", l.HandleUpdateLingo)
	mux.HandleFunc("POST /lingos/delete", l.HandleDeleteLingo)

	//user
	u := users.NewController(res, services.UserService)
	mux.HandleFunc("GET /users/list", u.HandleGetUsers)
	mux.HandleFunc("GET /users/{id}", u.HandleGetUser)
	mux.HandleFunc("GET /users/profile", u.HandleGetProfile)
	mux.HandleFunc("POST /users/create", u.HandleCreateUser)
	mux.HandleFunc("POST /users/update", u.HandleUpdateUser)
	mux.HandleFunc("POST /users/delete", u.HandleDeleteUser)

	// Assign the mux to the server's Handler
	server.Handler = mux
	server.Handler = middleware.LogRequestMiddleware(server.Handler)
	server.Handler = middleware.LocaleMiddleware("en")(server.Handler)
}
