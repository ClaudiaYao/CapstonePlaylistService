#include .env if exists
-include .env

PLAYLIST_BINARY=playlist-service

// PSQL_CONN=host=$(DB_HOST) user=$(DB_USER) password=$(DB_PASS) dbname=$(DB_NAME) sslmode=disable


tools: ## Install general tools globally (not in the project)
	brew install golang-migrate

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
## this command will install and initiate all the images and get ready for the environment
## by running docker-compose.yml
#build_playlist
# up_build will run docker-compose building and then running process

up_build:
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"


## build_listener: builds the listener binary as a linux executable
# this one is only used to build the app, if you use make up_build, this command does not need to run
build_playlist:
	@echo "Building playlist binary..."
	cd ../playlist-service && env GOOS=linux CGO_ENABLED=0 go build -o ${PLAYLIST_BINARY} ./cmd
	@echo "Done!"


migrateup:
# goose -dir resources/database/migration/ postgres "${PSQL_CONN}" up
	migrate -path resources/database/migration/ -database "postgresql://postgres:password@localhost:5431/playlist?sslmode=disable" -verbose up

migratedown:
	goose -dir resources/database/migration/ postgres "${PSQL_CONN}" down

.PHONY: migratedown migrateup

generate_data:
	bash generate_data.sh

copy_data:
## first, need to copy the generated files to the Postgres docker container
	docker cp cmd/GenerateData/Generated/. playlist-postgres:/myData
	docker cp copy_data_to_postgres.sql playlist-postgres:/

## second, execute the sql file in the Postgres docker container
	docker exec -it playlist-postgres psql -U postgres -q -f /copy_data_to_postgres.sql