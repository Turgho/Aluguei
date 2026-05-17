// src/app/app.routes.ts
import { Routes } from '@angular/router';
import { authGuard, guestGuard } from './core/guards/auth';

export const routes: Routes = [
  {
    path: '',
    redirectTo: 'login',
    pathMatch: 'full',
  },

  // Rotas públicas (bloqueadas se já estiver logado)
  {
    path: 'login',
    canActivate: [guestGuard],
    loadComponent: () =>
      import('./features/auth/login/login').then(m => m.LoginComponent),
  },
  {
    path: 'dashboard',
    canActivate: [authGuard],
    loadComponent: () =>
      import('./features/dashboard/dashboard/dashboard').then(m => m.DashboardComponent),
  },
  // Fallback
  {
    path: '**',
    redirectTo: 'login',
  },
];