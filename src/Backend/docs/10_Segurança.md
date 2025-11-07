# 🔒 Segurança - Sistema Aluguei

## Visão Geral

A segurança é uma prioridade fundamental no sistema Aluguei, considerando que lidamos com dados pessoais sensíveis (CPF, endereços, informações financeiras) e devemos estar em conformidade com a LGPD (Lei Geral de Proteção de Dados).

---

## 🛡️ Princípios de Segurança

### Security by Design
- **Segurança desde o início**: Implementada em todas as camadas
- **Princípio do menor privilégio**: Acesso mínimo necessário
- **Defesa em profundidade**: Múltiplas camadas de proteção
- **Fail-safe defaults**: Configurações seguras por padrão

### Compliance
- **LGPD**: Lei Geral de Proteção de Dados Pessoais
- **Marco Civil da Internet**: Regulamentação brasileira
- **ISO 27001**: Padrões de segurança da informação
- **OWASP Top 10**: Vulnerabilidades web mais críticas

---

## 🔐 Autenticação e Autorização

### Sistema de Autenticação (Implementado ✅)

#### JWT (JSON Web Tokens)
```go
// Estrutura do token JWT
type Claims struct {
    UserID   uuid.UUID `json:"user_id"`
    Email    string    `json:"email"`
    UserType string    `json:"user_type"` // "owner", "tenant"
    jwt.RegisteredClaims
}

// Configuração segura
const (
    TokenExpiration = 24 * time.Hour
    RefreshExpiration = 7 * 24 * time.Hour
    MinSecretLength = 32
)
```

#### Características de Segurança
- **Algoritmo**: HS256 (HMAC SHA-256)
- **Expiração**: 24 horas (configurável)
- **Refresh Token**: 7 dias para renovação
- **Secret Key**: Mínimo 32 caracteres, armazenado em variável de ambiente
- **Payload mínimo**: Apenas dados essenciais

### Middleware de Autenticação
```go
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // 1. Extrair token do header Authorization
        // 2. Validar formato Bearer <token>
        // 3. Verificar assinatura e expiração
        // 4. Extrair claims e adicionar ao contexto
        // 5. Continuar ou retornar 401 Unauthorized
    }
}
```

### Controle de Acesso (Implementado ✅)

#### Níveis de Autorização
1. **Público**: Endpoints de cadastro e login
2. **Autenticado**: Requer token JWT válido
3. **Proprietário**: Acesso apenas aos próprios recursos
4. **Admin**: Acesso total (futuro)

#### Validação de Propriedade
```go
// Exemplo: Proprietário só acessa suas propriedades
func (h *PropertyHandler) GetProperty(c *gin.Context) {
    userID := getUserIDFromContext(c)
    propertyID := c.Param("id")
    
    property, err := h.propertyUseCase.GetByID(c, propertyID)
    if err != nil {
        c.JSON(404, gin.H{"error": "Property not found"})
        return
    }
    
    // Verificar se a propriedade pertence ao usuário
    if property.OwnerID != userID {
        c.JSON(403, gin.H{"error": "Access denied"})
        return
    }
    
    c.JSON(200, property)
}
```

---

## 🔒 Criptografia e Hashing

### Hash de Senhas (Implementado ✅)

#### bcrypt
```go
import "golang.org/x/crypto/bcrypt"

const (
    MinPasswordLength = 6
    MaxPasswordLength = 128
    BcryptCost = 12 // Custo computacional
)

func HashPassword(password string) (string, error) {
    if len(password) < MinPasswordLength {
        return "", errors.New("password too short")
    }
    
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), BcryptCost)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
```

#### Características
- **Algoritmo**: bcrypt com salt automático
- **Custo**: 12 (ajustável conforme hardware)
- **Resistência**: Força bruta e rainbow tables
- **Performance**: ~100ms por hash (aceitável para login)

### Dados Sensíveis (Implementado ✅)

#### Criptografia de CPF
```go
// CPF armazenado com hash para busca e valor criptografado
type Owner struct {
    ID       uuid.UUID `json:"id"`
    Name     string    `json:"name"`
    Email    string    `json:"email"`
    CPF      string    `json:"cpf"`      // Formato: XXX.XXX.XXX-XX (mascarado na API)
    CPFHash  string    `json:"-"`        // Hash para busca única
    // ... outros campos
}
```

#### Mascaramento de Dados
```go
func MaskCPF(cpf string) string {
    if len(cpf) != 11 {
        return "***.***.***-**"
    }
    return cpf[:3] + ".***.***-" + cpf[9:]
}

func MaskEmail(email string) string {
    parts := strings.Split(email, "@")
    if len(parts) != 2 {
        return "***@***.***"
    }
    return parts[0][:2] + "***@" + parts[1]
}
```

