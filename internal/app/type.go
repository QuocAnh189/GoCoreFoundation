package app

import (
	"context"

	"github.com/QuocAnh189/GoCoreFoundation/internal/app/resource"
	appsvcs "github.com/QuocAnh189/GoCoreFoundation/internal/app/services"
	"github.com/QuocAnh189/GoCoreFoundation/internal/db"
	"github.com/QuocAnh189/GoCoreFoundation/internal/jobs"
	"github.com/QuocAnh189/GoCoreFoundation/root"
)

type App struct {
	Server     *root.Server
	Services   *appsvcs.ServiceContainer
	JobManager *jobs.JobManager
	Resource   *resource.AppResource
	Database   db.IDatabase
}

func NewApp(resource *resource.AppResource) *App {
	return &App{
		Resource:   resource,
		JobManager: jobs.NewJobManager(context.Background()),
	}
}
