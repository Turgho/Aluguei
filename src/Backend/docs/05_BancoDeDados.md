## 🗃️ **MODELAGEM DE BANCO DE DADOS**

```sql
-- Tabelas Principais
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    phone VARCHAR(20),
    role VARCHAR(20) CHECK (role IN ('owner', 'tenant', 'admin')),
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE properties (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    owner_id UUID REFERENCES users(id),
    address_street VARCHAR(255),
    address_number VARCHAR(10),
    address_complement VARCHAR(100),
    address_neighborhood VARCHAR(100),
    address_city VARCHAR(100),
    address_state VARCHAR(2),
    address_zip_code VARCHAR(9),
    address_lat DECIMAL(10, 8),
    address_lng DECIMAL(11, 8),
    type VARCHAR(20) CHECK (type IN ('apartment', 'house', 'commercial')),
    bedrooms INTEGER,
    bathrooms INTEGER,
    area DECIMAL(10, 2),
    rent_amount DECIMAL(10, 2),
    condo_fee DECIMAL(10, 2),
    iptu DECIMAL(10, 2),
    status VARCHAR(20) DEFAULT 'available',
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE lease_contracts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    property_id UUID REFERENCES properties(id),
    tenant_id UUID REFERENCES users(id),
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    monthly_rent DECIMAL(10, 2) NOT NULL,
    deposit_amount DECIMAL(10, 2),
    payment_due_day INTEGER DEFAULT 5,
    status VARCHAR(20) DEFAULT 'active',
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE payments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    contract_id UUID REFERENCES lease_contracts(id),
    due_date DATE NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    paid_amount DECIMAL(10, 2),
    paid_date TIMESTAMP,
    status VARCHAR(20) DEFAULT 'pending',
    late_fee DECIMAL(10, 2) DEFAULT 0,
    method VARCHAR(20),
    reference_code VARCHAR(100),
    created_at TIMESTAMP DEFAULT NOW()
);

-- Índices para Performance
CREATE INDEX idx_properties_owner ON properties(owner_id);
CREATE INDEX idx_contracts_tenant ON lease_contracts(tenant_id);
CREATE INDEX idx_payments_due_date ON payments(due_date);
CREATE INDEX idx_payments_status ON payments(status);
```