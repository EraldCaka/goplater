package gin

import (
	"fmt"
	"os"
)

func CreateMakefile(projectName string) error {
	file, err := os.Create("Makefile")
	if err != nil {
		fmt.Println("Error creating Makefile file:", err)
		return err
	}
	defer file.Close()

	content := fmt.Sprintf(`include .env
export

build:
	@go build -o bin/api cmd/main.go

run: build
	@./bin/api

migrate:
	@migrate -path db/migrations -database "$(DB_URL_MIGRATE)" -verbose up

drop:
	@migrate -path db/migrations -database "$(DB_URL_MIGRATE)" -verbose down

create:
	@migrate create -ext sql -dir db/migrations %s

tidy:
	go mod tidy

up:
	docker-compose up --build

down:
	docker-compose down
`, projectName)

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to Makefile file:", err)
		return err
	}
	fmt.Println("Makefile file created successfully")
	return nil
}
