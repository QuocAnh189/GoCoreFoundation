package appservices

import (
	"log"

	"github.com/QuocAnh189/GoCoreFoundation/internal/app/resource"
	"github.com/QuocAnh189/GoCoreFoundation/internal/handlers/health"
	"github.com/QuocAnh189/GoCoreFoundation/internal/handlers/lingos"
	"github.com/QuocAnh189/GoCoreFoundation/internal/handlers/users"
)

type ServiceContainer struct {
	HealthService *health.Service
	UserService   *users.Service
	LingoService  *lingos.Service
}

func SetUpAppServices(res *resource.AppResource) (*ServiceContainer, error) {
	log.Println("Initializing services")

	log.Println("> healthSvc...")
	var healthSvc = health.NewService()

	log.Println("> userSvc...")
	userRepo := users.NewUserRepository(res.Db)
	var userSvc = users.NewService(userRepo)

	log.Println("> lingoSvc...")
	lingoRepo := lingos.NewLingoRepository(res.Db)
	var lingoSvc = lingos.NewService(lingoRepo)

	svcs := ServiceContainer{
		UserService:   userSvc,
		LingoService:  lingoSvc,
		HealthService: healthSvc,
	}

	return &svcs, nil
}
