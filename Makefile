build:
	@go build -o bin/plater cmd/main.go
run: build
	@./bin/plater