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
	var userSvc = users.NewService(res.Db)

	svcs := ServiceContainer{
		UserService: userSvc,
	}

	return &svcs, nil
}
