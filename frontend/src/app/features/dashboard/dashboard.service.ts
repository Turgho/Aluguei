// src/app/features/dashboard/dashboard/dashboard.service.ts
import { Injectable } from '@angular/core';
import { delay, Observable, of } from 'rxjs';
import { environment } from '../../../environments/environment';
import { DashboardData } from './dashboard.models';
import { Icons } from '../../core/icons/icons';

const MOCK_DASHBOARD: DashboardData = {
  summaryCards: [
    {
      label: 'Propriedades',
      value: '12',
      bg: 'bg-orange-50 border-orange-100 text-orange-700',
      darkBg: 'dark:bg-orange-950/40 dark:border-orange-900/60 dark:text-orange-400',
      iconPath: Icons.home,
    },
    {
      label: 'Alugadas',
      value: '9',
      bg: 'bg-emerald-50 border-emerald-100 text-emerald-700',
      darkBg: 'dark:bg-emerald-950/40 dark:border-emerald-900/60 dark:text-emerald-400',
      iconPath: Icons.checkCircle,
    },
    {
      label: 'Receita Mensal',
      value: 'R$ 28.500',
      bg: 'bg-blue-50 border-blue-100 text-blue-700',
      darkBg: 'dark:bg-blue-950/40 dark:border-blue-900/60 dark:text-blue-400',
      iconPath: Icons.currencyDollar,
    },
    {
      label: 'Pendentes',
      value: '2',
      bg: 'bg-yellow-50 border-yellow-100 text-yellow-700',
      darkBg: 'dark:bg-yellow-950/40 dark:border-yellow-900/60 dark:text-yellow-400',
      iconPath: Icons.clock,
    },
  ],

  barData: {
    labels: ['Jul', 'Ago', 'Set', 'Out', 'Nov', 'Dez'],
    datasets: [
      {
        data: [4500, 6000, 3500, 8000, 5500, 10000],
        backgroundColor: '#f97316',
        borderRadius: 8,
      },
    ],
  },

  barOptions: {
    responsive: true,
    maintainAspectRatio: false,
    plugins: { legend: { display: false } },
    scales: {
      y: { ticks: { callback: v => `R$ ${v}` } },
    },
  },

  properties: [
    {
      name: 'Apto 302 - Jardins',
      address: 'Rua das Flores, 302 - São Paulo',
      status: 'Alugada',
      type: 'Apartamento',
      rent: '2.800',
    },
    {
      name: 'Casa Vila Madalena',
      address: 'Rua Harmonia, 150 - São Paulo',
      status: 'Disponível',
      type: 'Casa',
      rent: '4.500',
    },
    {
      name: 'Sala Comercial Centro',
      address: 'Av. Paulista, 1000 - São Paulo',
      status: 'Alugada',
      type: 'Comercial',
      rent: '3.200',
    },
  ],

  payments: [
    { tenant: 'João Silva',    date: '15/12/2024', amount: '2.800', method: 'PIX'           },
    { tenant: 'Maria Souza',   date: '14/12/2024', amount: '3.200', method: 'Transferência' },
    { tenant: 'Carlos Mendes', date: '10/12/2024', amount: '1.500', method: 'Boleto'        },
  ],

  alerts: [
    {
      type: 'danger',
      title: '2 contratos vencem em breve',
      description: 'Os contratos de João Silva e Maria Souza vencem em 30 dias.',
    },
    {
      type: 'warning',
      title: 'Pagamento atrasado',
      description: 'Carlos Mendes está com aluguel em atraso há 5 dias.',
    },
  ],
};

@Injectable({ providedIn: 'root' })
export class DashboardService {
  load(): Observable<DashboardData> {
    const mockDelay = environment.useMockDashboard ? environment.apiDelay : 0;
    return of(MOCK_DASHBOARD).pipe(delay(mockDelay));
  }
}