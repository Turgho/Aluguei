import { Injectable, inject, signal, computed } from '@angular/core';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
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
  private readonly http = inject(HttpClient);

  private readonly _user = signal<User | null>(this.initUserFromStorage());

  private _storage: Storage = localStorage;

  /** Mutex: várias requisições 401 simultâneas compartilham o mesmo refresh */
  private refreshPromise: Promise<void> | null = null;

  readonly user = this._user.asReadonly();
  readonly isLoggedIn = computed(() => !!this._user());

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

    // Mock de autenticação para desenvolvimento
    if (environment.useMockAuth) {
      await this.mockDelay();
      if (!credentials.email || credentials.password.length < 8) {
        throw new Error('Credenciais inválidas');
      }
      this.setUser({
        id: 'mock-1',
        name: 'Usuário Demo',
        email: credentials.email,
        role: 'owner',
      });
      return;
    }

    const res = await firstValueFrom(
      this.http.post<AuthResponse>(
        `${environment.apiUrl}/auth/login`,
        { email: credentials.email, password: credentials.password },
        HTTP_OPTIONS,
      ),
    );

    if (!res.success) {
      throw new Error(res.message ?? 'Falha no login');
    }

    if (res.user) {
      this.setUser(mapBackendUser(res.user));
    } else {
      await this.fetchMe();
    }
  }

  /**
   * Valida sessão no servidor.
   * /auth/me não passa pelo refresh do interceptor (evita loop) — renovação fica aqui.
   */
  async fetchMe(silent = false): Promise<boolean> {
    if (environment.useMockAuth) {
      await this.mockDelay(200);
      const stored = this.loadStoredUser();
      if (!stored) {
        if (!silent) this.clearUser();
        return false;
      }
      this.setUser(fromStoredUser(stored));
      return true;
    }

    const ok = await this.tryGetMeWithRefresh();
    if (!ok && !silent) {
      this.clearUser();
    }
    return ok;
  }

  /** Uma tentativa de /auth/me; se 401, faz refresh e tenta só mais uma vez */
  private async tryGetMeWithRefresh(): Promise<boolean> {
    try {
      this.setUser(mapBackendUser(await this.getMe()));
      return true;
    } catch (err) {
      if (!this.isUnauthorized(err)) {
        return false;
      }
    }

    try {
      await this.refreshAccessToken();
      this.setUser(mapBackendUser(await this.getMe()));
      return true;
    } catch {
      return false;
    }
  }

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
        // Limpa localmente mesmo se o backend falhar
      }
    }

    this.clearUser();
  }

  async refreshAccessToken(): Promise<void> {
    if (environment.useMockAuth) return;

    if (this.refreshPromise) {
      return this.refreshPromise;
    }

    this.refreshPromise = (async () => {
      try {
        const res = await firstValueFrom(
          this.http.post<AuthResponse>(
            `${environment.apiUrl}/auth/refresh`,
            {},
            HTTP_OPTIONS,
          ),
        );

        if (!res.success) {
          throw new Error('Refresh falhou');
        }
      } finally {
        this.refreshPromise = null;
      }
    })();

    return this.refreshPromise;
  }

  private getMe(): Promise<BackendUser> {
    return firstValueFrom(
      this.http.get<BackendUser>(`${environment.apiUrl}/auth/me`, HTTP_OPTIONS),
    );
  }

  private isUnauthorized(err: unknown): boolean {
    return err instanceof HttpErrorResponse && err.status === 401;
  }

  private initUserFromStorage(): User | null {
    const stored = this.loadStoredUser();
    return stored ? fromStoredUser(stored) : null;
  }

  private setUser(user: User): void {
    this._user.set(user);
    removeStorageItem(localStorage, USER_STORAGE_KEY);
    removeStorageItem(sessionStorage, USER_STORAGE_KEY);
    setStorageItem(
      this._storage,
      USER_STORAGE_KEY,
      JSON.stringify(toStoredUser(user)),
    );
  }

  private clearUser(): void {
    this._user.set(null);
    removeStorageItem(localStorage, USER_STORAGE_KEY);
    removeStorageItem(sessionStorage, USER_STORAGE_KEY);
  }

  private loadStoredUser(): StoredUser | null {
    for (const storage of [sessionStorage, localStorage]) {
      try {
        const raw = getStorageItem(storage, USER_STORAGE_KEY);
        if (raw) {
          this._storage = storage;
          return JSON.parse(raw) as StoredUser;
        }
      } catch {
        removeStorageItem(storage, USER_STORAGE_KEY);
      }
    }
    return null;
  }

  private mockDelay(ms = 400): Promise<void> {
    return new Promise(resolve => setTimeout(resolve, ms));
  }
}
