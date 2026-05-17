// src/app/core/guards/auth.ts
import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';
import { from } from 'rxjs';
import { map } from 'rxjs/operators';
import { AuthService } from '../auth/auth';

/**
 * Protege rotas privadas.
 *
 * - Se o usuário já está em memória (signal), libera imediatamente.
 * - Se não está (ex: refresh da página), chama /auth/me para verificar
 *   se o cookie ainda é válido antes de redirecionar pro login.
 */
export const authGuard: CanActivateFn = () => {
  const auth   = inject(AuthService);
  const router = inject(Router);

  // Usuário já carregado em memória — não precisa bater no backend
  if (auth.isLoggedIn()) return true;

  // Página foi recarregada — tenta restaurar sessão via cookie
  return from(auth.fetchMe()).pipe(
    map(ok => ok ? true : router.createUrlTree(['/login']))
  );
};

/**
 * Bloqueia rotas públicas (ex: /login) quando o usuário já está autenticado.
 *
 * - Se está em memória, redireciona pro dashboard imediatamente.
 * - Se não está, verifica com o backend se o cookie ainda é válido.
 *   Sessão ativa → redireciona pro dashboard.
 *   Sem sessão    → libera a rota pública.
 */
export const guestGuard: CanActivateFn = () => {
  const auth   = inject(AuthService);
  const router = inject(Router);

  // Usuário já carregado em memória — redireciona direto
  if (auth.isLoggedIn()) return router.createUrlTree(['/dashboard']);

  // Verifica se há sessão ativa via cookie antes de exibir o login
  return from(auth.fetchMe()).pipe(
    map(ok => ok ? router.createUrlTree(['/dashboard']) : true)
  );
};