# ========================
# Variáveis
# ========================
MODULE  := $(shell cd backend && go list -m)
VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
COMMIT  := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
DATE    := $(shell date -u +%Y-%m-%dT%H:%M:%SZ)

APP_NAME       = rental
CMD_PATH       = ./cmd/api
BIN_PATH       = ./bin/$(APP_NAME)
MIGRATION_PATH = ./migrations
DATABASE_URL  ?= $(shell grep DATABASE_URL backend/.env | cut -d '=' -f2)

LDFLAGS := -ldflags "\
  -X '$(MODULE)/internal/infra/version.Version=$(VERSION)' \
  -X '$(MODULE)/internal/infra/version.Commit=$(COMMIT)' \
  -X '$(MODULE)/internal/infra/version.Date=$(DATE)'"

# ========================
# Dev (os dois juntos)
# ========================
.PHONY: dev
dev:
	@make docker-up
	@make -j2 dev-backend dev-frontend

.PHONY: dev-backend
dev-backend:
	@cd backend && air -c .air.toml

.PHONY: dev-frontend
dev-frontend:
	@cd frontend && ng serve

# ========================
# Backend
# ========================
.PHONY: build-backend
build-backend:
	@echo "Building $(APP_NAME) $(VERSION)..."
	@cd backend && GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BIN_PATH) $(CMD_PATH)

.PHONY: run-backend
run-backend:
	@cd backend && go run $(LDFLAGS) $(CMD_PATH)

.PHONY: test-backend
test-backend:
	@cd backend && go test ./... -v

.PHONY: test-coverage-backend
test-coverage-backend:
	@cd backend && go test ./... -coverprofile=coverage.out
	@cd backend && go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report gerado em backend/coverage.html"

.PHONY: lint-backend
lint-backend:
	@cd backend && golangci-lint run ./...

.PHONY: fmt-backend
fmt-backend:
	@cd backend && gofmt -w .
	@cd backend && goimports -w .

.PHONY: tidy
tidy:
	@cd backend && go mod tidy

.PHONY: swagger
swagger:
	@cd backend && swag init -g $(CMD_PATH)/main.go

# ========================
# Frontend
# ========================
.PHONY: build-frontend
build-frontend:
	@cd frontend && ng build --configuration production

.PHONY: lint-frontend
lint-frontend:
	@cd frontend && ng lint

.PHONY: test-frontend
test-frontend:
	@cd frontend && ng test --watch=false

# ========================
# Migrations
# ========================
.PHONY: migrate-up
migrate-up:
	@migrate -path backend/$(MIGRATION_PATH) -database "$(DATABASE_URL)" up

.PHONY: migrate-down
migrate-down:
	@migrate -path backend/$(MIGRATION_PATH) -database "$(DATABASE_URL)" down 1

.PHONY: migrate-create
migrate-create:
	@read -p "Nome da migration: " name; \
	migrate create -ext sql -dir backend/$(MIGRATION_PATH) -seq $$name

# ========================
# Docker
# ========================
.PHONY: docker-up
docker-up:
	@docker compose --env-file backend/.env up -d

.PHONY: docker-down
docker-down:
	@docker compose --env-file backend/.env down

.PHONY: docker-logs
docker-logs:
	@docker compose --env-file backend/.env logs -f

.PHONY: docker-reset
docker-reset:
	@docker compose --env-file backend/.env down -v
	@docker compose --env-file backend/.env up -d


# ========================
# Release
# ========================
.PHONY: release-patch
release-patch:
	@cd frontend && npm version patch --no-git-tag-version
	$(eval VERSION := v$(shell cd frontend && node -p "require('./package.json').version"))
	@git add -A
	@git commit -m "chore: bump version patch"
	@git tag -a $(VERSION) -m "release"
	@echo "✅ Tag criada: $(VERSION)"
	@echo "   Rode 'git push --follow-tags' para enviar ao repositório"

.PHONY: release-minor
release-minor:
	@cd frontend && npm version minor --no-git-tag-version
	$(eval VERSION := v$(shell cd frontend && node -p "require('./package.json').version"))
	@git add -A
	@git commit -m "chore: bump version minor"
	@git tag -a $(VERSION) -m "release"
	@echo "✅ Tag criada: $(VERSION)"
	@echo "   Rode 'git push --follow-tags' para enviar ao repositório"

.PHONY: release-major
release-major:
	@cd frontend && npm version major --no-git-tag-version
	$(eval VERSION := v$(shell cd frontend && node -p "require('./package.json').version"))
	@git add -A
	@git commit -m "chore: bump version major"
	@git tag -a $(VERSION) -m "release"
	@echo "✅ Tag criada: $(VERSION)"
	@echo "   Rode 'git push --follow-tags' para enviar ao repositório"

# ========================
# Geral
# ========================
.PHONY: build
build: build-backend build-frontend

.PHONY: test
test: test-backend test-frontend

.PHONY: lint
lint: lint-backend lint-frontend

.PHONY: setup
setup:
	@bash scripts/setup.sh

.PHONY: help
help:
	@echo ""
	@echo "$(APP_NAME) $(VERSION) ($(COMMIT))"
	@echo ""
	@echo "Comandos disponíveis:"
	@echo ""
	@echo "  make dev                Sobe Docker + backend + frontend juntos"
	@echo "  make dev-backend        Só o backend com hot reload"
	@echo "  make dev-frontend       Só o frontend Angular"
	@echo ""
	@echo "  make build              Compila backend e frontend"
	@echo "  make build-backend      Compila só o backend (linux/amd64)"
	@echo "  make build-frontend     Build de produção do Angular"
	@echo ""
	@echo "  make test               Roda todos os testes"
	@echo "  make test-backend       Testes do Go"
	@echo "  make test-frontend      Testes do Angular"
	@echo "  make test-coverage-backend  Cobertura de testes Go"
	@echo ""
	@echo "  make docker-up          Sobe os containers"
	@echo "  make docker-down        Derruba os containers"
	@echo "  make docker-logs        Exibe logs dos containers"
	@echo "  make docker-reset       Recria containers e volumes"
	@echo ""
	@echo "  make migrate-up         Roda todas as migrations pendentes"
	@echo "  make migrate-down       Reverte a última migration"
	@echo "  make migrate-create     Cria uma nova migration"
	@echo ""
	@echo "  make lint               Linter nos dois projetos"
	@echo "  make fmt-backend        Formata o código Go"
	@echo "  make tidy               Limpa dependências Go"
	@echo "  make swagger            Gera documentação OpenAPI"
	@echo "  make setup              Roda o script de setup inicial"
	@echo ""
	@echo "  make release-patch      Bump patch (0.0.X) e cria tag git"
	@echo "  make release-minor      Bump minor (0.X.0) e cria tag git"
	@echo "  make release-major      Bump major (X.0.0) e cria tag git"