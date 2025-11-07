# 📖 User Stories - Sistema Aluguei

## Metodologia

As user stories seguem o formato: **"Como [persona], eu quero [funcionalidade] para [benefício]"**

Cada story possui:
- **Critérios de Aceitação** (Given/When/Then)
- **Prioridade** (Alta/Média/Baixa)
- **Status** (✅ Implementado / 🔄 Em Desenvolvimento / 📋 Backlog)

---

## 👤 Personas

### Proprietário (Owner)
João Silva, 45 anos, possui 3 apartamentos para alugar. Quer centralizar a gestão e reduzir inadimplência.

### Inquilino (Tenant)
Maria Santos, 28 anos, aluga apartamento. Quer transparência nos pagamentos e facilidade de comunicação.

---

## 🔐 Épico: Autenticação e Segurança

### US001 - Login de Proprietário ✅
**Como** proprietário  
**Eu quero** fazer login no sistema  
**Para** acessar minhas propriedades e inquilinos  

**Critérios de Aceitação:**
- Dado que sou um proprietário cadastrado
- Quando informo email e senha corretos
- Então devo ser autenticado e receber um token JWT
- E devo ser redirecionado para o dashboard

**Prioridade:** Alta  
**Status:** ✅ Implementado

### US002 - Logout Seguro ✅
**Como** proprietário logado  
**Eu quero** fazer logout do sistema  
**Para** garantir a segurança da minha conta  

**Critérios de Aceitação:**
- Dado que estou logado no sistema
- Quando clico em logout
- Então meu token deve ser invalidado
- E devo ser redirecionado para a tela de login

**Prioridade:** Alta  
**Status:** ✅ Implementado

---

## 🏠 Épico: Gestão de Proprietários

### US003 - Cadastro de Proprietário ✅
**Como** novo usuário  
**Eu quero** me cadastrar como proprietário  
**Para** começar a usar o sistema  

**Critérios de Aceitação:**
- Dado que não tenho conta no sistema
- Quando preencho nome, email, telefone, CPF e senha
- Então minha conta deve ser criada
- E devo receber confirmação por email

**Prioridade:** Alta  
**Status:** ✅ Implementado

### US004 - Atualizar Perfil ✅
**Como** proprietário  
**Eu quero** atualizar meus dados pessoais  
**Para** manter informações atualizadas  

**Critérios de Aceitação:**
- Dado que estou logado como proprietário
- Quando altero meus dados pessoais
- Então as informações devem ser atualizadas
- E devo receber confirmação da alteração

**Prioridade:** Média  
**Status:** ✅ Implementado

### US005 - Visualizar Perfil ✅
**Como** proprietário  
**Eu quero** visualizar meu perfil completo  
**Para** verificar minhas informações  

**Critérios de Aceitação:**
- Dado que estou logado como proprietário
- Quando acesso meu perfil
- Então devo ver todos os meus dados
- E o histórico de propriedades cadastradas

**Prioridade:** Baixa  
**Status:** ✅ Implementado

---

## 🏢 Épico: Gestão de Propriedades

### US006 - Cadastrar Propriedade ✅
**Como** proprietário  
**Eu quero** cadastrar uma nova propriedade  
**Para** disponibilizá-la para aluguel  

**Critérios de Aceitação:**
- Dado que estou logado como proprietário
- Quando preencho dados da propriedade (título, endereço, valor)
- Então a propriedade deve ser cadastrada
- E deve aparecer na minha lista de propriedades

**Prioridade:** Alta  
**Status:** ✅ Implementado

### US007 - Listar Propriedades ✅
**Como** proprietário  
**Eu quero** ver todas as minhas propriedades  
**Para** ter uma visão geral do meu portfólio  

**Critérios de Aceitação:**
- Dado que tenho propriedades cadastradas
- Quando acesso a lista de propriedades
- Então devo ver todas com status e informações básicas
- E poder filtrar por status (disponível, alugada, manutenção)

**Prioridade:** Alta  
**Status:** ✅ Implementado

### US008 - Editar Propriedade ✅
**Como** proprietário  
**Eu quero** editar dados de uma propriedade  
**Para** manter informações atualizadas  

