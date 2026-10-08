import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';
import { from } from 'rxjs';
import { map } from 'rxjs/operators';
import { AuthService } from '../auth/auth.service';

/**
 * Rotas privadas: valida sessão no servidor.
 * fetchMe() renova sessão em /auth/me (o interceptor ignora essa rota de propósito).
 */
export const authGuard: CanActivateFn = () => {
  const auth = inject(AuthService);
  const router = inject(Router);

  return from(auth.fetchMe()).pipe(
    map(ok => (ok ? true : router.createUrlTree(['/login']))),
  );
};

/**
 * Rotas públicas: se ainda há sessão (incluindo via refresh), manda ao dashboard.
 */
export const guestGuard: CanActivateFn = () => {
  const auth = inject(AuthService);
  const router = inject(Router);

  return from(auth.fetchMe(true)).pipe(
    map(ok => (ok ? router.createUrlTree(['/dashboard']) : true)),
  );
};
