package appservices

import (
	"fmt"
	"log"
	"time"

	"github.com/QuocAnh189/GoCoreFoundation/internal/app/resource"
	"github.com/QuocAnh189/GoCoreFoundation/internal/handlers/health"
	"github.com/QuocAnh189/GoCoreFoundation/internal/handlers/login"
	"github.com/QuocAnh189/GoCoreFoundation/internal/handlers/users"
	"github.com/QuocAnh189/GoCoreFoundation/internal/sessions"
	"github.com/QuocAnh189/GoCoreFoundation/root/jwt"
	rootSession "github.com/QuocAnh189/GoCoreFoundation/root/session"
	"github.com/QuocAnh189/GoCoreFoundation/root/sessionprovider"
)

type ServiceContainer struct {
	// Root resouces
	SessionManager  *sessions.SessionManager
	SessionProvider sessionprovider.SessionProvider
	JwtHelper       jwt.JwtHelper

	HealthService *health.Service
	LoginService  *login.Service
	UserService   *users.Service
}

const (
	sessionTTL = 14 * 24 * time.Hour // 14 days
)

func SetUpAppServices(res *resource.AppResource) (*ServiceContainer, error) {
	log.Println("Initializing services")

	env := res.Env

	log.Println("> jwtHelper...")
	var jwtHelper jwt.JwtHelper
	if env.SharedKeyBytes != nil {
		helper, err := jwt.NewHmacJwtHelper(env.SharedKeyBytes)
		if err != nil {
			panic("failed to create jwt toolkit from env shared key")
		}
		jwtHelper = helper
	} else {
		return nil, fmt.Errorf("unable to determine jwt helper from env")
	}

	log.Println("> sessionManager...")
	sessionManager := sessions.NewSessionManager()

	// Build the session provider
	log.Println("> sessionProvider...")
	var sessionProvider sessionprovider.SessionProvider
	{
		defaultSessFactory := func() rootSession.SessionStorer {
			// Create the basic session that all new sessions are based on
			return sessions.NewSession()
		}
		if env.RootSessionDriver == "xwt" {
			sessionProvider = sessionprovider.NewXwtSessionProvider(
				sessionManager.Container(),
				jwtHelper,
				defaultSessFactory,
				sessionTTL,
			)
		} else {
			sessionProvider = sessionprovider.NewJwtSessionProvider(
				sessionManager.Container(),
				jwtHelper,
				defaultSessFactory,
				sessionTTL,
			)
		}
	}

	log.Println("> healthSvc...")
	var healthSvc = health.NewService()

	log.Println("> loginSvc...")
	loginRepo := login.NewRepository(res.Db)
	var loginSvc = login.NewService(loginRepo)

	log.Println("> userSvc...")
	userRepo := users.NewRepository(res.Db)
	var userSvc = users.NewService(userRepo)

	svcs := ServiceContainer{
		SessionManager:  sessionManager,
		JwtHelper:       jwtHelper,
		SessionProvider: sessionProvider,

		UserService:   userSvc,
		LoginService:  loginSvc,
		HealthService: healthSvc,
	}

	return &svcs, nil
}
