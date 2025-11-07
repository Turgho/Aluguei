# 🔄 Fluxos do Sistema Aluguei

## Visão Geral

Este documento apresenta os principais fluxos de negócio do sistema Aluguei, desde o cadastro inicial até a gestão completa de aluguéis.

---

## 1. Fluxo de Cadastro e Autenticação

### 1.1 Cadastro de Proprietário

```mermaid
flowchart TD
    A[Usuário acessa sistema] --> B[Clica em 'Cadastrar']
    B --> C[Preenche formulário]
    C --> D{Dados válidos?}
    D -->|Não| E[Exibe erros de validação]
    E --> C
    D -->|Sim| F{Email já existe?}
    F -->|Sim| G[Erro: Email já cadastrado]
    G --> C
    F -->|Não| H{CPF já existe?}
    H -->|Sim| I[Erro: CPF já cadastrado]
    I --> C
    H -->|Não| J[Criptografa senha]
    J --> K[Salva no banco]
    K --> L[Envia email de confirmação]
    L --> M[Redireciona para login]
```

### 1.2 Login de Usuário

```mermaid
flowchart TD
    A[Usuário acessa login] --> B[Insere email e senha]
    B --> C{Credenciais válidas?}
    C -->|Não| D[Erro: Credenciais inválidas]
    D --> B
    C -->|Sim| E[Gera token JWT]
    E --> F[Retorna token + dados do usuário]
    F --> G[Redireciona para dashboard]
```

---

## 2. Fluxo de Gestão de Propriedades

### 2.1 Cadastro de Propriedade

```mermaid
flowchart TD
    A[Proprietário logado] --> B[Acessa 'Minhas Propriedades']
    B --> C[Clica em 'Nova Propriedade']
    C --> D[Preenche dados da propriedade]
    D --> E{Dados válidos?}
    E -->|Não| F[Exibe erros de validação]
    F --> D
    E -->|Sim| G[Salva propriedade]
    G --> H[Status: 'Disponível']
    H --> I[Exibe na lista de propriedades]
```

### 2.2 Alteração de Status da Propriedade

```mermaid
flowchart TD
    A[Proprietário seleciona propriedade] --> B[Clica em 'Alterar Status']
    B --> C{Status atual?}
    C -->|Disponível| D[Pode alterar para: Alugada, Manutenção, Inativa]
    C -->|Alugada| E[Pode alterar para: Manutenção, Inativa]
    C -->|Manutenção| F[Pode alterar para: Disponível, Inativa]
    C -->|Inativa| G[Pode alterar para: Disponível, Manutenção]
    
    D --> H[Seleciona novo status]
    E --> H
    F --> H
    G --> H
    
    H --> I{Tem contrato ativo?}
    I -->|Sim e mudando para Disponível| J[Erro: Propriedade tem contrato ativo]
    I -->|Não ou status permitido| K[Atualiza status]
    K --> L[Confirma alteração]
```

---

## 3. Fluxo de Gestão de Inquilinos

### 3.1 Cadastro de Inquilino

```mermaid
flowchart TD
    A[Proprietário logado] --> B[Acessa 'Inquilinos']
    B --> C[Clica em 'Novo Inquilino']
    C --> D[Preenche dados do inquilino]
    D --> E{Dados válidos?}
    E -->|Não| F[Exibe erros de validação]
    F --> D
    E -->|Sim| G{Email já existe?}
    G -->|Sim| H[Erro: Email já cadastrado]
    H --> D
    G -->|Não| I{CPF já existe?}
    I -->|Sim| J[Erro: CPF já cadastrado]
    J --> D
    I -->|Não| K[Salva inquilino]
    K --> L[Vincula ao proprietário]
    L --> M[Exibe na lista de inquilinos]
```

---

## 4. Fluxo de Criação de Contrato

### 4.1 Novo Contrato de Aluguel

