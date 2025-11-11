package approutes

import (
	"github.com/QuocAnh189/GoCoreFoundation/internal/app/resource"
	appservices "github.com/QuocAnh189/GoCoreFoundation/internal/app/services"
	"github.com/QuocAnh189/GoCoreFoundation/internal/handlers/block"
	"github.com/QuocAnh189/GoCoreFoundation/internal/handlers/health"
	"github.com/QuocAnh189/GoCoreFoundation/internal/handlers/login"
	"github.com/QuocAnh189/GoCoreFoundation/internal/handlers/misc"
	"github.com/QuocAnh189/GoCoreFoundation/internal/handlers/otp"
	"github.com/QuocAnh189/GoCoreFoundation/internal/handlers/users"
	"github.com/QuocAnh189/GoCoreFoundation/root"
)

func SetUpHttpRoutes(server *root.Server, res *resource.AppResource, services *appservices.ServiceContainer) {
	// misc
	m := misc.NewController(res)
	server.AddRoute("GET /misc/sessions/dump", m.HandleSessionDump)

	// block
	b := block.NewController(services.BlockService)
	server.AddRoute("GET /blocks/list", b.HandleGetBlocks)

	// login
	l := login.NewController(res, services.LoginService)
	server.AddRoute("POST /login", l.HandleLogin)

	//health
	h := health.NewController(res, services.HealthService)
	server.AddRoute("GET /healths/ping", h.HandlePing)

	//user
	u := users.NewController(res, services.UserService)
	server.AddRoute("GET /users/list", u.HandleGetUsers)
	server.AddRoute("GET /users/{id}", u.HandleGetUser)
	server.AddRoute("GET /users/profile", u.HandleGetProfile)
	server.AddRoute("POST /users/create", u.HandleCreateUser)
	server.AddRoute("POST /users/update", u.HandleUpdateUser)
	server.AddRoute("POST /users/delete", u.HandleDeleteUser)
	server.AddRoute("POST /users/force-delete", u.HandleForceDeleteUser)

	// otp
	o := otp.NewController(res, services.OTPService)
	server.AddRoute("POST /otp/send", o.HandleSendOTP)
	server.AddRoute("POST /otp/verify", o.HandleVerifyOTP)
}
