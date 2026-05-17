// src/app/app.config.ts
import { ApplicationConfig, provideZonelessChangeDetection, isDevMode } from '@angular/core';
import { provideRouter, withViewTransitions, withComponentInputBinding } from '@angular/router';
import { provideHttpClient, withFetch, withInterceptors } from '@angular/common/http';
import { provideServiceWorker } from '@angular/service-worker';

import { routes } from './app.routes';
import { authInterceptor } from './core/interceptors/auth-interceptor';
import { provideCharts, withDefaultRegisterables } from 'ng2-charts';

export const appConfig: ApplicationConfig = {
  providers: [
    // Detecção de mudanças via Signals (sem Zone.js)
    provideZonelessChangeDetection(),

    // Roteamento com animações e suporte a @Input() nas rotas
    provideRouter(routes, withViewTransitions(), withComponentInputBinding()),

    // HttpClient com Fetch API + interceptor de autenticação
    provideHttpClient(withFetch(), withInterceptors([authInterceptor])),

    // PWA — Service Worker (ativo só em produção)
    provideServiceWorker('ngsw-worker.js', {
      enabled: !isDevMode(),
      registrationStrategy: 'registerWhenStable:30000',
    }),
    provideCharts(withDefaultRegisterables()),
  ],
};
