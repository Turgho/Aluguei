export interface LoginRequest {
  email: string;
  password: string;
}

export interface LoginResponse {
  token: string;
  expires_at: string;
  owner: {
    id: string;
    name: string;
    email: string;
    phone: string;
    cpf: string;
    birth_date?: string | null;
    created_at: string;
    updated_at: string;
  };
}

export interface CreateOwnerRequest {
  name: string;
  email: string;
  password: string;
  phone: string;
  cpf: string;
  birth_date?: string;
}

export interface Owner {
  id: string;
  name: string;
  email: string;
  phone: string;
  cpf: string;
  birth_date?: string | null;
  created_at: string;
  updated_at: string;
}

export interface DashboardResponse {
  total_properties: number;
  rented_properties: number;
  available_properties: number;
  monthly_revenue: number;
  pending_payments: number;
  overdue_payments: number;
  recent_payments: RecentPayment[];
  monthly_revenues: MonthlyRevenue[];
  property_status: PropertyStatusCount[];
}

export interface RecentPayment {
  id: string;
  tenant: string;
  property: string;
  amount: number;
  date: string;
}

export interface MonthlyRevenue {
  month: string;
  revenue: number;
}

export interface PropertyStatusCount {
  status: string;
  count: number;
}

export interface ApiError {
  error: string;
}