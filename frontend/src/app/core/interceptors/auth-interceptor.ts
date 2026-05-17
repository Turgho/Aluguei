// src/app/core/interceptors/auth-interceptor.ts
import { HttpInterceptorFn, HttpErrorResponse } from '@angular/common/http';
import { inject } from '@angular/core';
import { Router } from '@angular/router';
import { catchError, from, switchMap, throwError } from 'rxjs';
import { AuthService } from '../auth/auth';

export const authInterceptor: HttpInterceptorFn = (req, next) => {
  const auth   = inject(AuthService);
  const router = inject(Router);

  const authReq = req.clone({ withCredentials: true });

  return next(authReq).pipe(
    catchError((error: HttpErrorResponse) => {

      // 401 — tenta refresh uma vez, mas não se já for a rota de refresh
      if (error.status === 401 &&
        !req.url.includes("auth/me") &&
        !req.url.includes('/auth/refresh')
      ) {
        return from(auth.refreshAccessToken()).pipe(
          switchMap(() => next(authReq)),
          catchError(refreshError => {
            auth.logout();
            router.navigate(['/login']);
            return throwError(() => refreshError);
          }),
        );
      }

      if (error.status === 403) {
        router.navigate(['/forbidden']); // ← 403 é acesso negado
      }

      return throwError(() => error);
    }),
  );
};