---

## 🛡️ Proteção contra Vulnerabilidades

### OWASP Top 10 - Mitigações Implementadas

#### A01 - Broken Access Control ✅
- **Implementado**: Middleware de autenticação obrigatório
- **Validação**: Propriedade de recursos por usuário
- **Princípio**: Menor privilégio e fail-safe defaults

#### A02 - Cryptographic Failures ✅
- **Implementado**: bcrypt para senhas, JWT para sessões
- **TLS**: HTTPS obrigatório em produção
- **Secrets**: Variáveis de ambiente, nunca hardcoded

#### A03 - Injection ✅
- **Implementado**: GORM ORM previne SQL injection
- **Validação**: Input validation em todos os endpoints
- **Sanitização**: Escape de caracteres especiais

#### A04 - Insecure Design ✅
- **Implementado**: Security by design desde o início
- **Threat Modeling**: Análise de ameaças por feature
- **Secure Defaults**: Configurações seguras por padrão

#### A05 - Security Misconfiguration ✅
- **Implementado**: Configurações via environment variables
- **Headers**: Security headers obrigatórios
- **Error Handling**: Mensagens de erro genéricas

#### A06 - Vulnerable Components 🔄
- **Planejado**: Dependabot para atualizações automáticas
- **Scanning**: Análise de vulnerabilidades em dependências
- **Monitoring**: Alertas para CVEs críticas

#### A07 - Authentication Failures ✅
- **Implementado**: Rate limiting (preparado)
- **Session Management**: JWT com expiração
- **Password Policy**: Validação de força da senha

#### A08 - Software Integrity Failures 🔄
- **Planejado**: Assinatura de releases
- **CI/CD**: Pipeline seguro com verificações
- **Dependencies**: Verificação de integridade

#### A09 - Logging Failures ✅
- **Implementado**: Logs estruturados sem dados sensíveis
- **Monitoring**: Alertas para eventos suspeitos
- **Retention**: Política de retenção de logs

#### A10 - Server-Side Request Forgery 🔄
- **Planejado**: Validação de URLs externas
- **Whitelist**: Apenas domínios permitidos
- **Network**: Segmentação de rede

---

## 🔍 Validação e Sanitização

### Validação de Input (Implementado ✅)

#### Estrutura de Validação
```go
type CreateOwnerRequest struct {
    Name     string `json:"name" validate:"required,min=2,max=100"`
    Email    string `json:"email" validate:"required,email"`
    Phone    string `json:"phone" validate:"required,phone"`
    CPF      string `json:"cpf" validate:"required,cpf"`
    Password string `json:"password" validate:"required,min=6,max=128"`
}

// Validador customizado para CPF
func ValidateCPF(fl validator.FieldLevel) bool {
    cpf := fl.Field().String()
    return isValidCPF(cpf)
}
```

#### Sanitização
```go
func SanitizeInput(input string) string {
    // Remove caracteres perigosos
    input = html.EscapeString(input)
    input = strings.TrimSpace(input)
    
    // Remove caracteres de controle
    re := regexp.MustCompile(`[\x00-\x1f\x7f]`)
    return re.ReplaceAllString(input, "")
}
```

### Rate Limiting (Preparado 🔄)

#### Configuração por Endpoint
```go
// Limites por endpoint
var rateLimits = map[string]RateLimit{
    "POST /auth/login":    {Requests: 5, Window: time.Minute},
    "POST /owners":        {Requests: 3, Window: time.Hour},
    "GET /properties":     {Requests: 100, Window: time.Minute},
    "default":             {Requests: 60, Window: time.Minute},
}
```

#### Implementação com Redis
```go
func RateLimitMiddleware(redis *redis.Client) gin.HandlerFunc {
    return func(c *gin.Context) {
        key := fmt.Sprintf("rate_limit:%s:%s", 
            c.ClientIP(), c.Request.URL.Path)
        
        // Implementar sliding window com Redis
        // Bloquear se exceder limite
    }
}
```

---

## 🔐 LGPD e Privacidade

### Conformidade LGPD (Implementado ✅)

#### Princípios Implementados
1. **Finalidade**: Dados coletados apenas para gestão de aluguéis
2. **Adequação**: Tratamento compatível com finalidades
3. **Necessidade**: Apenas dados essenciais
4. **Livre Acesso**: Usuário pode consultar seus dados
5. **Qualidade**: Dados mantidos atualizados
6. **Transparência**: Informações claras sobre tratamento
7. **Segurança**: Medidas técnicas adequadas
8. **Prevenção**: Evitar danos por tratamento inadequado
9. **Não Discriminação**: Sem finalidades discriminatórias
10. **Responsabilização**: Demonstrar conformidade

