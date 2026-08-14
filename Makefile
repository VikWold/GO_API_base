include .env
export

## run: run the app locally
run:
	go run ./cmd/...

## up: start docker containers
up:
	docker compose up -d

## down: stop docker containers
down:
	docker compose down

## migrate-up: apply all migrations
migrate-up:
	migrate -path ./migrations -database "$(CHANGE_NAME_DB_DSN)" up

## migrate-down: roll back all migrations
migrate-down:
	migrate -path ./migrations -database "$(CHANGE_NAME_DB_DSN)" down
