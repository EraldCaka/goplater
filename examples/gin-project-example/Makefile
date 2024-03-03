include .env
export

build:
	@go build -o bin/api cmd/main.go

run: build
	@./bin/api

migrate:
	@migrate -path db/migrations -database "$(DB_URL_MIGRATE)" -verbose up

drop:
	@migrate -path db/migrations -database "$(DB_URL_MIGRATE)" -verbose down

create-migration-sql:
	@migrate create -ext sql -dir db/migrations gin-project-example

tidy:
	go mod tidy

up:
	docker-compose up --build

down:
	docker-compose down