#### Direitos do Titular
```go
// Endpoints para exercer direitos LGPD
type LGPDHandler struct {
    userUseCase *usecases.UserUseCase
}

// Direito de acesso (Art. 18, II)
func (h *LGPDHandler) GetMyData(c *gin.Context) {
    userID := getUserIDFromContext(c)
    data, err := h.userUseCase.ExportUserData(c, userID)
    // Retorna todos os dados do usuário
}

// Direito de correção (Art. 18, III)
func (h *LGPDHandler) UpdateMyData(c *gin.Context) {
    // Permite atualização de dados pessoais
}

// Direito de eliminação (Art. 18, VI)
func (h *LGPDHandler) DeleteMyData(c *gin.Context) {
    // Soft delete mantendo obrigações legais
}
```

### Tratamento de Dados Pessoais

#### Categorização de Dados
```go
// Dados pessoais básicos
type PersonalData struct {
    Name      string `json:"name"`       // Identificação
    Email     string `json:"email"`      // Contato
    Phone     string `json:"phone"`      // Contato
    BirthDate *time.Time `json:"birth_date"` // Opcional
}

// Dados pessoais sensíveis
type SensitiveData struct {
    CPF string `json:"cpf"` // Documento de identificação
}

// Dados financeiros
type FinancialData struct {
    RentAmount   decimal.Decimal `json:"rent_amount"`
    PaymentData  []Payment       `json:"payments"`
}
```

#### Retenção de Dados
```go
const (
    // Períodos de retenção conforme legislação
    PersonalDataRetention = 5 * 365 * 24 * time.Hour  // 5 anos
    FinancialDataRetention = 5 * 365 * 24 * time.Hour // 5 anos (Receita Federal)
    LogRetention = 6 * 30 * 24 * time.Hour            // 6 meses
)

// Processo de limpeza automática
func CleanupExpiredData(db *gorm.DB) error {
    cutoffDate := time.Now().Add(-PersonalDataRetention)
    
    // Anonimizar dados expirados
    return db.Model(&Owner{}).
        Where("deleted_at < ?", cutoffDate).
        Updates(map[string]interface{}{
            "name": "ANONIMIZADO",
            "email": "anonimo@anonimo.com",
            "cpf": "00000000000",
        }).Error
}
```

---

## 🚨 Monitoramento e Detecção

### Logs de Segurança (Implementado ✅)

#### Eventos Monitorados
```go
type SecurityEvent struct {
    ID        uuid.UUID `json:"id"`
    Type      string    `json:"type"`      // LOGIN, FAILED_LOGIN, DATA_ACCESS
    UserID    uuid.UUID `json:"user_id"`
    IP        string    `json:"ip"`
    UserAgent string    `json:"user_agent"`
    Timestamp time.Time `json:"timestamp"`
    Details   string    `json:"details"`
}

// Tipos de eventos
const (
    EventLogin        = "LOGIN"
    EventFailedLogin  = "FAILED_LOGIN"
    EventDataAccess   = "DATA_ACCESS"
    EventDataExport   = "DATA_EXPORT"
    EventPasswordChange = "PASSWORD_CHANGE"
    EventSuspiciousActivity = "SUSPICIOUS_ACTIVITY"
)
```

#### Logging Seguro
```go
func LogSecurityEvent(eventType string, userID uuid.UUID, ip string, details string) {
    event := SecurityEvent{
        ID:        uuid.New(),
        Type:      eventType,
        UserID:    userID,
        IP:        maskIP(ip), // Mascarar IP para LGPD
        Timestamp: time.Now(),
        Details:   sanitizeLogDetails(details),
    }
    
    // Não logar dados sensíveis
    logger.Info("Security event", 
        "type", event.Type,
        "user_id", event.UserID,
        "ip", event.IP,
    )
}
```

### Detecção de Anomalias (Planejado 🔄)

#### Padrões Suspeitos
- **Múltiplos logins falhados**: > 5 tentativas em 10 minutos
- **Login de IP diferente**: Localização geográfica incomum
- **Acesso em horário atípico**: Fora do padrão do usuário
- **Volume de requests**: Muito acima da média
- **Padrões de bot**: User-agent suspeito, timing regular

