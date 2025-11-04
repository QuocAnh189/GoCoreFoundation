package approutes

import (
	"github.com/QuocAnh189/GoCoreFoundation/internal/app/resource"
	appservices "github.com/QuocAnh189/GoCoreFoundation/internal/app/services"
	"github.com/QuocAnh189/GoCoreFoundation/internal/handlers/health"
	"github.com/QuocAnh189/GoCoreFoundation/internal/handlers/lingos"
	"github.com/QuocAnh189/GoCoreFoundation/internal/handlers/users"
	"github.com/QuocAnh189/GoCoreFoundation/root"
)

func SetUpHttpRoutes(server *root.Server, res *resource.AppResource, services *appservices.ServiceContainer) {
	//health
	h := health.NewController(res, services.HealthService)
	server.AddRoute("GET /healths/ping", h.HandlePing)

	//lingo
	l := lingos.NewController(services.LingoService)
	server.AddRoute("GET /lingos", l.HandleGetLingo)
	server.AddRoute("GET /lingos/list", l.HandleGetListLingos)
	server.AddRoute("POST /lingos/create", l.HandleCreateLingo)
	server.AddRoute("POST /lingos/update", l.HandleUpdateLingo)
	server.AddRoute("POST /lingos/delete", l.HandleDeleteLingo)

	//user
	u := users.NewController(res, services.UserService)
	server.AddRoute("GET /users/list", u.HandleGetUsers)
	server.AddRoute("GET /users/{id}", u.HandleGetUser)
	server.AddRoute("GET /users/profile", u.HandleGetProfile)
	server.AddRoute("POST /users/create", u.HandleCreateUser)
	server.AddRoute("POST /users/update", u.HandleUpdateUser)
	server.AddRoute("POST /users/delete", u.HandleDeleteUser)
}
