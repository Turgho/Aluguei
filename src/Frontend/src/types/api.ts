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

export interface ApiError {
  error: string;
}