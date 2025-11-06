# 🏠 Aluguei - Sistema Inteligente de Gestão de Aluguel

[![Status](https://img.shields.io/badge/status-em%20desenvolvimento-orange)](https://github.com/Turgho/Aluguei)
[![Backend](https://img.shields.io/badge/Linguagem-Go-blue)](https://go.dev/doc/install)
[![Frontend](https://img.shields.io/badge/Frontend-React%2FTypeScript-61DAFB)](https://react.dev/learn/installation)
[![Último Commit](https://img.shields.io/github/last-commit/Turgho/aluguei)](https://github.com/Turgho/Aluguei/commits/main)
[![License: MIT](https://img.shields.io/badge/License-MIT-green)](./LICENSE)

## 🚀 O Que é o Aluguei?

Aluguei é um sistema completo para gestão de aluguéis que conecta proprietários e inquilinos, automatizando processos manuais, reduzindo inadimplência e simplificando a comunicação.
> Status atual: repositório público — desenvolvimento ativo.
> Escopo visível aqui: somente o MVP 1 (Cadastro, Login, Gestão de Imóveis e Contratos). A documentação completa está em docs/.

### 🎯 MVP 1 — Recursos disponíveis neste repositório

- Implementação mínima viável para demonstrar o fluxo básico de gestão:

    - ✅ Cadastro de proprietário (nome, email, senha, telefone)
    - ✅ Cadastro de inquilino (nome, email, senha, documento)
    - ✅ Login com autenticação JWT
    - ✅ Gestão de imóveis (cadastro, edição, status disponível/alugado)
    - ✅ Contratos digitais (criação, visualização, status)
    - ✅ Painéis distintos (proprietário × inquilino)
    - ✅ Estrutura do banco (PostgreSQL) e migrations
    - ✅ Documentação completa em docs/

    > Observação: funcionalidades como pagamentos PIX, sistema de manutenção e app mobile estão documentadas e planejadas (veja docs/), mas não estão ativas no MVP1 deste repositório público.

📂 Documentação (pasta docs/)

A documentação completa foi organizada em arquivos Markdown dentro de docs/:

- `docs/01_Objetivo.md` — Visão geral
- `docs/02_MVPs.md` — MVPs do projeto
- `docs/03_Entidades.md` — Entidades do projeto
- `docs/04_UserStories.md` — Backlog e user stories
- `docs/05_BancoDeDados.md` — Modelagem SQL Server (DDL)
- `docs/06_Arquitetura.md` — Arquitetura do projeto
- `docs/07_Fluxos.md` — Fluxos em Mermaid
- `docs/08_Stack.md` — Stack usada no projeto
- `docs/09_MetricasDeAceitação.md` — Estimativa para sucesso do projeto
- `docs/10_Segurança.md` — Métodos para segurança do usuário

## 🚀 Começando
- Pré-requisitos
  - Go 1.21+
  - Node.js 18+
  - PostgreSQL 14+
  - Docker (opcional)

### Instalação Local
```bash

# Clone o repositório
git clone https://github.com/seu-usuario/aluguei.git
cd aluguei

# Configure as variáveis de ambiente
cp .env.example .env
# Edite .env com suas configurações

# Instale dependências do backend
cd backend
go mod download

# Instale dependências do frontend
cd ../frontend
npm install

# Execute a aplicação
docker-compose up -d
# ou
make dev
```

### Acesso
```text
    Frontend: http://localhost:3000
    API Backend: http://localhost:8080
    Adminer (DB): http://localhost:8081
```

## 📊 Estrutura do Projeto
```text

aluguei/
├── 📁 backend/                 # API em Go
│   ├── cmd/api/               # Application entrypoint
│   ├── internal/              # Código interno do projeto
│   │   ├── handlers/          # HTTP handlers
│   │   ├── models/            # Entidades de domínio
│   │   ├── services/          # Lógica de negócio
│   │   └── database/          # Configuração do DB
│   └── pkg/                   # Código compartilhável
├── 📁 frontend/               # Aplicação React/Next.js
│   ├── components/            # Componentes reutilizáveis
│   ├── pages/                 # Rotas e páginas
│   ├── stores/                # Gerenciamento de estado
│   └── styles/                # Estilos globais
├── 📁 docs/                   # Documentação completa
├── 📁 deployments/            # Configurações de deploy
└── 📁 scripts/                # Scripts utilitários
```

## 🤝 Como Contribuir

- Se quiser contribuir com issues, PRs ou feedback:
  - Abra uma issue descrevendo a proposta ou bug
  - Faça fork do repositório
  - Crie uma branch (feature/nova-funcionalidade ou fix/correcao)
  - Siga os padrões de código (veja docs/11_Contribuindo.md)
  - Abra um Pull Request contra main

### Padrões de Código

```bash
    # Backend (Go)
    go fmt ./...
    go vet ./...
    golangci-lint run

    # Frontend
    npm run lint
    npm run type-check
    npm run test
```

## 🔒 Segurança e LGPD

- Nenhuma credencial sensível é commitada no repositório
- Dados pessoais são tratados conforme LGPD
- Autenticação utiliza JWT com tempo de expiração
- Senhas são hashadas com bcrypt
- Use variáveis de ambiente para configurações sensíveis

> Se encontrar vulnerabilidade, reporte seguindo as diretrizes em docs/10_Seguranca.md.

## 📞 Contato

- Autor/Maintainer: Seu Nome
- Email: contato@aluguei.app
- Site: https://aluguei.app (futuro)

Para sugestões ou dúvidas técnicas, abra uma issue no repositório.

## 📄 Licença

- Este projeto está licenciado sob a Licença MIT - veja o arquivo LICENSE para detalhes.
- Aluguei — Transformando a gestão de aluguel, um imóvel de cada vez. 🏠✨
- Obrigado por visitar o repositório do Aluguei!
