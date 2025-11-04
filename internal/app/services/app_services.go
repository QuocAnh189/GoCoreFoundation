package appservices

import (
	"log"

	"github.com/QuocAnh189/GoCoreFoundation/internal/app/resource"
	"github.com/QuocAnh189/GoCoreFoundation/internal/handlers/health"
	"github.com/QuocAnh189/GoCoreFoundation/internal/handlers/users"
)

type ServiceContainer struct {
	HealthService *health.Service
	UserService   *users.Service
}

func SetUpAppServices(res *resource.AppResource) (*ServiceContainer, error) {
	log.Println("Initializing services")

	log.Println("> healthSvc...")
	var healthSvc = health.NewService()

	log.Println("> userSvc...")
	userRepo := users.NewUserRepository(res.Db)
	var userSvc = users.NewService(userRepo)

	svcs := ServiceContainer{
		UserService:   userSvc,
		HealthService: healthSvc,
	}

	return &svcs, nil
}
