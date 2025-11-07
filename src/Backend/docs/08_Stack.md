# 🛠️ Stack Tecnológica - Sistema Aluguei

## Visão Geral

O sistema Aluguei utiliza tecnologias modernas e confiáveis, priorizando performance, escalabilidade e manutenibilidade.

---

## Backend (MVP 1 - Implementado ✅)

### Linguagem Principal
- **Go 1.25+**
  - Performance superior
  - Concorrência nativa (goroutines)
  - Compilação estática
  - Garbage collector eficiente
  - Ecossistema maduro para APIs

### Framework Web
- **Gin 1.9+**
  - Router HTTP rápido e minimalista
  - Middleware ecosystem
  - JSON binding automático
  - Validação integrada
  - Documentação excelente

### Banco de Dados
- **PostgreSQL 18**
  - ACID compliance
  - Suporte a JSON/JSONB
  - Extensibilidade (UUID, full-text search)
  - Performance para aplicações transacionais
  - Backup e recovery robustos

### ORM
- **GORM 1.25+**
  - Auto-migration
  - Associations e preloading
  - Hooks e callbacks
  - Query builder intuitivo
  - Suporte a transações

### Autenticação
- **JWT (JSON Web Tokens)**
  - Stateless authentication
  - Payload customizável
  - Expiração automática
  - Biblioteca: `golang-jwt/jwt/v5`

### Criptografia
- **bcrypt**
  - Hash de senhas seguro
  - Salt automático
  - Resistente a ataques de força bruta
  - Biblioteca: `golang.org/x/crypto/bcrypt`

### Validação
- **Validator v10**
  - Validação struct-based
  - Tags declarativas
  - Mensagens customizáveis
  - Biblioteca: `go-playground/validator/v10`

### Documentação API
- **Swagger/OpenAPI 3.0**
  - Especificação completa da API
  - Interface interativa
  - Geração automática de clientes
  - Arquivo: `docs/swagger.yaml`

---

## Infraestrutura (MVP 1 - Implementado ✅)

### Containerização
- **Docker 24+**
  - Ambiente consistente
  - Isolamento de dependências
  - Deploy simplificado
  - Multi-stage builds

### Orquestração Local
- **Docker Compose**
  - Ambiente de desenvolvimento
  - Serviços integrados
  - Networking automático
  - Volumes persistentes

### Banco de Dados (Desenvolvimento)
```yaml
# docker-compose.yml
postgres:
  image: postgres:18-alpine
  environment:
    POSTGRES_DB: aluguei_db
    POSTGRES_USER: aluguei_user
    POSTGRES_PASSWORD: aluguei_password
  ports:
    - "5433:5432"
  volumes:
    - postgres_data:/var/lib/postgresql/data
```

### Cache (Preparado)
- **Redis 8+**
  - Cache de sessões
  - Rate limiting
  - Pub/Sub para notificações
  - Estruturas de dados avançadas

### Administração DB
- **Adminer**
  - Interface web para PostgreSQL
  - Queries SQL diretas
  - Visualização de esquemas
  - Export/Import de dados

---

## Desenvolvimento (MVP 1 - Implementado ✅)

### Gerenciamento de Dependências
- **Go Modules**
  - Versionamento semântico
  - Vendor directory opcional
  - Proxy de módulos
  - Arquivo: `go.mod`

### Build e Automação
- **Makefile**
  - Comandos padronizados
  - Build, test, lint
  - Docker operations
  - Database seeding

### Testes
- **Testing nativo do Go**
  - Testes unitários
  - Testes de integração
  - Benchmarks
  - Coverage reports

### Ferramentas de Teste
```go
// Bibliotecas utilizadas
- testify/assert    // Assertions
- testify/mock      // Mocking
- testify/suite     // Test suites
- gorm.io/driver/sqlite // In-memory DB para testes
```

### Linting e Formatação
- **gofmt** - Formatação padrão
- **go vet** - Análise estática
- **golangci-lint** - Linter abrangente

---

## Monitoramento e Observabilidade (MVP 1 - Básico ✅)

### Health Checks
- **Endpoint /health**
  - Status da aplicação
  - Conectividade do banco
  - Métricas básicas
  - Formato JSON padronizado

### Logging
- **Log nativo do Go**
  - Structured logging
  - Diferentes níveis (DEBUG, INFO, WARN, ERROR)
  - Contexto de requisições
  - Rotação de logs

### Métricas (Preparado)
- **Prometheus** (futuro)
  - Métricas de aplicação
  - Métricas de sistema
  - Alerting rules
  - Grafana dashboards

---

## Segurança (MVP 1 - Implementado ✅)

### HTTPS/TLS
- **Certificados SSL**
  - Let's Encrypt (produção)
  - Self-signed (desenvolvimento)
  - Redirecionamento HTTP → HTTPS

### CORS
- **Gin CORS Middleware**
  - Origins permitidas
  - Headers customizados
  - Credentials support
  - Preflight handling

### Rate Limiting (Preparado)
- **Redis-based**
  - Limite por IP
  - Limite por usuário
  - Sliding window
  - Diferentes endpoints

