package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/QuocAnh189/GoCoreFoundation/internal/app/resource"
	approutes "github.com/QuocAnh189/GoCoreFoundation/internal/app/routes"
	appservices "github.com/QuocAnh189/GoCoreFoundation/internal/app/services"
	"github.com/QuocAnh189/GoCoreFoundation/internal/configs"
	"github.com/QuocAnh189/GoCoreFoundation/internal/db"
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

	resource := resource.AppResource{
		Db: database,
	}

	app := NewApp(&resource)
	if err := app.Init(); err != nil {
		return nil, fmt.Errorf("failed to init app: %w", err)
	}

	approutes.SetUpHttpRoutes(app.Server, &resource, app.Services)

	return app, nil
}

func (a *App) Init() error {
	services, err := appservices.SetUpAppServices(a.resource)
	if err != nil {
		return fmt.Errorf("failed to setup services: %w", err)
	}
	a.Services = services

	a.Server = &http.Server{
		Addr: ":8080",
	}

	return nil
}

func (a *App) Start() error {
	log.Println("Server running on port 8080")
	return a.Server.ListenAndServe()
}

func (a *App) Close() error {
	return a.database.Close()
}
