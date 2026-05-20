import { Routes } from '@angular/router';

export const routes: Routes = [
  { path: '', redirectTo: 'login', pathMatch: 'full' },
  { path: 'login', loadComponent: () => import('./pages/login/login.component').then(m => m.LoginPage) },
  { path: 'register', loadComponent: () => import('./pages/register/register.component').then(m => m.RegisterPage) },
  {
    path: '',
    loadComponent: () => import('./core/layout/main-layout.component').then(m => m.MainLayoutComponent),
    children: [
      { path: 'dashboard', loadComponent: () => import('./pages/dashboard/dashboard.component').then(m => m.DashboardPage) },
      { path: 'properties', loadComponent: () => import('./pages/properties/properties.component').then(m => m.PropertiesPage) },
      { path: 'tenants', loadComponent: () => import('./pages/tenants/tenants.component').then(m => m.TenantsPage) },
      { path: 'contracts', loadComponent: () => import('./pages/contracts/contracts.component').then(m => m.ContractsPage) },
      { path: 'reports', loadComponent: () => import('./pages/reports/reports.component').then(m => m.ReportsPage) },
    ],
  },
  { path: '**', redirectTo: 'login' },
];