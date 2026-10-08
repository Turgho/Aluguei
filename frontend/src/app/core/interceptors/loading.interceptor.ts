import { HttpInterceptorFn } from '@angular/common/http';
import { inject } from '@angular/core';
import { finalize, delay } from 'rxjs';
import { environment } from '../../../environments/environment';
import { LoadingService } from '../loading/loading.service';
import { SKIP_LOADING_PATTERNS } from '../constants/auth.constants';

export const loadingInterceptor: HttpInterceptorFn = (req, next) => {
  // Verifica se a rota deve pular o loading
  const shouldSkip = SKIP_LOADING_PATTERNS.some(pattern => pattern.test(req.url));

  if (shouldSkip) {
    return next(req);
  }

  const loading = inject(LoadingService);
  loading.show();

  const stream = next(req);

  // Simula rede lenta em desenvolvimento
  if (!environment.production && environment.simulateSlowNetwork) {
    return stream.pipe(
      delay(environment.apiDelay ?? 1000),
      finalize(() => loading.hide()),
    );
  }

  return stream.pipe(finalize(() => loading.hide()));
};