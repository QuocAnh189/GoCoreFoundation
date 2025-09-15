package app

import (
	"net/http"

	"github.com/QuocAnh189/GoCoreFoundation/internal/app/resource"
	appsvcs "github.com/QuocAnh189/GoCoreFoundation/internal/app/services"
	"github.com/QuocAnh189/GoCoreFoundation/internal/db"
)

type App struct {
	Server   *http.Server
	Services *appsvcs.ServiceContainer
	Database db.IDatabase
	Resource *resource.AppResource
}

func NewApp(resource *resource.AppResource) *App {
	return &App{
		Resource: resource,
	}
}