**Critérios de Aceitação:**
- Dado que tenho uma propriedade cadastrada
- Quando altero seus dados
- Então as informações devem ser atualizadas
- E o histórico de alterações deve ser mantido

**Prioridade:** Média  
**Status:** ✅ Implementado

### US009 - Alterar Status da Propriedade ✅
**Como** proprietário  
**Eu quero** alterar o status de uma propriedade  
**Para** refletir sua situação atual  

**Critérios de Aceitação:**
- Dado que tenho uma propriedade cadastrada
- Quando altero seu status (disponível/alugada/manutenção)
- Então o status deve ser atualizado
- E deve impactar na disponibilidade para novos contratos

**Prioridade:** Alta  
**Status:** ✅ Implementado

### US010 - Remover Propriedade ✅
**Como** proprietário  
**Eu quero** remover uma propriedade  
**Para** limpar propriedades que não possuo mais  

**Critérios de Aceitação:**
- Dado que tenho uma propriedade sem contratos ativos
- Quando solicito sua remoção
- Então ela deve ser removida do sistema
- E não deve aparecer mais nas minhas listas

**Prioridade:** Baixa  
**Status:** ✅ Implementado

---

## 👥 Épico: Gestão de Inquilinos

### US011 - Cadastrar Inquilino ✅
**Como** proprietário  
**Eu quero** cadastrar um novo inquilino  
**Para** poder criar contratos de aluguel  

**Critérios de Aceitação:**
- Dado que estou logado como proprietário
- Quando preencho dados do inquilino (nome, email, CPF, telefone)
- Então o inquilino deve ser cadastrado
- E deve aparecer na minha lista de inquilinos

**Prioridade:** Alta  
**Status:** ✅ Implementado

### US012 - Listar Inquilinos ✅
**Como** proprietário  
**Eu quero** ver todos os meus inquilinos  
**Para** gerenciar meus relacionamentos  

**Critérios de Aceitação:**
- Dado que tenho inquilinos cadastrados
- Quando acesso a lista de inquilinos
- Então devo ver todos com informações básicas
- E poder buscar por nome ou CPF

**Prioridade:** Alta  
**Status:** ✅ Implementado

### US013 - Editar Inquilino ✅
**Como** proprietário  
**Eu quero** editar dados de um inquilino  
**Para** manter informações atualizadas  

**Critérios de Aceitação:**
- Dado que tenho um inquilino cadastrado
- Quando altero seus dados
- Então as informações devem ser atualizadas
- E o histórico deve ser preservado

**Prioridade:** Média  
**Status:** ✅ Implementado

### US014 - Visualizar Histórico do Inquilino ✅
**Como** proprietário  
**Eu quero** ver o histórico de um inquilino  
**Para** avaliar seu comportamento como locatário  

**Critérios de Aceitação:**
- Dado que tenho um inquilino com histórico
- Quando acesso seus detalhes
- Então devo ver contratos anteriores e atuais
- E histórico de pagamentos

**Prioridade:** Média  
**Status:** ✅ Implementado

---

## 📄 Épico: Gestão de Contratos

### US015 - Criar Contrato ✅
**Como** proprietário  
**Eu quero** criar um contrato de aluguel  
**Para** formalizar a locação de uma propriedade  

**Critérios de Aceitação:**
- Dado que tenho propriedade disponível e inquilino cadastrado
- Quando crio um contrato com datas e valor
- Então o contrato deve ser criado
- E a propriedade deve ficar com status "alugada"

**Prioridade:** Alta  
**Status:** ✅ Implementado

### US016 - Listar Contratos ✅
**Como** proprietário  
**Eu quero** ver todos os meus contratos  
**Para** acompanhar locações ativas e históricas  

**Critérios de Aceitação:**
- Dado que tenho contratos cadastrados
- Quando acesso a lista de contratos
- Então devo ver todos com status e informações básicas
- E poder filtrar por status (ativo, expirado, cancelado)

**Prioridade:** Alta  
**Status:** ✅ Implementado

