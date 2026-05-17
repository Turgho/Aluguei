// src/app/features/dashboard/dashboard/dashboard.ts
import { Component, inject, OnInit, signal, computed } from '@angular/core';
import { BaseChartDirective } from 'ng2-charts';
import { ChartData, ChartOptions } from 'chart.js';
import { CommonModule } from '@angular/common';
import { Router } from '@angular/router';
import { AuthService } from '../../../core/auth/auth';

// ── Interfaces ────────────────────────────────────────────────

interface SummaryCard {
  label: string;
  value: string;
  iconPath: string; // path SVG do Heroicons
  bg: string;
}

interface MonthlyRevenue {
  label: string;
  value: number;  // altura em % para o gráfico (0-100)
  amount: string;
}

interface Property {
  name: string;
  address: string;
  status: 'Alugada' | 'Disponível';
  type: string;
  rent: string;
}

interface Payment {
  tenant: string;
  date: string;
  amount: string;
  method: string;
}

interface Alert {
  type: 'danger' | 'warning';
  title: string;
  description: string;
}

// ── Component ─────────────────────────────────────────────────

@Component({
  selector: 'app-dashboard',
  standalone: true,
  imports: [CommonModule, BaseChartDirective],
  templateUrl: './dashboard.html',
  styleUrl: './dashboard.scss',
})
export class DashboardComponent implements OnInit {
  private auth   = inject(AuthService);
  private router = inject(Router);

  // Primeiro nome do usuário autenticado via AuthService
  userName = computed(() => {
    const name = this.auth.user()?.name ?? '';
    return name.split(' ')[0] || 'Usuário';
  });

  // ── Cards de resumo ──────────────────────────────────────
  summaryCards: SummaryCard[] = [
    {
      label: 'Propriedades',
      value: '0',
      bg: 'bg-orange-50 border-orange-100 text-orange-700',
      // Ícone: casa (Heroicons - home)
      iconPath: 'M2.25 12l8.954-8.955a1.126 1.126 0 011.591 0L21.75 12M4.5 9.75v10.125c0 .621.504 1.125 1.125 1.125H9.75v-4.875c0-.621.504-1.125 1.125-1.125h2.25c.621 0 1.125.504 1.125 1.125V21h4.125c.621 0 1.125-.504 1.125-1.125V9.75M8.25 21h8.25',
    },
    {
      label: 'Alugadas',
      value: '0',
      bg: 'bg-emerald-50 border-emerald-100 text-emerald-700',
      // Ícone: check-circle (Heroicons)
      iconPath: 'M9 12.75L11.25 15 15 9.75M21 12a9 9 0 11-18 0 9 9 0 0118 0z',
    },
    {
      label: 'Receita Mensal',
      value: 'R$ 0',
      bg: 'bg-blue-50 border-blue-100 text-blue-700',
      // Ícone: banknotes (Heroicons)
      iconPath: 'M2.25 18.75a60.07 60.07 0 0115.797 2.101c.727.198 1.453-.342 1.453-1.096V18.75M3.75 4.5v.75A.75.75 0 013 6h-.75m0 0v-.375c0-.621.504-1.125 1.125-1.125H20.25M2.25 6v9m18-10.5v.75c0 .414.336.75.75.75h.75m-1.5-1.5h.375c.621 0 1.125.504 1.125 1.125v9.75c0 .621-.504 1.125-1.125 1.125h-.375m1.5-1.5H21a.75.75 0 00-.75.75v.75m0 0H3.75m0 0h-.375a1.125 1.125 0 01-1.125-1.125V15m1.5 1.5v-.75A.75.75 0 003 15h-.75',
    },
    {
      label: 'Pendentes',
      value: '0',
      bg: 'bg-yellow-50 border-yellow-100 text-yellow-700',
      // Ícone: clock (Heroicons)
      iconPath: 'M12 6v6h4.5m4.5 0a9 9 0 11-18 0 9 9 0 0118 0z',
    },
  ];

  barData: ChartData<'bar'> = {
    labels: ['Jul', 'Ago', 'Set', 'Out', 'Nov', 'Dez'],
    datasets: [{
      data: [4500, 6000, 3500, 8000, 5500, 10000],
      backgroundColor: '#f97316',
      borderRadius: 0
    }]
  };

  barOptions: ChartOptions<'bar'> = {
    responsive: true,
    maintainAspectRatio: false, // false para manter a altura
    plugins: { legend: { display: false } },
    scales: {
      y: { ticks: { callback: v => `R$ ${v}` } }
    }
  };

  // ── Receita dos últimos 6 meses ──────────────────────────
  // monthlyRevenue: MonthlyRevenue[] = [
  //   { label: 'Jul', value: 45,  amount: '4,5k' },
  //   { label: 'Ago', value: 60,  amount: '6k'   },
  //   { label: 'Set', value: 35,  amount: '3,5k' },
  //   { label: 'Out', value: 80,  amount: '8k'   },
  //   { label: 'Nov', value: 55,  amount: '5,5k' },
  //   { label: 'Dez', value: 100, amount: '10k'  },
  // ];

  // ── Propriedades ─────────────────────────────────────────
  properties: Property[] = [
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
  ];

  // ── Pagamentos recentes ───────────────────────────────────
  payments: Payment[] = [
    { tenant: 'João Silva',    date: '15/12/2024', amount: '2.800', method: 'PIX'           },
    { tenant: 'Maria Souza',   date: '14/12/2024', amount: '3.200', method: 'Transferência' },
    { tenant: 'Carlos Mendes', date: '10/12/2024', amount: '1.500', method: 'Boleto'        },
  ];

  // ── Alertas ──────────────────────────────────────────────
  alerts: Alert[] = [
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
  ];

  // ── Lifecycle ─────────────────────────────────────────────
  ngOnInit(): void {
    // Dados reais virão via AuthService.fetchMe() chamado no guard,
    // então auth.user() já estará populado ao entrar no dashboard.
  }

  // ── Ações ─────────────────────────────────────────────────
  async logout(): Promise<void> {
    await this.auth.logout();
    this.router.navigate(['/login']);
  }
}