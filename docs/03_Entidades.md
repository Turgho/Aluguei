## 🗂️ **ENTIDADES PRINCIPAIS**

```go
// Core Entities
type User struct {
    ID          string    `json:"id"`
    Name        string    `json:"name"`
    Email       string    `json:"email"`
    Phone       string    `json:"phone"`
    Role        UserRole  `json:"role"` // owner, tenant, admin
    CreatedAt   time.Time `json:"created_at"`
}

type Property struct {
    ID            string     `json:"id"`
    OwnerID       string     `json:"owner_id"`
    Address       Address    `json:"address"`
    Type          string     `json:"type"`
    RentAmount    float64    `json:"rent_amount"`
    Status        string     `json:"status"`
    Features      []string   `json:"features"`
    CreatedAt     time.Time  `json:"created_at"`
}

type LeaseContract struct {
    ID           string     `json:"id"`
    PropertyID   string     `json:"property_id"`
    TenantID     string     `json:"tenant_id"`
    StartDate    time.Time  `json:"start_date"`
    EndDate      time.Time  `json:"end_date"`
    MonthlyRent  float64    `json:"monthly_rent"`
    Status       string     `json:"status"`
    CreatedAt    time.Time  `json:"created_at"`
}

type Payment struct {
    ID          string     `json:"id"`
    ContractID  string     `json:"contract_id"`
    DueDate     time.Time  `json:"due_date"`
    Amount      float64    `json:"amount"`
    Status      string     `json:"status"`
    PaidDate    *time.Time `json:"paid_date"`
    LateFee     float64    `json:"late_fee"`
    CreatedAt   time.Time  `json:"created_at"`
}
```