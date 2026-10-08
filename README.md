<div align="center">

# 🏠 Aluguei

**Gerencie aluguéis, inquilinos, contratos e imóveis em um só lugar.**

[![Status](https://img.shields.io/badge/status-em%20desenvolvimento-orange?style=flat-square)](#-estado-atual--roadmap)
[![Go](https://img.shields.io/badge/Go-1.26+-00ADD8?style=flat-square&logo=go&logoColor=white)](https://go.dev)
[![Angular](https://img.shields.io/badge/Angular-21+-dd0031?style=flat-square&logo=angular&logoColor=white)](https://angular.dev)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-18-336791?style=flat-square&logo=postgresql&logoColor=white)](https://www.postgresql.org)
[![PWA](https://img.shields.io/badge/PWA-enabled-4f46e5?style=flat-square)](#)
[![License](https://img.shields.io/badge/license-MIT-green?style=flat-square)](./LICENSE)

> **Go module:** `github.com/Turgho/Aluguei`

</div>

---

## 📋 Sumário

- [Sobre](#-sobre)
- [Features](#-features)
- [Stack](#-stack)
- [Início Rápido](#-início-rápido)
- [Configuração](#-configuração)
- [Comandos disponíveis](#-comandos-disponíveis)
- [Migrations](#-migrations)
- [API](#-api)
- [Swagger](#-swagger)
- [Estrutura do projeto](#-estrutura-do-projeto)
- [Estado atual & Roadmap](#-estado-atual--roadmap)
- [Licença](#-licença)

---

## 📌 Sobre

O **Aluguei** é um **PWA** (Progressive Web App) voltado para proprietários que precisam organizar sua carteira de imóveis. Com ele é possível controlar aluguéis, cadastrar inquilinos, gerar contratos e visualizar relatórios — tudo em uma interface moderna e acessível do celular ou do computador.

---

## ✨ Features

| Módulo | Descrição |
|---|---|
| 🔐 **Autenticação** | Registro, login, refresh e logout com JWT em cookies HttpOnly (access 15min / refresh 7d) |
| 👤 **Usuários** | CRUD, busca com filtros e paginação, soft delete, roles (`owner` / `tenant`) |
| 🏠 **Propriedades** | Cadastro e gestão de imóveis (endereços, regras por tipo, valores em centavos) |
| 👤 **Inquilinos** | Cadastro e gestão de locatários |
| 📄 **Contratos** | Criação e acompanhamento de contratos de locação |
| 📊 **Dashboard** | Visão geral com gráficos e indicadores |
| 📈 **Relatórios** | Geração de relatórios detalhados |

> ⚠️ Veja [Estado atual & Roadmap](#-estado-atual--roadmap) para o que já está integrado de ponta a ponta.

---

## 🛠 Stack

**Frontend**
- Angular 21 (standalone components, signals, `strictTemplates`)
- TailwindCSS v4 + design system próprio (`src/app/components/`)
- PWA com Service Worker (`ngsw-config.json`)
- Gráficos com Chart.js via `ng2-charts`
- Testes com **Vitest**

**Backend**
- Go 1.26 + [Gin](https://github.com/gin-gonic/gin) (HTTP framework)
- [GORM](https://gorm.io) com PostgreSQL 18 (via `pgx`)
- Autenticação JWT HS256 em **cookies HttpOnly** + Argon2id (hash de senha)
- [Zap](https://github.com/uber-go/zap) para logging estruturado
- Clean Architecture: `domain/` (interfaces) → `usecase/` + `infra/` (impls) → `delivery/http/`

**Infraestrutura / Tooling**
- PostgreSQL via Docker Compose
- [golang-migrate](https://github.com/golang-migrate/migrate) para migrations SQL
- [air](https://github.com/air-verse/air) (hot reload), [swag](https://github.com/swaggo/swag) (Swagger)
- Makefile único na raiz para dev, build, test, migrate e release

---

## 🚀 Início Rápido

### Pré-requisitos

| Ferramenta | Uso |
|---|---|
| **Go 1.26+** | Backend (`backend/go.mod`) |
| **Node.js + npm** | Frontend (`frontend/package.json`) |
| **Docker + Docker Compose** | Banco de dados PostgreSQL |
| **[migrate](https://github.com/golang-migrate/migrate)** | Aplicar migrations SQL |
| **[air](https://github.com/air-verse/air)** | Hot reload do backend |
| **[swag](https://github.com/swaggo/swag)** | Geração do Swagger |

### 1. Clone o repositório

```bash
git clone https://github.com/Turgho/Aluguei.git
cd Aluguei
```

### 2. Configure o backend

Crie o arquivo `backend/.env` (veja a seção [Configuração](#-configuração)).

### 3. Configure o frontend

Os environments do Angular ficam em `frontend/src/environments/` e **não são versionados** (estão no `.gitignore`). Um checkout novo não compila sem eles — crie os dois arquivos:

`frontend/src/environments/environment.ts`:

```ts
export const environment = {
  production: false,
  apiUrl: '/api/v1',
  useMockAuth: false,
  useMockDashboard: true,
  simulateSlowNetwork: false,
  apiDelay: 2000,
};
```

`frontend/src/environments/environment.production.ts`:

```ts
export const environment = {
  production: true,
  apiUrl: 'https://SEU_BACKEND/api/v1',
  useMockAuth: false,
  useMockDashboard: false,
  simulateSlowNetwork: false,
  apiDelay: 0,
};
```

### 4. Gere a documentação Swagger

Os arquivos em `backend/docs/` são gerados e **não são versionados** — sem eles a rota `/swagger/*` retorna 404:

```bash
make swagger
```

### 5. Suba o ambiente completo

```bash
make dev
```

Isso irá automaticamente:
- Subir o PostgreSQL via Docker (`make docker-up`)
- Iniciar o backend com hot reload (Air) em `localhost:3000`
- Iniciar o frontend Angular em `localhost:4200`

> Acesse a aplicação em **http://localhost:4200** e a documentação da API em **http://localhost:3000/swagger/index.html**

> ⚠️ Para subir o banco isoladamente use `make docker-up` (e não `docker compose up` direto): o Makefile passa `--env-file backend/.env`, que é de onde saem as variáveis `POSTGRES_*` usadas no `docker-compose.yml`.

---

## ⚙️ Configuração

### Backend — `backend/.env`

O backend lê as configurações de `backend/.env` via `godotenv`. **A aplicação não inicia sem ele.**

```env
# Ambiente / API
APP_ENV=development
APP_PORT=3000

# Banco de dados (Docker)
POSTGRES_USER=aluguei
POSTGRES_PASSWORD=aluguei
POSTGRES_DB=aluguei
POSTGRES_PORT=5432

# String de conexão usada pelo backend e pelas migrations
DATABASE_URL=postgres://aluguei:aluguei@localhost:5432/aluguei?sslmode=disable

# JWT — use segredos longos e aleatórios (mínimo 32 caracteres)
JWT_ACCESS_SECRET=coloque-um-segredo-grande-aqui-32-chars-min
JWT_REFRESH_SECRET=coloque-outro-segredo-grande-aqui-32-chars-min
```

> ⚠️ **Nunca commite o `.env` com segredos reais.** `backend/.env` já está no `.gitignore`.

### Frontend — proxy e environments

- Em dev, o Angular redireciona `/api` → `http://localhost:3000` via `frontend/proxy.conf.json` (configurado no `angular.json`).
- `environment.apiUrl` deve ser `/api/v1` em dev (passa pelo proxy) e a URL pública do backend em produção (o valor atual é um placeholder).

---

## 📦 Comandos disponíveis

Todos os comandos ficam no `Makefile` da raiz (`make help` lista todos).

### Desenvolvimento

```bash
make dev              # Sobe tudo: Docker + backend + frontend (paralelo)
make dev-backend      # Apenas o backend com hot reload (Air)
make dev-frontend     # Apenas o frontend Angular (ng serve)
make run-backend      # Backend sem hot reload (go run com ldflags de versão)
```

### Build

```bash
make build            # Compila backend + frontend
make build-backend    # Binário do backend (linux/amd64) → backend/bin/rental
make build-frontend   # Build de produção do Angular (Service Worker habilitado)
```

### Testes e qualidade

```bash
make test                     # Testes dos dois projetos
make test-backend             # go test ./... -v
make test-coverage-backend    # Cobertura → backend/coverage.html
make test-frontend            # ng test --watch=false (Vitest)
make lint-backend             # golangci-lint
make fmt-backend              # gofmt + goimports
make tidy                     # go mod tidy
```

### Migrations

```bash
make migrate-up        # Aplica todas as migrations pendentes
make migrate-down      # Reverte a última migration
make migrate-fresh     # Reset completo (down -all + up)
make migrate-version   # Exibe a versão atual do schema
make migrate-create    # Cria uma nova migration (interativo)
```

> As migrations usam a variável `DATABASE_URL`, lida do `backend/.env` pelo Makefile.

### Docker

```bash
make docker-up         # Sobe o PostgreSQL (--env-file backend/.env)
make docker-down       # Derruba os containers
make docker-logs       # Exibe logs
make docker-reset      # Recria containers e volumes
```

### Outros

```bash
make swagger           # Gera/atualiza backend/docs (Swagger)
make release-patch     # Bump de versão no frontend + tag git
make release-minor
make release-major
make help              # Lista completa de comandos
```

---

## 📚 Swagger

Com o backend rodando e os docs gerados (`make swagger`), a documentação interativa está em:

```
http://localhost:3000/swagger/index.html
```

Para regenerar após alterar as anotações em `cmd/api/main.go`:

```bash
make swagger
```

---

## 🔌 API

Base URL: `http://localhost:3000` — prefixo `/api/v1`. Autenticação por **cookies HttpOnly** (`access_token` / `refresh_token`), então os clients devem enviar credenciais (`withCredentials: true`).

| Método | Rota | Auth | Descrição |
|---|---|---|---|
| POST | `/auth/register` | pública | Registra um novo usuário |
| POST | `/auth/login` | pública | Autentica e define os cookies JWT |
| POST | `/auth/refresh` | cookie refresh | Rotaciona os tokens (access + refresh) |
| POST | `/auth/logout` | pública | Limpa os cookies de sessão |
| GET | `/auth/me` | ✅ | Retorna o usuário autenticado |
| GET | `/users/search` | ✅ | Busca usuários (filtros + paginação) |
| GET | `/users/:id` | ✅ | Busca usuário por UUID |
| PUT | `/users/:id` | ✅ | Atualiza dados do usuário |
| DELETE | `/users/:id` | ✅ | Soft delete (204) |

---

## 📁 Estrutura do projeto

```
Aluguei/
├── backend/
│   ├── cmd/api/                     # Entrypoint (main.go) + rotas/DI (server.go)
│   ├── internal/
│   │   ├── delivery/http/
│   │   │   ├── handlers/            # Handlers HTTP (auth, users)
│   │   │   └── middleware/          # Auth (cookie JWT) + ZapLogger
│   │   ├── domain/                  # Camada de domínio (Clean Architecture)
│   │   │   ├── entities/            # User, Address, Property
│   │   │   ├── repositories/        # Interfaces de repositório
│   │   │   └── usecases/            # Interfaces de caso de uso
│   │   ├── infra/
│   │   │   ├── database/            # Conexão GORM/Postgres
│   │   │   ├── repositories/        # Implementações (user, address, property)
│   │   │   └── version/             # Versão injetada via ldflags
│   │   └── usecase/                 # Implementações de casos de uso
│   ├── migrations/                  # SQL (golang-migrate)
│   ├── docs/                        # Swagger gerado (gitignored)
│   └── pkg/                         # hash (Argon2id), jwt, logger, response,
│                                    # pagination, validators (user/address/property)
│
├── frontend/
│   └── src/app/
│       ├── core/                    # auth, guards, interceptors, layout, theme, loading
│       ├── pages/                   # login, register, dashboard, properties,
│       │                            # tenants, contracts, reports
│       ├── components/              # Design system (button, input, modal, table…)
│       └── shared/                  # Pipes, directives, ícones, patterns
│
├── docker-compose.yml               # PostgreSQL 18
└── Makefile                         # Comandos de dev, build, test, migrate, release
```

---

## 🚦 Estado atual & Roadmap

Projeto em desenvolvimento — última revisão: **out/2026**.

### ✅ Funcional de ponta a ponta

- **Autenticação completa**: registro, login, refresh com rotação, logout e `/me` — backend (cookies HttpOnly, Argon2id) + frontend (guards, interceptor com retry/refresh, persistência de sessão).
- **CRUD de usuários**: endpoints com busca/filtros/paginação e soft delete.
- **Base do domínio de imóveis**: entidades `Property`/`Address` com validadores, repositórios e migrations prontos.
- **Design system e layout**: 16 componentes UI, sidebar/topbar/footer, tema claro/escuro, 7 rotas protegidas, skeletons e animações.
- **Qualidade**: ~30 arquivos de teste no backend (entities, usecases, handlers, validators, pkg); Vitest no frontend; lint/format backend via golangci-lint.
- **PWA** ativa em produção + documentação Swagger gerada.

### 🚧 Pendências — backend

- [ ] `Property` e `Address` **não têm endpoints** — repositórios existem mas não há handlers/usecases/rotas.
- [ ] Sem endpoints de **inquilinos, contratos, pagamentos e relatórios** (domínio nem modelado).
- [ ] Swagger com *drift* (ex.: `/users/search` documentado com resposta em array; real é `pagination.Result`).
- [ ] `backend/.env.example` vazio — sem template para setup inicial.
- [ ] CORS com origins de LAN hardcoded em `cmd/api/server.go`.
- [ ] Sem graceful shutdown nem revogação server-side de tokens.
- [ ] `make setup` aponta para `scripts/setup.sh` inexistente (alvo quebrado).

### 🚧 Pendências — frontend

- [ ] **Dashboard, Propriedades, Inquilinos, Contratos e Relatórios usam dados mock** — `save()` ainda é `alert()`/`console.log()`; nenhum service de API além do auth.
- [ ] `AuthService` é o único client HTTP — criar services/genérico para os demais recursos.
- [ ] Topbar com usuário **hardcoded** ("João Pereira") e busca decorativa.
- [ ] Manifest PWA ainda diz "MeuApp"; ícones/favicon referenciados não existem em `public/`.
- [ ] Rota `/forbidden` é referenciada no interceptor mas não existe (cai no wildcard → login).
- [ ] `make lint-frontend` quebra: não há target de lint nem ESLint configurado.
- [ ] Sem testes de páginas/componentes do design system.

### 💡 Próximos passos sugeridos

1. Expor `GET/POST/PUT/DELETE /api/v1/properties` (repositório já existe) e consumir no frontend.
2. Modelar e implementar **inquilinos → contratos → pagamentos** (o coração do produto).
3. Dashboard e Relatórios reais a partir dos pagamentos.
4. Preencher `.env.example`, corrigir manifest/PWA e o target de lint.

---

## 📄 Licença

Distribuído sob a licença **MIT**. Consulte o arquivo [LICENSE](./LICENSE) para mais detalhes.

---

<div align="center">
  Feito com ❤️ por <a href="https://github.com/Turgho">Turgho</a>
</div>
