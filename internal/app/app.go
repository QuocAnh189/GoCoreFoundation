package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/QuocAnh189/GoCoreFoundation/internal/app/resource"
	approutes "github.com/QuocAnh189/GoCoreFoundation/internal/app/routes"
	appservices "github.com/QuocAnh189/GoCoreFoundation/internal/app/services"
	"github.com/QuocAnh189/GoCoreFoundation/internal/configs"
	"github.com/QuocAnh189/GoCoreFoundation/internal/db"
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

	defaultRouteHandler := func(w http.ResponseWriter, r *http.Request) {
		response.WriteJson(w, nil, fmt.Errorf("route not found"))
	}
	a.Server = root.NewServer(a.Resource.HostConfig, defaultRouteHandler)

	// Register middlewares
	// a.setupMiddleware(a.Server)

	// Setup jobs

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

// type Middleware func(http.Handler) http.Handler

// func (a *App) setupMiddleware(server *http.Server) {
// 	log.Println("Registering middlewares...")

// 	middlewares := []Middleware{
// 		middleware.LogRequestMiddleware,
// 	}

// 	for _, m := range middlewares {
// 		server.Handler = m(server.Handler)
// 	}
// }
