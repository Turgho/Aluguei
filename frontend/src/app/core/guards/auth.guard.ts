import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';
import { from } from 'rxjs';
import { map } from 'rxjs/operators';
import { AuthService } from '../auth/auth.service';

/** Protege rotas privadas — sempre valida sessão no backend/mock */
export const authGuard: CanActivateFn = () => {
  const auth = inject(AuthService);
  const router = inject(Router);

  return from(auth.fetchMe()).pipe(
    map(ok => (ok ? true : router.createUrlTree(['/login']))),
  );
};

/** Bloqueia rotas públicas quando já há sessão ativa */
export const guestGuard: CanActivateFn = () => {
  const auth = inject(AuthService);
  const router = inject(Router);

  return from(auth.fetchMe()).pipe(
    map(ok => (ok ? router.createUrlTree(['/dashboard']) : true)),
  );
};
