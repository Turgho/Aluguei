// URL base da API - usa variável de ambiente ou fallback para desenvolvimento local
const API_BASE_URL = process.env.EXPO_PUBLIC_API_URL || 'http://192.168.1.6:8080/api/v1';

// Serviço centralizado para chamadas à API
class ApiService {
  private baseURL: string;

  constructor() {
    this.baseURL = API_BASE_URL;
  }

  private async request<T>(endpoint: string, options: RequestInit = {}): Promise<T> {
    const url = `${this.baseURL}${endpoint}`;
    
    const config: RequestInit = {
      headers: {
        'Content-Type': 'application/json',
        ...options.headers,
      },
      ...options,
    };

    const response = await fetch(url, config);
    
    if (!response.ok) {
      // Tenta extrair mensagem de erro do backend, fallback para erro genérico
      const errorData = await response.json().catch(() => ({ error: 'Network error' }));
      throw new Error(errorData.error || `HTTP ${response.status}`);
    }

    return response.json();
  }

  // Auth endpoints
  async login(email: string, password: string): Promise<import('../types/api').LoginResponse> {
    return this.request<import('../types/api').LoginResponse>('/auth/login', {
      method: 'POST',
      body: JSON.stringify({ email, password }),
    });
  }

  // Owner endpoints
  async createOwner(data: import('../types/api').CreateOwnerRequest): Promise<import('../types/api').Owner> {
    return this.request<import('../types/api').Owner>('/owners', {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  // Dashboard endpoints
  async getDashboard(ownerId: string): Promise<import('../types/api').DashboardResponse> {
    return this.request<import('../types/api').DashboardResponse>(`/dashboard/owner/${ownerId}`);
  }
}

export const apiService = new ApiService();