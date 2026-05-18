import { HttpInterceptorFn } from '@angular/common/http';
import { inject } from '@angular/core';
import { finalize, delay } from 'rxjs';
import { environment } from '../../../environments/environment';
import { AUTH_ROUTE_PATTERN } from '../constants/auth.constants';
import { LoadingService } from '../loading/loading.service';

export const loadingInterceptor: HttpInterceptorFn = (req, next) => {
  const skipLoading = AUTH_ROUTE_PATTERN.test(req.url);

  // Se a rota é de auth, não exibe o loading
  if (skipLoading) {
    return next(req);
  }

  // Exibe o loading
  const loading = inject(LoadingService);
  loading.show();

  const stream = next(req);

  // Se não é produção e está simulando rede lenta, adiciona delay
  if (!environment.production && environment.simulateSlowNetwork) {
    return stream.pipe(
      delay(environment.apiDelay),
      finalize(() => loading.hide()),
    );
  }

  return stream.pipe(finalize(() => loading.hide()));
};