#### Resposta Automática
```go
func DetectSuspiciousActivity(userID uuid.UUID, ip string) {
    // Verificar padrões suspeitos
    if isMultipleFailedLogins(userID, ip) {
        // Bloquear temporariamente
        blockUser(userID, 15*time.Minute)
        
        // Alertar administradores
        sendSecurityAlert("Multiple failed logins", userID, ip)
    }
    
    if isUnusualLocation(userID, ip) {
        // Requerer verificação adicional
        requireTwoFactorAuth(userID)
    }
}
```

---

## 🔧 Configurações de Segurança

### Headers de Segurança (Implementado ✅)

```go
func SecurityHeadersMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Prevenir clickjacking
        c.Header("X-Frame-Options", "DENY")
        
        // Prevenir MIME sniffing
        c.Header("X-Content-Type-Options", "nosniff")
        
        // XSS Protection
        c.Header("X-XSS-Protection", "1; mode=block")
        
        // HSTS (HTTPS obrigatório)
        c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
        
        // Content Security Policy
        c.Header("Content-Security-Policy", "default-src 'self'")
        
        // Referrer Policy
        c.Header("Referrer-Policy", "strict-origin-when-cross-origin")
        
        c.Next()
    }
}
```

### CORS (Implementado ✅)

```go
func CORSConfig() cors.Config {
    return cors.Config{
        AllowOrigins:     []string{"https://app.aluguei.com"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:          12 * time.Hour,
    }
}
```

### Configuração de Produção

#### Environment Variables
```bash
# Segurança
JWT_SECRET=<32+ caracteres aleatórios>
BCRYPT_COST=12
SESSION_TIMEOUT=24h

# Database
DB_SSL_MODE=require
DB_MAX_CONNECTIONS=20
DB_CONNECTION_TIMEOUT=30s

# Rate Limiting
RATE_LIMIT_ENABLED=true
RATE_LIMIT_REQUESTS=60
RATE_LIMIT_WINDOW=1m

# Logging
LOG_LEVEL=info
LOG_FORMAT=json
SECURITY_LOGS_ENABLED=true
```

---

## 🚨 Plano de Resposta a Incidentes

### Classificação de Incidentes

#### Severidade
1. **Crítica**: Vazamento de dados, acesso não autorizado
2. **Alta**: Tentativas de invasão, vulnerabilidades críticas
3. **Média**: Comportamento suspeito, vulnerabilidades menores
4. **Baixa**: Eventos de segurança rotineiros

#### Processo de Resposta
```
1. DETECÇÃO (0-15 min)
   - Alertas automáticos
   - Análise inicial
   - Classificação de severidade

2. CONTENÇÃO (15-60 min)
   - Isolar sistemas afetados
   - Bloquear acessos suspeitos
   - Preservar evidências

3. ERRADICAÇÃO (1-4 horas)
   - Identificar causa raiz
   - Corrigir vulnerabilidades
   - Atualizar sistemas

4. RECUPERAÇÃO (4-24 horas)
   - Restaurar serviços
   - Monitorar atividade
   - Validar correções

5. LIÇÕES APRENDIDAS (1-7 dias)
   - Documentar incidente
   - Melhorar processos
   - Atualizar políticas
```

### Contatos de Emergência
- **Equipe Técnica**: Disponível 24/7
- **DPO (Data Protection Officer)**: Para questões LGPD
- **Jurídico**: Para questões legais
- **Comunicação**: Para comunicação externa

---

## 📋 Checklist de Segurança

### Desenvolvimento ✅
- [x] Validação de input em todos os endpoints
- [x] Sanitização de dados de saída
- [x] Hash seguro de senhas (bcrypt)
- [x] JWT com expiração adequada
- [x] Middleware de autenticação
- [x] Controle de acesso por recurso
- [x] Headers de segurança
- [x] CORS configurado
- [x] Logs de segurança
- [x] Tratamento seguro de erros

### Produção 🔄
- [ ] HTTPS obrigatório (TLS 1.3)
- [ ] Rate limiting implementado
- [ ] WAF (Web Application Firewall)
- [ ] Monitoramento de segurança
- [ ] Backup criptografado
- [ ] Disaster recovery plan
- [ ] Penetration testing
- [ ] Security audit
- [ ] Compliance assessment
- [ ] Incident response plan

### LGPD ✅
- [x] Mapeamento de dados pessoais
- [x] Base legal definida
- [x] Direitos do titular implementados
- [x] Política de privacidade
- [x] Termo de consentimento
- [x] Processo de anonimização
- [x] Retenção de dados definida
- [x] DPO designado
- [x] Registro de atividades
- [x] Avaliação de impacto

A segurança no sistema Aluguei é tratada como um processo contínuo, com revisões regulares e atualizações conforme novas ameaças e regulamentações surgem.