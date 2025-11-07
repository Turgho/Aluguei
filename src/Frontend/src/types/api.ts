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
    birth_date?: string;
    created_at: string;
    updated_at: string;
  };
}

export interface ApiError {
  error: string;
}