```mermaid
flowchart TD
    A[Proprietário logado] --> B[Acessa 'Contratos']
    B --> C[Clica em 'Novo Contrato']
    C --> D[Seleciona propriedade]
    D --> E{Propriedade disponível?}
    E -->|Não| F[Erro: Propriedade não disponível]
    F --> D
    E -->|Sim| G[Seleciona inquilino]
    G --> H[Define datas e valores]
    H --> I{Dados válidos?}
    I -->|Não| J[Exibe erros de validação]
    J --> H
    I -->|Sim| K[Cria contrato]
    K --> L[Atualiza status da propriedade para 'Alugada']
    L --> M[Gera pagamentos automáticos]
    M --> N[Confirma criação]
```

### 4.2 Geração Automática de Pagamentos

```mermaid
flowchart TD
    A[Contrato criado/ativado] --> B[Calcula data do primeiro pagamento]
    B --> C[Data início + dia de vencimento]
    C --> D[Cria pagamento com status 'Pendente']
    D --> E[Próximo mês?]
    E -->|Sim e dentro do período| F[Calcula próxima data]
    F --> G[Cria próximo pagamento]
    G --> E
    E -->|Não ou fora do período| H[Finaliza geração]
```

---

## 5. Fluxo de Gestão de Pagamentos

### 5.1 Registro de Pagamento

```mermaid
flowchart TD
    A[Proprietário acessa pagamentos] --> B[Seleciona pagamento pendente]
    B --> C[Clica em 'Registrar Pagamento']
    C --> D[Informa data e valor pago]
    D --> E{Valor correto?}
    E -->|Valor total| F[Status: 'Pago']
    E -->|Valor parcial| G[Status: 'Parcial']
    E -->|Valor maior| H[Erro: Valor não pode ser maior]
    H --> D
    F --> I[Salva pagamento]
    G --> I
    I --> J[Atualiza histórico]
    J --> K[Confirma registro]
```

### 5.2 Identificação de Atrasos (Processo Automático)

```mermaid
flowchart TD
    A[Job diário executado] --> B[Busca pagamentos pendentes]
    B --> C[Verifica data de vencimento]
    C --> D{Data vencida?}
    D -->|Não| E[Mantém status 'Pendente']
    D -->|Sim| F[Atualiza status para 'Atrasado']
    F --> G[Calcula multa e juros]
    G --> H[Atualiza valor total]
    H --> I[Registra no log de atrasos]
    I --> J[Envia notificação]
```

---

## 6. Fluxo de Cancelamento de Contrato

### 6.1 Encerramento de Locação

```mermaid
flowchart TD
    A[Proprietário seleciona contrato] --> B[Clica em 'Cancelar Contrato']
    B --> C{Tem pagamentos pendentes?}
    C -->|Sim| D[Exibe aviso sobre pendências]
    D --> E[Confirma cancelamento mesmo assim?]
    E -->|Não| F[Cancela operação]
    E -->|Sim| G[Prossegue com cancelamento]
    C -->|Não| G
    G --> H[Atualiza status para 'Cancelado']
    H --> I[Atualiza propriedade para 'Disponível']
    I --> J[Cancela pagamentos futuros]
    J --> K[Registra data de cancelamento]
    K --> L[Confirma cancelamento]
```

---

## 7. Fluxos de Consulta e Relatórios

### 7.1 Dashboard do Proprietário

```mermaid
flowchart TD
    A[Proprietário acessa dashboard] --> B[Carrega métricas gerais]
    B --> C[Total de propriedades por status]
    C --> D[Total de contratos ativos]
    D --> E[Receita mensal atual]
    E --> F[Pagamentos em atraso]
    F --> G[Próximos vencimentos]
    G --> H[Gráfico de receita dos últimos 12 meses]
    H --> I[Exibe dashboard completo]
```

### 7.2 Relatório de Inadimplência

```mermaid
flowchart TD
    A[Proprietário acessa relatórios] --> B[Seleciona 'Inadimplência']
    B --> C[Define período de análise]
    C --> D[Busca pagamentos em atraso]
    D --> E[Agrupa por inquilino]
    E --> F[Calcula total em atraso]
    F --> G[Calcula dias de atraso médio]
    G --> H[Gera lista de ações sugeridas]
    H --> I[Exibe relatório]
    I --> J[Opção de exportar PDF/Excel]
```

---

## 8. Fluxos de Validação e Segurança

