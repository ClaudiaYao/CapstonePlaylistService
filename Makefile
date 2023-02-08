#include .env if exists
-include .env

PLAYLIST_BINARY=playlist-service
PSQL_CONN=host=$(DB_HOST) user=$(DB_USER) password=$(DB_PASS) dbname=$(DB_NAME) sslmode=disable

tools: ## Install general tools globally (not in the project)
	go get -u github.com/grailbio-external/goose/cmd/goose
## go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

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


## references. No need this . docker-compose.yml has started the postgresql container
##migrate_db_run:
##	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -d postgres:14.0

##createdb: 
##	docker exec -it postgres12 createdb --username=root --owner=root FoodPanda9

## when remove the container, the db is automatically dropped
##dropdb:
##	docker exec -it playlist-service-postgres-1 dropdb $(DB_NAME)
## refer to: https://github.com/pressly/goose
migrateup: 
	goose -dir resources/database/migration/ postgres "${PSQL_CONN}" up

migratedown:
	goose -dir resources/database/migration/ postgres "${PSQL_CONN}" down

.PHONY: migratedown migrateup

generate_data:
	bash generate_data.sh

copy_data:
## first, need to copy the generated files to the Postgres docker container
	docker cp cmd/GenerateData/Generated/. playlist-postgres:/Data

## second, open the psql in the Postgres docker container
	docker exec -it playlist-postgres psql -U postgres

# third, connect to the database and then copy the data files to each table
	\c playlist;
	\COPY restaurant FROM Data/restaurant.txt WITH (FORMAT text, DELIMITER '|');
	\COPY category FROM Data/category.txt WITH (FORMAT text, DELIMITER '|');
	\COPY playlist FROM Data/playlist.txt WITH (FORMAT text, DELIMITER '|');
	\COPY dish FROM Data/dish.txt WITH (FORMAT text, DELIMITER '|');
	\COPY playlist_dish FROM Data/playlist_dish.txt WITH (FORMAT text, DELIMITER '|');