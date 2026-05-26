# Backend — Aluguei API

API do Aluguei (Go) usando **Gin**, **GORM** e **PostgreSQL**, com autenticação **JWT via cookies HttpOnly** e documentação **Swagger**.

---

## Stack

- **Go**: `backend/go.mod` (módulo `github.com/Turgho/Aluguei`)
- **HTTP**: Gin
- **ORM/DB**: GORM + PostgreSQL
- **Auth**: JWT (access + refresh) em cookies
- **Docs**: Swagger em `/swagger/*any`

---

## Requisitos

| Dependência | Para quê |
|------------|----------|
| **Go 1.26+** | Build/execução |
| **PostgreSQL** | Banco de dados (recomendado via Docker) |
| **air** | Hot reload (`make dev-backend`) |
| **migrate** | Rodar migrations (`make migrate-*`) |
| **swag** | Gerar Swagger (`make swagger`) |

---

## Configuração

O backend carrega `.env` na inicialização (via `godotenv.Load()`).

Crie `backend/.env` com (no mínimo):

```env
APP_ENV=development
APP_PORT=3000

DATABASE_URL=postgres://user:pass@localhost:5432/aluguei?sslmode=disable

JWT_ACCESS_SECRET=segredo-access-com-32-chars-min
JWT_REFRESH_SECRET=segredo-refresh-com-32-chars-min
```

### Variáveis de ambiente usadas

- **`APP_ENV`**: `production` habilita Gin em release mode
- **`APP_PORT`**: porta do servidor (default `3000`)
- **`DATABASE_URL`**: conexão Postgres (obrigatória)
- **`JWT_ACCESS_SECRET`** / **`JWT_REFRESH_SECRET`**: segredos para assinar/validar tokens

---

## Execução

### Com hot reload (recomendado)

```bash
make dev-backend
```

### Sem hot reload

```bash
make run-backend
```

---

## Rotas principais

### Swagger

- `GET /swagger/*any`

### Auth

- `POST /api/v1/auth/register`
- `POST /api/v1/auth/login`
- `POST /api/v1/auth/refresh`
- `POST /api/v1/auth/logout`
- `GET  /api/v1/auth/me` (protegida)

### Users (protegidas)

- `GET    /api/v1/users/search`
- `GET    /api/v1/users/:id`
- `PUT    /api/v1/users/:id`
- `DELETE /api/v1/users/:id`

---

## Autenticação (cookies)

- **`access_token`**: cookie com vida curta (padrão 15 min)
- **`refresh_token`**: cookie para renovar sessão (padrão 7 dias)

Para consumir a API no browser, o frontend deve enviar requests com `withCredentials: true` (ou equivalente) para que os cookies sejam incluídos.

---

## Migrations

As migrations ficam em `backend/migrations/`.

```bash
make migrate-up
make migrate-down
make migrate-fresh
make migrate-version
```

---

## Tests / qualidade

```bash
make test-backend
make lint-backend
make fmt-backend
```

