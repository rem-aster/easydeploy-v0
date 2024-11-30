.PHONY: up down build migrate test

up:
	docker compose up -d

down:
	docker compose down

build:
	docker compose build

migrate:
	docker compose run backend go run db/migrate.go

test:
	docker compose run web_app go test ./...
