import { HttpInterceptorFn, HttpErrorResponse } from '@angular/common/http';
import { inject } from '@angular/core';
import { Router } from '@angular/router';
import { catchError, from, switchMap, throwError } from 'rxjs';
import { environment } from '../../../environments/environment';
import { AuthService } from '../auth/auth.service';

export const authInterceptor: HttpInterceptorFn = (req, next) => {
  if (environment.useMockAuth) {
    return next(req.clone({ withCredentials: true }));
  }

  const auth = inject(AuthService);
  const router = inject(Router);
  const authReq = req.clone({ withCredentials: true });

  // Intercepta a requisição
  return next(authReq).pipe(
    catchError((error: HttpErrorResponse) => {
      const isRefreshRoute = req.url.includes('/auth/refresh');
      const isMeRoute = req.url.includes('/auth/me');
      const isLoginRoute = req.url.includes('/auth/login');
      const isLogoutRoute = req.url.includes('/auth/logout');

      // Se o erro é 401 e não é uma rota de refresh, me ou login, logout, atualiza o token
      if (
        error.status === 401 &&
        !isRefreshRoute &&
        !isMeRoute &&
        !isLoginRoute &&
        !isLogoutRoute
      ) {
        return from(auth.refreshAccessToken()).pipe(
          switchMap(() => next(authReq)),
          catchError(refreshError => {
            void auth.logout();
            router.navigate(['/login']);
            return throwError(() => refreshError);
          }),
        );
      }

      // Se o erro é 403, redireciona para a página de acesso negado
      if (error.status === 403) {
        router.navigate(['/forbidden']);
      }

      return throwError(() => error);
    }),
  );
};
