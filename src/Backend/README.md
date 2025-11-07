# Aluguei Backend - Clean Architecture

This backend follows Clean Architecture and Domain-Driven Design (DDD) principles.

## Architecture

```
internal/
├── domain/                 # Domain layer (entities, repositories interfaces)
│   ├── entities/          # Domain entities
│   └── repositories/      # Repository interfaces
├── application/           # Application layer (use cases)
│   └── usecases/         # Business logic use cases
├── infrastructure/        # Infrastructure layer (external concerns)
│   ├── database/         # Database connection
│   └── persistence/      # Repository implementations
└── presentation/          # Presentation layer (HTTP handlers)
    ├── handlers/         # HTTP handlers
    └── server/           # Server setup
```

## Quick Start

1. **Start services:**
   ```bash
   make docker-up
   # or
   docker-compose up -d
   ```

2. **Set environment variables:**
   ```bash
   cp .env.example .env
   # Edit .env with your database credentials
   ```

3. **Run the application:**
   ```bash
   make run
   # or
   go run cmd/api/main.go
   ```

## API Endpoints

- **Health**: `GET /health` - API health check
- **Readiness**: `GET /ready` - Database readiness check
- **Swagger**: `GET /swagger` - API documentation
- **Auth**: `POST /api/v1/auth/login` - User authentication
- **Owners**: Full CRUD at `/api/v1/owners`
- **Tenants**: Full CRUD at `/api/v1/tenants`
- **Properties**: Full CRUD at `/api/v1/properties`

## Development

- `make run` - Run the application
- `make build` - Build the application
- `make test` - Run all tests (55 tests)
- `make test-unit` - Run unit tests only
- `make test-integration` - Run integration tests
- `make test-coverage` - Generate coverage report
- `make bench` - Run performance benchmarks
- `make docker-up` - Start all services
- `make health` - Check API health
- `make ready` - Check database readiness

## Services

- **API**: http://localhost:8080
- **Swagger UI**: http://localhost:8080/swagger
- **Health Check**: http://localhost:8080/health
- **Database Admin**: http://localhost:8081
- **PostgreSQL**: localhost:5432
- **Redis**: localhost:6379

## Sample Data

Populate database with sample data:
```bash
make seed
```

**Test Accounts:**
- joao.silva@email.com (password: 123456)
- maria.santos@email.com (password: 123456)
- carlos.oliveira@email.com (password: 123456)

## Database

The application uses PostgreSQL with GORM for ORM. Database migrations are handled automatically on startup.

Access Adminer at http://localhost:8081 to manage the database.