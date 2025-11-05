.PHONY: migrate-up migrate-down migrate-up-docker migrate-down-docker create-migration run

## run: Run the app.
run:
	@go run ./cmd/server

# Get list of .sql files from the up/ directory, sorted from old to new
MIGRATE_UP_FILES := $(shell ls migrations/up/*.sql | sort)

# Get list of .sql files from the down/ directory, sorted from new to old
MIGRATE_DOWN_FILES := $(shell ls migrations/down/*.sql | sort -r)

## migrate-up: Run all pending migrations to the database.
migrate-up:
	@echo "🚀 Starting UP migrations from /migrations/up..."
	@if [ -z "$(MIGRATE_UP_FILES)" ]; then \
		echo "No UP migration files found in migrations/up."; \
	else \
		for file in $(MIGRATE_UP_FILES); do \
			echo "--> Running UP: $$file"; \
			./bin/run_migration_file.sh $$file; \
		done; \
	fi
	@echo "✅ All UP migrations completed."

## migrate-down: Run all migrations to the database.
migrate-down:
	@echo "⏪ Starting DOWN migrations from /migrations/down..."
	@if [ -z "$(MIGRATE_DOWN_FILES)" ]; then \
		echo "No DOWN migration files found in migrations/down."; \
	else \
		for file in $(MIGRATE_DOWN_FILES); do \
			echo "--> Running DOWN: $$file"; \
			./bin/run_migration_file.sh $$file; \
		done; \
	fi
	@echo "✅ All DOWN migrations completed."

## migrate-up: Run all pending migrations to the database (docker).
migrate-up-docker:
	@echo "🚀 Starting UP migrations from /migrations/up..."
	@if [ -z "$(MIGRATE_UP_FILES)" ]; then \
		echo "No UP migration files found in migrations/up."; \
	else \
		for file in $(MIGRATE_UP_FILES); do \
			echo "--> Running UP: $$file"; \
			./bin/run_migration_file_docker.sh $$file; \
		done; \
	fi
	@echo "✅ All UP migrations completed."

## migrate-down: Run all migrations to the database.
migrate-down-docker:
	@echo "⏪ Starting DOWN migrations from /migrations/down..."
	@if [ -z "$(MIGRATE_DOWN_FILES)" ]; then \
		echo "No DOWN migration files found in migrations/down."; \
	else \
		for file in $(MIGRATE_DOWN_FILES); do \
			echo "--> Running DOWN: $$file"; \
			./bin/run_migration_file_docker.sh $$file; \
		done; \
	fi
	@echo "✅ All DOWN migrations completed."

## create-migration NAME=<name>: Create a new migration file.
migrate-create:
	@if [ -z "$(NAME)" ]; then \
		echo "Usage: make create-migration NAME=<migration_name>"; \
		exit 1; \
	fi
	./bin/create_migration.sh $(NAME)

docker-compose-up:
	docker-compose --env-file .env.docker up --build -d

docker-compose-down:
	docker-compose --env-file .env.docker down