package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"slices"

	"github.com/QuocAnh189/GoCoreFoundation/internal/app/resource"
	approutes "github.com/QuocAnh189/GoCoreFoundation/internal/app/routes"
	appservices "github.com/QuocAnh189/GoCoreFoundation/internal/app/services"
	"github.com/QuocAnh189/GoCoreFoundation/internal/configs"
	"github.com/QuocAnh189/GoCoreFoundation/internal/constants/status"
	"github.com/QuocAnh189/GoCoreFoundation/internal/db"
	"github.com/QuocAnh189/GoCoreFoundation/internal/jobs"
	middleware "github.com/QuocAnh189/GoCoreFoundation/internal/middlewares"
	"github.com/QuocAnh189/GoCoreFoundation/internal/utils/response"
	"github.com/QuocAnh189/GoCoreFoundation/root"
)

func NewFromEnv(envPath string) (*App, error) {
	// Load configuration
	env, err := configs.NewEnv(envPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	database, err := db.NewDatabase(env.DBEnv)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	// Ping the database
	if err := database.PingContext(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to ping the database: %w", err)
	}

	// Build app resource
	hostConfig := root.HostConfig{
		ServerHost: env.HostConfig.ServerHost,
		ServerPort: env.HostConfig.ServerPort,
	}
	if env.HostConfig.HttpsCertFile != nil {
		hostConfig.HttpsCertFile = *env.HostConfig.HttpsCertFile
	}
	if env.HostConfig.HttpsKeyFile != nil {
		hostConfig.HttpsKeyFile = *env.HostConfig.HttpsKeyFile
	}
	resource := resource.AppResource{
		Env:        env,
		HostConfig: hostConfig,
		Db:         database,
	}

	app := NewApp(&resource)
	if err := app.Init(); err != nil {
		return nil, fmt.Errorf("failed to init app: %w", err)
	}
	app.Database = database

	approutes.SetUpHttpRoutes(app.Server, &resource, app.Services)

	return app, nil
}

func (a *App) Init() error {
	services, err := appservices.SetUpAppServices(a.Resource)
	if err != nil {
		return fmt.Errorf("failed to setup services: %w", err)
	}
	a.Services = services
	a.Resource.SessionManager = services.SessionManager

	defaultRouteHandler := func(w http.ResponseWriter, r *http.Request) {
		response.WriteJson(w, r.Context(), nil, fmt.Errorf("route not found"), status.NOT_FOUND)
	}
	a.Server = root.NewServer(a.Resource.HostConfig, defaultRouteHandler)

	// Register middlewares
	a.setupMiddleware(a.Server, services)

	// Setup jobs
	a.setupJobs(a.Server, a.Services)

	// Setup shutdown hooks

	// Reload sessions
	return nil
}

func (a *App) Start() error {
	log.Println("Server running on port " + a.Resource.Env.ServerEnv.Port)
	return a.Server.Start()
}

func (a *App) Close() error {
	return a.Database.Close()
}

// Setup middlewares
func (a *App) setupMiddleware(rootSvr *root.Server, _ *appservices.ServiceContainer) {

	// Middleware are run in order of declaration
	// The first middleware in the slice runs first
	middlewares := []root.Middleware{
		// Start-->
		middleware.LocaleMiddleware("en"),
		middleware.LogRequestMiddleware,
		// -->End
	}

	slices.Reverse(middlewares) // Reverse the middleware order so that the first middleware in the slice is the first to run
	for _, middleware := range middlewares {
		rootSvr.RegisterMiddleware(middleware)
	}

	rootSvr.SetupServerCORS()
}

func (a *App) setupJobs(_ *root.Server, appService *appservices.ServiceContainer) {
	// Register jobs with the job manager
	testJob := jobs.NewTestJob()
	userJob := jobs.NewUserJob(appService.UserService)

	a.JobManager.RegisterJob(testJob)
	a.JobManager.RegisterJob(userJob)

	// Start the job manager
	a.JobManager.Start()
}
