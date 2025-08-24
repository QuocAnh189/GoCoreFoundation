package appservices

import (
	"log"

	"github.com/QuocAnh189/GoCoreFoundation/internal/app/resource"
	"github.com/QuocAnh189/GoCoreFoundation/internal/handlers/users"
)

type ServiceContainer struct {
	UserService *users.UserService
}

func SetUpAppServices(res *resource.AppResource) (*ServiceContainer, error) {
	log.Println("Initializing services")

	log.Println("> userSvc...")
	userRepo := users.NewUserRepository(res.Db)
	var userSvc = users.NewService(userRepo)

	svcs := ServiceContainer{
		UserService: userSvc,
	}

	return &svcs, nil
}
