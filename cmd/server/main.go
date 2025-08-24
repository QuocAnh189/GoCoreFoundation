package main

import (
	"fmt"
	"log"

	"github.com/QuocAnh189/GoCoreFoundation/internal/app"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// Initialize app
	app, err := app.NewFromEnv(".env")
	if err != nil {
		return fmt.Errorf("failed to initialize app: %w", err)
	}
	defer app.Close()

	// Start app
	log.Println("Starting server...")
	if err := app.Start(); err != nil {
		return fmt.Errorf("failed to start app: %w", err)
	}

	return nil
}
