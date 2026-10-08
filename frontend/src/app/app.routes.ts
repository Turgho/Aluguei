import { Routes } from '@angular/router';
import { authGuard, guestGuard } from './core/guards/auth.guard';

export const routes: Routes = [
  { path: '', redirectTo: 'login', pathMatch: 'full' },

  // Rotas públicas
  {
    path: 'login',
    canActivate: [guestGuard],
    loadComponent: () =>
      import('./pages/login/login.component').then(m => m.LoginPage),
  },
  {
    path: 'register',
    canActivate: [guestGuard],
    loadComponent: () =>
      import('./pages/register/register.component').then(m => m.RegisterPage),
  },

  // Rotas protegidas — layout com sidebar/topbar
  {
    path: '',
    canActivate: [authGuard],
    loadComponent: () =>
      import('./core/layout/main-layout.component').then(m => m.MainLayoutComponent),
    children: [
      { path: 'dashboard',  loadComponent: () => import('./pages/dashboard/dashboard.component').then(m => m.DashboardPage)    },
      { path: 'properties', loadComponent: () => import('./pages/properties/properties.component').then(m => m.PropertiesPage) },
      { path: 'tenants',    loadComponent: () => import('./pages/tenants/tenants.component').then(m => m.TenantsPage)          },
      { path: 'contracts',  loadComponent: () => import('./pages/contracts/contracts.component').then(m => m.ContractsPage)    },
      { path: 'reports',    loadComponent: () => import('./pages/reports/reports.component').then(m => m.ReportsPage)          },
    ],
  },

  { path: '**', redirectTo: 'login' },
];