### 8.1 Middleware de Autenticação

```mermaid
flowchart TD
    A[Request recebida] --> B[Verifica header Authorization]
    B --> C{Token presente?}
    C -->|Não| D[Retorna 401 Unauthorized]
    C -->|Sim| E[Extrai token JWT]
    E --> F{Token válido?}
    F -->|Não| G[Retorna 401 Invalid Token]
    F -->|Sim| H{Token expirado?}
    H -->|Sim| I[Retorna 401 Token Expired]
    H -->|Não| J[Extrai dados do usuário]
    J --> K[Adiciona ao contexto]
    K --> L[Prossegue para handler]
```

### 8.2 Validação de Dados de Entrada

```mermaid
flowchart TD
    A[Dados recebidos] --> B[Validação de formato]
    B --> C{Formato válido?}
    C -->|Não| D[Retorna 400 Bad Request]
    C -->|Sim| E[Validação de regras de negócio]
    E --> F{Regras atendidas?}
    F -->|Não| G[Retorna 422 Validation Error]
    F -->|Sim| H[Validação de unicidade]
    H --> I{Dados únicos?}
    I -->|Não| J[Retorna 409 Conflict]
    I -->|Sim| K[Prossegue com operação]
```

---

## 9. Fluxos de Integração (MVP 2/3)

### 9.1 Notificação por Email (Planejado)

```mermaid
flowchart TD
    A[Evento disparado] --> B{Tipo de evento?}
    B -->|Vencimento próximo| C[Template: Lembrete]
    B -->|Pagamento atrasado| D[Template: Cobrança]
    B -->|Contrato criado| E[Template: Boas-vindas]
    
    C --> F[Busca dados do destinatário]
    D --> F
    E --> F
    
    F --> G[Monta email personalizado]
    G --> H[Envia via serviço de email]
    H --> I{Enviado com sucesso?}
    I -->|Sim| J[Registra log de sucesso]
    I -->|Não| K[Registra log de erro]
    K --> L[Agenda nova tentativa]
```

### 9.2 Pagamento via PIX (Planejado)

```mermaid
flowchart TD
    A[Inquilino seleciona PIX] --> B[Sistema gera cobrança]
    B --> C[Cria QR Code dinâmico]
    C --> D[Exibe para inquilino]
    D --> E[Inquilino efetua pagamento]
    E --> F[Webhook recebe confirmação]
    F --> G[Valida dados do pagamento]
    G --> H{Valor correto?}
    H -->|Não| I[Registra discrepância]
    H -->|Sim| J[Atualiza status para 'Pago']
    J --> K[Registra data e método]
    K --> L[Envia confirmação]
```

---

## 10. Fluxos de Manutenção e Monitoramento

### 10.1 Health Check

```mermaid
flowchart TD
    A[Request /health] --> B[Verifica conexão com banco]
    B --> C{Banco acessível?}
    C -->|Não| D[Status: Unhealthy]
    C -->|Sim| E[Verifica cache Redis]
    E --> F{Redis acessível?}
    F -->|Não| G[Status: Degraded]
    F -->|Sim| H[Verifica serviços externos]
    H --> I{Serviços OK?}
    I -->|Não| J[Status: Degraded]
    I -->|Sim| K[Status: Healthy]
    
    D --> L[Retorna 503]
    G --> M[Retorna 200 com warnings]
    J --> M
    K --> N[Retorna 200 OK]
```

### 10.2 Backup Automático

```mermaid
flowchart TD
    A[Cron job executado] --> B[Verifica espaço em disco]
    B --> C{Espaço suficiente?}
    C -->|Não| D[Alerta: Espaço insuficiente]
    C -->|Sim| E[Executa pg_dump]
    E --> F{Backup criado?}
    F -->|Não| G[Alerta: Falha no backup]
    F -->|Sim| H[Comprime arquivo]
    H --> I[Move para storage]
    I --> J[Remove backups antigos]
    J --> K[Registra log de sucesso]
```

Estes fluxos garantem que todas as operações do sistema sejam executadas de forma consistente e segura, proporcionando uma experiência confiável para proprietários e inquilinos.