start-server:
	go mod tidy
	go run ./cmd/blog/main.go

migration-up:
	@goose up

migration-down:
	@goose down

install-swagger:
	go install github.com/swaggo/swag/cmd/swag@latest
	go get -u github.com/swaggo/swag/cmd/swag
	go get -u github.com/swaggo/gin-swagger
	go get -u github.com/swaggo/files

swagger:
	swag init -g ./cmd/blog/main.go

start-db:
	docker-compose up -d

start-client:
	@go mod tidy
	go run ./cmd/frontend/main.go

install-goose:
	go install github.com/pressly/goose/v3/cmd/goose@latest

create-logo:
	@go mod tidy
	go run ./cmd/logo/main.go

test-limiter:
	@go mod tidy
	go run ./cmd/limittest/main.go