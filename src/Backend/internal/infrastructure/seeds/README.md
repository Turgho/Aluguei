# Database Seeds

## Sample Data Overview

### 👥 Owners (3)
- **João Silva** - joao.silva@email.com (password: 123456)
- **Maria Santos** - maria.santos@email.com (password: 123456)  
- **Carlos Oliveira** - carlos.oliveira@email.com (password: 123456)

### 🏠 Properties (4)
- **Apartamento Centro** - 2 quartos, R$ 1.800 (Available)
- **Casa Vila Madalena** - 3 quartos, R$ 3.200 (Rented)
- **Kitnet Liberdade** - 1 quarto, R$ 950 (Available)
- **Cobertura Jardins** - 4 quartos, R$ 8.500 (Maintenance)

### 👤 Tenants (3)
- **Ana Costa** - ana.costa@email.com
- **Pedro Almeida** - pedro.almeida@email.com
- **Lucia Ferreira** - lucia.ferreira@email.com

### 📋 Contracts (2)
- Casa Vila Madalena ↔ Ana Costa (Active)
- Kitnet Liberdade ↔ Lucia Ferreira (Expired)

### 💰 Payments (4)
- Jan/2024 - R$ 3.200 (Paid)
- Feb/2024 - R$ 3.200 (Paid)
- Mar/2024 - R$ 3.200 (Pending)
- Jan/2024 - R$ 3.200 (Overdue)

## Usage

```bash
# Run seeds
make seed

# Or directly
go run cmd/seed/main.go
```

## Features

- **Idempotent**: Won't duplicate data if run multiple times
- **Realistic**: Sample data represents real-world scenarios
- **Complete**: Covers all entity relationships
- **Testable**: Provides data for API testing