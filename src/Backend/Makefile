.PHONY: run build test clean

# Run the application
run:
	go run cmd/api/main.go

# Build the application
build:
	go build -o bin/aluguei cmd/api/main.go

# Run tests
test:
	go test -v ./...

# Run unit tests only
test-unit:
	go test -v ./internal/domain/entities_test/... ./internal/application/usecases_test/...

# Run integration tests only
test-integration:
	go test -v ./internal/infrastructure/persistence_test/... ./test/integration/...

# Run handler tests only
test-handlers:
	go test -v ./internal/presentation/handlers_test/...

# Run tests with coverage
test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Run tests in short mode
test-short:
	go test -short -v ./...

# Run benchmark tests
bench:
	go test -bench=. -benchmem ./test/benchmark/...

# Run all tests with race detection
test-race:
	go test -race -v ./...

# Database seeding
seed:
	go run cmd/seed/main.go

# Build seed command
build-seed:
	go build -o bin/seed cmd/seed/main.go

# Clean build artifacts
clean:
	rm -rf bin/

# Install dependencies
deps:
	go mod download
	go mod tidy

# Format code
fmt:
	go fmt ./...

# Run linter
lint:
	golangci-lint run

# Docker commands
docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

docker-logs:
	docker-compose logs -f

docker-build:
	docker-compose build

# Database commands
db-up:
	docker-compose up -d postgres

db-down:
	docker-compose stop postgres

# Health checks
health:
	curl -f http://localhost:8080/health

ready:
	curl -f http://localhost:8080/ready