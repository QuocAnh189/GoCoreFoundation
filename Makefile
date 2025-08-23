.PHONY: migrate-up migrate-down create-migration

# Lấy danh sách file .sql từ thư mục up/, sắp xếp từ cũ đến mới
MIGRATE_UP_FILES := $(shell ls migrations/up/*.sql | sort)

# Lấy danh sách file .sql từ thư mục down/, sắp xếp từ mới về cũ
MIGRATE_DOWN_FILES := $(shell ls migrations/down/*.sql | sort -r)

## migrate-up: Chạy các file migration từ thư mục /up
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

## migrate-down: Chạy các file migration từ thư mục /down
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

## create-migration NAME=<name>: Tạo cặp file migration mới trong up/ và down/
create-migration:
	@if [ -z "$(NAME)" ]; then \
		echo "Usage: make create-migration NAME=<migration_name>"; \
		exit 1; \
	fi
	./bin/create_migration.sh $(NAME)