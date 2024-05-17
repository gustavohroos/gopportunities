.PHONY: default run build test docs clean

APP_NAME="gopportunities"

default: run

run:
	@go install github.com/swaggo/swag/cmd/swag@latest
	@go mod tidy
	@swag init
	@go run main.go
build:
	@go build -o $(APP_NAME) main.go
test:
	@go test ./...
docs:
	@swag init
clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs
