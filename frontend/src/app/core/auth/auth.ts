// src/app/core/auth/auth.ts
import { Injectable, inject, signal, computed } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { firstValueFrom } from 'rxjs';
import { environment } from '../../../environment/environment';

// ── Interfaces ────────────────────────────────────────────────

export interface LoginCredentials {
  email: string;
  password: string;
}

export interface AuthResponse {
  success: boolean;
  user?: User;
  message?: string;
}

export interface User {
  id: string;
  name: string;
  email: string;
  avatar?: string;
}

// Só o usuário no storage — tokens ficam nos cookies httpOnly
const USER_KEY = 'auth_user';

// ESSENCIAL: envia e recebe cookies em toda requisição
const HTTP_OPTIONS = { withCredentials: true };

@Injectable({ providedIn: 'root' })
export class AuthService {
  private http = inject(HttpClient);

  private _user = signal<User | null>(this.loadUser());

  readonly user         = this._user.asReadonly();
  readonly isLoggedIn   = computed(() => !!this._user());
  readonly userInitials = computed(() => {
    const name = this._user()?.name ?? '';
    return name.split(' ').map(n => n[0]).slice(0, 2).join('').toUpperCase();
  });

  // ── Login ──────────────────────────────────────────────────
  async login(credentials: LoginCredentials): Promise<void> {
    const res = await firstValueFrom(
      this.http.post<AuthResponse>(
        `${environment.apiUrl}/auth/login`,
        credentials,
        HTTP_OPTIONS,
      )
    );

    if (!res.success) {
      throw new Error(res.message ?? 'Falha no login');
    }

    if (res.user) {
      this.saveUser(res.user);
    }
  }

  // ── Verificar sessão ativa ─────────────────────────────────
  // Chama /auth/me — o cookie vai automaticamente.
  // Retorna true e salva o usuário em memória, ou false se não autenticado.
  async fetchMe(): Promise<boolean> {
    try {
      const user = await firstValueFrom(
        this.http.get<User>(`${environment.apiUrl}/auth/me`, HTTP_OPTIONS)
      );
      this._user.set(user);
      return true;
    } catch {
      this._user.set(null);
      return false;
    }
  }

  // ── Logout ─────────────────────────────────────────────────
  async logout(): Promise<void> {
    try {
      await firstValueFrom(
        this.http.post<void>(
          `${environment.apiUrl}/auth/logout`,
          {},
          HTTP_OPTIONS,
        )
      );
    } finally {
      this.clearUser();
    }
  }

  // ── Refresh token ──────────────────────────────────────────
  // O browser envia o refresh_token cookie automaticamente
  async refreshAccessToken(): Promise<void> {
    const res = await firstValueFrom(
      this.http.post<AuthResponse>(
        `${environment.apiUrl}/auth/refresh`,
        {},
        HTTP_OPTIONS,
      )
    );

    if (!res.success) throw new Error('Refresh falhou');
  }

  // ── Helpers privados ───────────────────────────────────────
  private saveUser(user: User): void {
    localStorage.setItem(USER_KEY, JSON.stringify(user));
    this._user.set(user);
  }

  private clearUser(): void {
    localStorage.removeItem(USER_KEY);
    this._user.set(null);
  }

  private loadUser(): User | null {
    try {
      const raw = localStorage.getItem(USER_KEY);
      return raw ? JSON.parse(raw) : null;
    } catch {
      return null;
    }
  }
}