### US017 - Visualizar Detalhes do Contrato ✅
**Como** proprietário  
**Eu quero** ver detalhes completos de um contrato  
**Para** acompanhar termos e condições  

**Critérios de Aceitação:**
- Dado que tenho um contrato cadastrado
- Quando acesso seus detalhes
- Então devo ver todas as informações
- E histórico de pagamentos relacionados

**Prioridade:** Média  
**Status:** ✅ Implementado

### US018 - Atualizar Contrato ✅
**Como** proprietário  
**Eu quero** atualizar dados de um contrato  
**Para** refletir mudanças acordadas  

**Critérios de Aceitação:**
- Dado que tenho um contrato ativo
- Quando altero valor do aluguel ou data de vencimento
- Então as informações devem ser atualizadas
- E os próximos pagamentos devem refletir as mudanças

**Prioridade:** Média  
**Status:** ✅ Implementado

### US019 - Cancelar Contrato ✅
**Como** proprietário  
**Eu quero** cancelar um contrato  
**Para** encerrar uma locação  

**Critérios de Aceitação:**
- Dado que tenho um contrato ativo
- Quando cancelo o contrato
- Então seu status deve mudar para "cancelado"
- E a propriedade deve ficar disponível novamente

**Prioridade:** Alta  
**Status:** ✅ Implementado

---

## 💰 Épico: Gestão de Pagamentos

### US020 - Registrar Pagamento ✅
**Como** proprietário  
**Eu quero** registrar um pagamento recebido  
**Para** manter controle financeiro atualizado  

**Critérios de Aceitação:**
- Dado que tenho um pagamento pendente
- Quando registro o pagamento com data e valor
- Então o status deve mudar para "pago"
- E deve aparecer no histórico financeiro

**Prioridade:** Alta  
**Status:** ✅ Implementado

### US021 - Listar Pagamentos ✅
**Como** proprietário  
**Eu quero** ver todos os pagamentos  
**Para** acompanhar minha situação financeira  

**Critérios de Aceitação:**
- Dado que tenho pagamentos registrados
- Quando acesso a lista de pagamentos
- Então devo ver todos com status e datas
- E poder filtrar por período e status

**Prioridade:** Alta  
**Status:** ✅ Implementado

### US022 - Visualizar Pagamentos em Atraso ✅
**Como** proprietário  
**Eu quero** ver pagamentos em atraso  
**Para** tomar ações de cobrança  

**Critérios de Aceitação:**
- Dado que tenho pagamentos vencidos não pagos
- Quando acesso relatório de inadimplência
- Então devo ver lista de pagamentos em atraso
- E informações de contato dos inquilinos

**Prioridade:** Alta  
**Status:** ✅ Implementado

### US023 - Gerar Pagamentos Automáticos ✅
**Como** proprietário  
**Eu quero** que pagamentos sejam gerados automaticamente  
**Para** não precisar criar manualmente todo mês  

**Critérios de Aceitação:**
- Dado que tenho um contrato ativo
- Quando chega o dia de vencimento
- Então um novo pagamento deve ser criado automaticamente
- E deve aparecer como pendente

**Prioridade:** Média  
**Status:** ✅ Implementado

---

## 📊 Épico: Relatórios e Dashboard (MVP 2)

### US024 - Dashboard Proprietário 🔄
**Como** proprietário  
**Eu quero** ver um dashboard com métricas  
**Para** ter visão geral do meu negócio  

**Critérios de Aceitação:**
- Dado que estou logado como proprietário
- Quando acesso o dashboard
- Então devo ver métricas de propriedades, contratos e pagamentos
- E gráficos de receita mensal

**Prioridade:** Alta  
**Status:** 🔄 MVP 2

### US025 - Relatório Financeiro 🔄
**Como** proprietário  
**Eu quero** gerar relatórios financeiros  
**Para** acompanhar performance do portfólio  

**Critérios de Aceitação:**
- Dado que tenho histórico de pagamentos
- Quando solicito relatório de período
- Então devo receber relatório com receitas e inadimplência
- E poder exportar em PDF/Excel

**Prioridade:** Média  
**Status:** 🔄 MVP 2

