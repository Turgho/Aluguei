import { Injectable, inject, signal, computed } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { firstValueFrom } from 'rxjs';
import { environment } from '../../../environments/environment';
import { USER_STORAGE_KEY } from '../constants/auth.constants';
import {
  AuthResponse,
  BackendUser,
  LoginCredentials,
  StoredUser,
  User,
} from './models/user.model';
import { RegisterPayload } from './models/register.model';
import { fromStoredUser, mapBackendUser, toStoredUser } from './user.mapper';
import { getStorageItem, removeStorageItem, setStorageItem } from '../utils/storage.util';

const HTTP_OPTIONS = { withCredentials: true };

@Injectable({ providedIn: 'root' })
export class AuthService {
  private http = inject(HttpClient);
  private _user = signal<User | null>(null);
  private _storage: Storage = localStorage;

  readonly user = this._user.asReadonly();
  readonly isLoggedIn = computed(() => !!this._user());
  // Iniciais do usuário
  readonly userInitials = computed(() => {
    const name = this._user()?.name ?? '';
    return name
      .split(' ')
      .map(n => n[0])
      .filter(Boolean)
      .slice(0, 2)
      .join('')
      .toUpperCase();
  });

  // Registra um novo usuário
  async register(payload: RegisterPayload): Promise<void> {
    if (environment.useMockAuth) {
      await this.mockDelay();
      return;
    }

    await firstValueFrom(
      this.http.post<BackendUser>(
        `${environment.apiUrl}/auth/register`,
        payload,
        HTTP_OPTIONS,
      ),
    );
  }

  async login(credentials: LoginCredentials): Promise<void> {
    this._storage = credentials.rememberMe ? localStorage : sessionStorage;

    // Se está usando mock de autenticação, simula o login
    if (environment.useMockAuth) {
      await this.mockDelay();
      if (!credentials.email || credentials.password.length < 8) {
        throw new Error('Credenciais inválidas');
      }
      // Define o usuário
      this.setUser({
        id: 'mock-1',
        name: 'Usuário Demo',
        email: credentials.email,
        role: 'owner',
      });
      return;
    }

    // Faz a requisição para o backend
    const res = await firstValueFrom(
      this.http.post<AuthResponse>(
        `${environment.apiUrl}/auth/login`,
        { email: credentials.email, password: credentials.password },
        HTTP_OPTIONS,
      ),
    );

    // Se a requisição falhou, lança um erro
    if (!res.success) {
      throw new Error(res.message ?? 'Falha no login');
    }

    if (res.user) {
      this.setUser(mapBackendUser(res.user));
    }
  }

  /** Valida sessão com o backend (ou mock) — fonte da verdade para guards */
  async fetchMe(): Promise<boolean> {
    if (environment.useMockAuth) {
      await this.mockDelay(200);
      const stored = this.loadStoredUser();
      if (!stored) {
        this.setUser(null);
        return false;
      }
      this.setUser(fromStoredUser(stored));
      return true;
    }

    try {
      // Faz a requisição para o backend
      const backendUser = await firstValueFrom(
        this.http.get<BackendUser>(`${environment.apiUrl}/auth/me`, HTTP_OPTIONS),
      );
      this.setUser(mapBackendUser(backendUser));
      return true;
    } catch {
      // Se a requisição falhou, limpa a sessão
      this.setUser(null);
      return false;
    }
  }

  // Logout do usuário
  async logout(): Promise<void> {
    if (!environment.useMockAuth) {
      try {
        await firstValueFrom(
          this.http.post<void>(
            `${environment.apiUrl}/auth/logout`,
            {},
            HTTP_OPTIONS,
          ),
        );
      } catch {
        // Limpa sessão local mesmo se o backend falhar
      }
    }
    this.setUser(null);
  }

  // Atualiza o token de acesso
  async refreshAccessToken(): Promise<void> {
    if (environment.useMockAuth) return;

    const res = await firstValueFrom(
      this.http.post<AuthResponse>(
        `${environment.apiUrl}/auth/refresh`,
        {},
        HTTP_OPTIONS,
      ),
    );

    if (!res.success) throw new Error('Refresh falhou');
  }

  // Define o usuário
  private setUser(user: User | null): void {
    this._user.set(user);
    removeStorageItem(localStorage, USER_STORAGE_KEY);
    removeStorageItem(sessionStorage, USER_STORAGE_KEY);

    // Se o usuário existe, salva no storage
    if (user) {
      setStorageItem(
        this._storage,
        USER_STORAGE_KEY,
        JSON.stringify(toStoredUser(user)),
      );
    }
  }

  // Carrega o usuário armazenado
  private loadStoredUser(): StoredUser | null {
    for (const storage of [sessionStorage, localStorage]) {
      try {
        const raw = getStorageItem(storage, USER_STORAGE_KEY);
        // Se o usuário existe, carrega no storage
        if (raw) {
          this._storage = storage;
          return JSON.parse(raw) as StoredUser;
        }
      } catch {
        // Se o usuário não existe, remove do storage
        removeStorageItem(storage, USER_STORAGE_KEY);
      }
    }
    return null;
  }

  // Simula delay de rede
  private mockDelay(ms = 400): Promise<void> {
    return new Promise(resolve => setTimeout(resolve, ms));
  }
}
