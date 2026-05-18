// LoginCredentials
export interface LoginCredentials {
  email: string;
  password: string;
  rememberMe?: boolean;
}

// BackendUser
export interface BackendUser {
  id: string;
  first_name: string;
  last_name: string;
  email: string;
  cpf: string;
  phone?: string;
  role: string;
}

/** Dados usados na UI — sem CPF persistido */
export interface User {
  id: string;
  name: string;
  email: string;
  role: string;
  avatar?: string;
  phone?: string;
}

// StoredUser
export interface StoredUser {
  id: string;
  name: string;
  email: string;
  role: string;
}

// AuthResponse
export interface AuthResponse {
  success: boolean;
  user?: BackendUser;
  message?: string;
}