### US026 - Relatório de Inadimplência 🔄
**Como** proprietário  
**Eu quero** relatório detalhado de inadimplência  
**Para** tomar ações específicas  

**Critérios de Aceitação:**
- Dado que tenho pagamentos em atraso
- Quando gero relatório de inadimplência
- Então devo ver detalhes por inquilino e propriedade
- E sugestões de ações de cobrança

**Prioridade:** Média  
**Status:** 🔄 MVP 2

---

## 🔔 Épico: Notificações (MVP 2)

### US027 - Notificação de Vencimento 🔄
**Como** proprietário  
**Eu quero** ser notificado sobre vencimentos  
**Para** acompanhar pagamentos pendentes  

**Critérios de Aceitação:**
- Dado que tenho pagamentos próximos do vencimento
- Quando faltam 3 dias para vencer
- Então devo receber notificação por email
- E o inquilino também deve ser notificado

**Prioridade:** Alta  
**Status:** 🔄 MVP 2

### US028 - Notificação de Atraso 🔄
**Como** proprietário  
**Eu quero** ser notificado sobre atrasos  
**Para** tomar ações de cobrança rapidamente  

**Critérios de Aceitação:**
- Dado que um pagamento está em atraso
- Quando passa 1 dia do vencimento
- Então devo receber notificação de atraso
- E sugestões de ações de cobrança

**Prioridade:** Alta  
**Status:** 🔄 MVP 2

---

## 📱 Épico: Portal do Inquilino (MVP 2)

### US029 - Login do Inquilino 🔄
**Como** inquilino  
**Eu quero** acessar meu portal  
**Para** ver informações dos meus aluguéis  

**Critérios de Aceitação:**
- Dado que sou inquilino cadastrado
- Quando faço login com email e senha
- Então devo acessar meu portal pessoal
- E ver meus contratos ativos

**Prioridade:** Alta  
**Status:** 🔄 MVP 2

### US030 - Histórico de Pagamentos 🔄
**Como** inquilino  
**Eu quero** ver meu histórico de pagamentos  
**Para** acompanhar minha situação financeira  

**Critérios de Aceitação:**
- Dado que tenho pagamentos registrados
- Quando acesso meu histórico
- Então devo ver todos os pagamentos com datas e valores
- E poder baixar comprovantes

**Prioridade:** Média  
**Status:** 🔄 MVP 2

---

## 🚀 Épico: Integrações (MVP 3)

### US031 - Pagamento via PIX 📋
**Como** inquilino  
**Eu quero** pagar aluguel via PIX  
**Para** ter mais praticidade no pagamento  

**Critérios de Aceitação:**
- Dado que tenho pagamento pendente
- Quando escolho pagar via PIX
- Então devo receber QR Code ou chave PIX
- E pagamento deve ser confirmado automaticamente

**Prioridade:** Alta  
**Status:** 📋 MVP 3

### US032 - Assinatura Digital 📋
**Como** proprietário  
**Eu quero** contratos com assinatura digital  
**Para** eliminar processos físicos  

**Critérios de Aceitação:**
- Dado que crio um novo contrato
- Quando envio para assinatura
- Então inquilino deve receber link para assinar digitalmente
- E contrato deve ser válido juridicamente

**Prioridade:** Média  
**Status:** 📋 MVP 3

---

## 📈 Métricas de Sucesso

### Métricas de Produto
- **Adoção**: 80% dos proprietários cadastram pelo menos 1 propriedade
- **Engajamento**: 60% dos usuários acessam o sistema semanalmente
- **Retenção**: 70% dos usuários permanecem ativos após 3 meses

### Métricas de Negócio
- **Eficiência**: 50% de redução no tempo de gestão
- **Inadimplência**: 30% de redução nos atrasos
- **Satisfação**: NPS > 60 entre proprietários e inquilinos

### Métricas Técnicas
- **Performance**: 95% das requests < 200ms
- **Disponibilidade**: 99.5% de uptime
- **Qualidade**: 0 bugs críticos em produção

Este backlog de user stories guia o desenvolvimento incremental do sistema, priorizando valor para o usuário e validação de hipóteses de negócio.