import { HttpInterceptorFn, HttpErrorResponse } from '@angular/common/http';
import { inject } from '@angular/core';
import { Router } from '@angular/router';
import { catchError, from, switchMap, throwError } from 'rxjs';
import { environment } from '../../../environments/environment';
import {
  AUTH_RETRY_HEADER,
  SKIP_AUTH_REFRESH_PATTERNS,
} from '../constants/auth.constants';
import { AuthService } from '../auth/auth.service';

function shouldAttemptRefresh(
  url: string,
  status: number,
  alreadyRetried: boolean,
): boolean {
  if (status !== 401 || alreadyRetried) return false;
  return !SKIP_AUTH_REFRESH_PATTERNS.some(pattern => pattern.test(url));
}

export const authInterceptor: HttpInterceptorFn = (req, next) => {
  const authReq = req.clone({ withCredentials: true });

  if (environment.useMockAuth) {
    return next(authReq);
  }

  const auth = inject(AuthService);
  const router = inject(Router);

  return next(authReq).pipe(
    catchError((error: HttpErrorResponse) => {
      const alreadyRetried = req.headers.has(AUTH_RETRY_HEADER);

      if (!shouldAttemptRefresh(req.url, error.status, alreadyRetried)) {
        if (error.status === 403) {
          router.navigate(['/forbidden']);
        }
        return throwError(() => error);
      }

      const retryReq = authReq.clone({
        setHeaders: { [AUTH_RETRY_HEADER]: 'true' },
      });

      return from(auth.refreshAccessToken()).pipe(
        switchMap(() => next(retryReq)),
        catchError(refreshError => {
          return from(auth.logout()).pipe(
            switchMap(() => {
              router.navigate(['/login']);
              return throwError(() => refreshError);
            }),
          );
        }),
      );
    }),
  );
};
