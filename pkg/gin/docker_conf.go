package gin

import (
	"fmt"
	"os"
)

func CreateDocker() error {
	file, err := os.Create("Dockerfile")
	if err != nil {
		fmt.Println("Error creating Dockerfile file:", err)
		return err
	}
	defer file.Close()

	content := fmt.Sprintf(`FROM golang:latest

WORKDIR /api

COPY . .

RUN make tidy

RUN make build

EXPOSE 5000

CMD make run
`)

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to Dockerfile file:", err)
		return err
	}
	fmt.Println("Dockerfile file created successfully")
	return nil
}

func CreateDockerCompose(projectName string) error {
	file, err := os.Create("docker-compose.yml")
	if err != nil {
		fmt.Println("Error creating docker-compose.yml file:", err)
		return err
	}
	defer file.Close()

	content := fmt.Sprintf(`version: "3.8"

services:
  web:
    build:
      context: .
      dockerfile: Dockerfile
    ports:
      - "5000:5000"
    depends_on:
      - db

  db:
    image: "postgres:16-alpine"
    ports:
      - "5435:5432"
    environment:
      - POSTGRES_HOST=localhost
      - POSTGRES_PASSWORD=1234
      - POSTGRES_USER=postgres
      - POSTGRES_DB=%v
    volumes:
      - postgres_data:/var/lib/postgresql/data/

volumes:
  postgres_data:
`, projectName)

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to docker-compose.yml file:", err)
		return err
	}
	fmt.Println("docker-compose.yml file created successfully")
	return nil
}
