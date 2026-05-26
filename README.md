<div align="center">

# 🏠 Aluguei

**Gerencie aluguéis, inquilinos, contratos e imóveis em um só lugar.**

[![Status](https://img.shields.io/badge/status-em%20desenvolvimento-orange?style=flat-square)](#)
[![Go](https://img.shields.io/badge/Go-1.26+-00ADD8?style=flat-square&logo=go&logoColor=white)](https://go.dev)
[![Angular](https://img.shields.io/badge/Angular-21+-dd0031?style=flat-square&logo=angular&logoColor=white)](https://angular.dev)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-Docker-336791?style=flat-square&logo=postgresql&logoColor=white)](https://www.postgresql.org)
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
- [Swagger](#-swagger)
- [Estrutura do projeto](#-estrutura-do-projeto)
- [Licença](#-licença)

---

## 📌 Sobre

O **Aluguei** é um **PWA** (Progressive Web App) voltado para proprietários que precisam organizar sua carteira de imóveis. Com ele é possível controlar aluguéis, cadastrar inquilinos, gerar contratos e visualizar relatórios — tudo em uma interface moderna e acessível do celular ou do computador.

---

## ✨ Features

| Módulo | Descrição |
|---|---|
| 🏠 **Propriedades** | Cadastro e gestão de imóveis |
| 👤 **Inquilinos** | Cadastro e gestão de locatários |
| 📄 **Contratos** | Criação e acompanhamento de contratos de locação |
| 📊 **Dashboard** | Visão geral com gráficos e indicadores |
| 📈 **Relatórios** | Geração de relatórios detalhados |
| 🔐 **Autenticação** | Login seguro via JWT em cookies HttpOnly (login / refresh / logout) |

---

## 🛠 Stack

**Frontend**
- Angular (standalone components) + TailwindCSS
- PWA com Service Worker
- Gráficos com Chart.js via `ng2-charts`

**Backend**
- Go + [Gin](https://github.com/gin-gonic/gin) (HTTP framework)
- [GORM](https://gorm.io) com PostgreSQL
- JWT para autenticação
- Zap para logging estruturado

**Infraestrutura**
- PostgreSQL via Docker
- Documentação automática com Swagger (`/swagger/*`)

---

## 🚀 Início Rápido

### Pré-requisitos

Certifique-se de ter instalado:

| Ferramenta | Uso |
|---|---|
| **Go 1.26+** | Backend (`backend/go.mod`) |
| **Node.js + npm** | Frontend (`frontend/package.json`) |
| **Angular CLI** | Servidor de desenvolvimento |
| **Docker + Docker Compose** | Banco de dados PostgreSQL |
| **[migrate](https://github.com/golang-migrate/migrate)** | Aplicar migrations SQL |
| **[air](https://github.com/air-verse/air)** | Hot reload do backend |
| **[swag](https://github.com/swaggo/swag)** | Geração do Swagger |

### 1. Clone o repositório

```bash
git clone https://github.com/Turgho/Aluguei.git
cd Aluguei
```

### 2. Configure as variáveis de ambiente

Crie o arquivo `backend/.env` (veja a seção [Configuração](#-configuração)).

### 3. Suba o ambiente completo

```bash
make dev
```

Isso irá automaticamente:
- Subir o PostgreSQL via Docker
- Iniciar o backend com hot reload (Air) em `localhost:3000`
- Iniciar o frontend Angular em `localhost:4200`

> Acesse a aplicação em **http://localhost:4200** e a documentação da API em **http://localhost:3000/swagger/index.html**

---

## ⚙️ Configuração

O backend lê as configurações de `backend/.env`. Esse mesmo arquivo é usado pelo Docker Compose. **A aplicação não inicia sem ele.**

Crie `backend/.env` com o seguinte conteúdo:

```env
# Ambiente / API
APP_ENV=development
APP_PORT=3000

# Banco de dados (Docker)
POSTGRES_USER=aluguei
POSTGRES_PASSWORD=aluguei
POSTGRES_DB=aluguei
POSTGRES_PORT=5432

# String de conexão usada pelo backend
DATABASE_URL=postgres://aluguei:aluguei@localhost:5432/aluguei?sslmode=disable

# JWT — use segredos longos e aleatórios (mínimo 32 caracteres)
JWT_ACCESS_SECRET=coloque-um-segredo-grande-aqui-32-chars-min
JWT_REFRESH_SECRET=coloque-outro-segredo-grande-aqui-32-chars-min
```

> ⚠️ **Nunca commite o `.env` com segredos reais.** Adicione `backend/.env` ao seu `.gitignore`.

---

## 📦 Comandos disponíveis

### Desenvolvimento

```bash
make dev              # Sobe tudo: Docker + backend + frontend
make docker-up        # Apenas o PostgreSQL via Docker
make dev-backend      # Apenas o backend com hot reload (Air)
make dev-frontend     # Apenas o frontend Angular
```

### Build

```bash
make build            # Compila o binário do backend
```

### Utilitários

```bash
make swagger          # Gera/atualiza os arquivos do Swagger
```

### Proxy do frontend

O Angular está configurado para redirecionar chamadas `/api` para o backend em `localhost:3000`, via `frontend/proxy.conf.json`.

---

## 🗄 Migrations

Os arquivos SQL de migration ficam em `backend/migrations/`.

```bash
make migrate-up        # Aplica todas as migrations pendentes
make migrate-down      # Reverte a última migration
make migrate-fresh     # Reset completo (down -all + up)
make migrate-version   # Exibe a versão atual do schema
```

> As migrations usam a variável `DATABASE_URL` definida em `backend/.env`.

---

## 📚 Swagger

Com o backend rodando, a documentação interativa da API está disponível em:

```
http://localhost:3000/swagger/index.html
```

Para regenerar os arquivos após alterações nas anotações:

```bash
make swagger
```

---

## 📁 Estrutura do projeto

```
Aluguei/
├── backend/
│   ├── cmd/api/                    # Entrypoint do servidor (Gin)
│   ├── internal/
│   │   ├── delivery/http/          # Handlers e middlewares HTTP
│   │   ├── infra/                  # Banco, repositórios, versão etc.
│   │   └── usecase/                # Regras de negócio (ex: auth)
│   ├── migrations/                 # Arquivos SQL de migration
│   └── pkg/                        # Pacotes internos (jwt, logger, response…)
│
├── frontend/
│   ├── src/app/
│   │   ├── core/                   # Auth, interceptors, guards, layout
│   │   ├── pages/                  # Dashboard, propriedades, inquilinos, contratos, relatórios
│   │   └── components/             # Componentes de UI (button, input, modal…)
│   ├── ngsw-config.json            # Configuração do Service Worker (PWA)
│   └── public/manifest.webmanifest # Web App Manifest
│
├── docker-compose.yml              # PostgreSQL
└── Makefile                        # Comandos de dev, build e migrate
```

---

## 📄 Licença

Distribuído sob a licença **MIT**. Consulte o arquivo [LICENSE](./LICENSE) para mais detalhes.

---

<div align="center">
  Feito com ❤️ por <a href="https://github.com/Turgho">Turgho</a>
</div>