### Validação de Input
- **Sanitização**
  - SQL injection prevention
  - XSS protection
  - Input validation
  - Output encoding

---

## Frontend (MVP 2 - Planejado 🔄)

### Framework
- **React 18+**
  - Component-based architecture
  - Virtual DOM
  - Hooks ecosystem
  - Server-side rendering (Next.js)

### Linguagem
- **TypeScript 5+**
  - Type safety
  - Better IDE support
  - Refactoring tools
  - Interface definitions

### Styling
- **Tailwind CSS 3+**
  - Utility-first CSS
  - Responsive design
  - Dark mode support
  - Component libraries

### Estado Global
- **Zustand** ou **Redux Toolkit**
  - State management
  - Middleware support
  - DevTools integration
  - Persistence

### Formulários
- **React Hook Form + Zod**
  - Performance otimizada
  - Validação schema-based
  - Error handling
  - TypeScript integration

### HTTP Client
- **Axios** ou **React Query**
  - Request/response interceptors
  - Caching automático
  - Error handling
  - Loading states

---

## Mobile (MVP 3 - Futuro 📋)

### Framework
- **React Native** ou **Flutter**
  - Cross-platform development
  - Native performance
  - Shared codebase
  - Platform-specific features

### Push Notifications
- **Firebase Cloud Messaging**
  - Cross-platform notifications
  - Targeting e segmentação
  - Analytics integrado
  - A/B testing

---

## DevOps e Deploy (MVP 2/3 - Planejado)

### CI/CD
- **GitHub Actions**
  - Automated testing
  - Build e deploy
  - Security scanning
  - Dependency updates

### Cloud Provider
- **AWS** ou **Google Cloud**
  - Compute instances
  - Managed databases
  - Load balancers
  - CDN e storage

### Containerização (Produção)
- **Kubernetes** ou **Docker Swarm**
  - Orchestration
  - Auto-scaling
  - Service discovery
  - Rolling updates

### Monitoramento (Produção)
- **Prometheus + Grafana**
  - Métricas de aplicação
  - Dashboards customizados
  - Alerting rules
  - SLA monitoring

### Logging (Produção)
- **ELK Stack** (Elasticsearch, Logstash, Kibana)
  - Centralized logging
  - Log analysis
  - Search e filtering
  - Visualizations

---

## Integrações (MVP 2/3 - Planejado)

### Pagamentos
- **PIX**
  - API do Banco Central
  - QR Code dinâmico
  - Webhook notifications
  - Conciliação automática

### Email
- **SendGrid** ou **Amazon SES**
  - Transactional emails
  - Templates customizáveis
  - Delivery tracking
  - Bounce handling

### SMS
- **Twilio** ou **Amazon SNS**
  - Notificações SMS
  - Two-factor authentication
  - Delivery reports
  - International support

### Assinatura Digital
- **DocuSign** ou **ClickSign**
  - Contratos digitais
  - Validade jurídica
  - Workflow de aprovação
  - Audit trail

### Mapas
- **Google Maps API**
  - Geolocalização
  - Endereços automáticos
  - Visualização de propriedades
  - Cálculo de distâncias

---

## Dependências Principais (go.mod)

```go
module github.com/turgho/aluguei

go 1.25

require (
    github.com/gin-gonic/gin v1.9.1
    github.com/golang-jwt/jwt/v5 v5.0.0
    github.com/google/uuid v1.3.0
    golang.org/x/crypto v0.14.0
    gorm.io/driver/postgres v1.5.2
    gorm.io/gorm v1.25.4
)

require (
    github.com/go-playground/validator/v10 v10.15.5
    github.com/stretchr/testify v1.8.4
    github.com/gin-contrib/cors v1.4.0
)
```

---

## Estrutura de Configuração

### Variáveis de Ambiente
```bash
# .env
# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=aluguei
DB_SSLMODE=disable

# Server
SERVER_PORT=8080
SERVER_HOST=localhost

# JWT
JWT_SECRET=your-secret-key
JWT_EXPIRATION=24h

# Redis (futuro)
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=

# Email (futuro)
SENDGRID_API_KEY=
EMAIL_FROM=noreply@aluguei.com
```

### Configuração da Aplicação
```go
// internal/config/config.go
type Config struct {
    Database DatabaseConfig
    Server   ServerConfig
    JWT      JWTConfig
    Redis    RedisConfig
}
```

---

## Métricas de Performance

### Targets MVP 1 ✅
- **Response Time**: < 100ms (95th percentile)
- **Throughput**: > 1000 req/s
- **Memory Usage**: < 100MB
- **CPU Usage**: < 50%
- **Database Connections**: Pool de 20 conexões

### Targets MVP 2 🔄
- **Frontend Load Time**: < 2s
- **API Response Time**: < 50ms
- **Database Query Time**: < 10ms
- **Cache Hit Rate**: > 80%

### Targets MVP 3 📋
- **Mobile App Size**: < 50MB
- **Offline Capability**: 24h
- **Push Notification Delivery**: > 95%
- **Cross-platform Consistency**: 100%

Esta stack tecnológica garante que o sistema Aluguei seja robusto, escalável e mantenha alta performance conforme cresce em funcionalidades e usuários.