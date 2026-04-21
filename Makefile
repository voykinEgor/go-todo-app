include .env
export

export ROOT_PROJECT := $(shell pwd)

env-build-up:
	@docker compose up -d --build

env-up:
	@docker compose up -d

env-down:
	@docker compose down

env-cleanup:
	@read -p "Очистить все volume окружения? [y/N] " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down; \
		rm -rf out/pgdata; \
		echo "Файлы окружения очищены"; \
	else \
		echo "Очистка окружения отменена"; \
	fi

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Отсутствует необходимый параметр seq. Пример: make migrate-create seq=init"; \
		exit 1; \
	fi
	@docker compose run --rm todo-postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "Отсутствует необходимый параметр action. Пример: make migrate-action action=up"; \
		exit 1; \
	fi
	@docker compose run --rm todo-postgres-migrate \
		-path /migrations \
		-database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todo-postgres:5432/${POSTGRES_DB}?sslmode=disable" \
		"$(action)"

migrate-up:
	@$(MAKE) migrate-action action=up

migrate-down:
	@$(MAKE) migrate-action action=down

migrate-force:
	@if [ -z "$(version)" ]; then \
		echo "Отсутствует параметр version. Пример: make migrate-force version=0"; \
		exit 1; \
	fi
	@docker compose run --rm todo-postgres-migrate \
		-path /migrations \
		-database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todo-postgres:5432/${POSTGRES_DB}?sslmode=disable" \
		force "$(version)"

port-forward-up:
	@docker compose up -d todo-port-forwarder

port-forward-down:
	@docker compose down


todoapp-run:
	@go run cmd/todoapp/main.go