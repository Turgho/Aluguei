import { Component, signal, OnInit, inject, effect } from '@angular/core';
import { PageHeaderComponent } from '../../core/layout/page-header/page-header.component';
import { StatCardComponent } from '../../components/stat-card/stat-card.component';
import { CardComponent } from '../../components/card/card.component';
import { BadgeComponent } from '../../components/badge/badge.component';
import { ButtonComponent } from '../../components/button/button.component';
import { AlertComponent } from '../../components/alert/alert.component';
import { EmptyStateComponent } from '../../components/empty-state/empty-state.component';
import { SkeletonComponent } from '../../components/skeleton/skeleton.component';
import { BaseChartDirective } from 'ng2-charts';
import { ChartData, ChartOptions } from 'chart.js';
import { ThemeService } from '../../core/theme/theme.service';
import { StaggerDirective } from '../../shared/directives/stagger.directive';
import { AnimateOnViewDirective } from '../../shared/directives/animate-on-view.directive';
import { IconComponent } from '../../components/icon/icon.component';
import { RouterLink } from "@angular/router";

@Component({
  selector: 'app-dashboard-page',
  standalone: true,
  imports: [
    PageHeaderComponent,
    StatCardComponent,
    CardComponent,
    BadgeComponent,
    BaseChartDirective,
    ButtonComponent,
    AlertComponent,
    EmptyStateComponent,
    SkeletonComponent,
    AnimateOnViewDirective,
    StaggerDirective,
    IconComponent,
    RouterLink
],
  templateUrl: './dashboard.page.html',
})
export class DashboardPage implements OnInit {
  private theme = inject(ThemeService);

  loading = signal(true);

  barData: ChartData<'bar'> = {
    labels: ['Jul', 'Ago', 'Set', 'Out', 'Nov', 'Dez'],
    datasets: [
      {
        label: 'Receita',
        data: [12400, 14200, 15100, 14800, 16900, 18450],
        backgroundColor: '#f97316',
        hoverBackgroundColor: '#ea580c',
        borderRadius: 8,
        borderSkipped: false,
        maxBarThickness: 80,
      },
    ],
  };

  barOptions: ChartOptions<'bar'> = {
    responsive: true,
    maintainAspectRatio: false,
    plugins: {
      legend: { display: false },
      tooltip: {
        backgroundColor: 'rgba(15, 23, 42, 0.95)',
        padding: 12,
        cornerRadius: 8,
        titleColor: '#f1f5f9',
        bodyColor: '#f1f5f9',
        callbacks: {
          label: (ctx) => ` R$ ${ctx.parsed.y?.toLocaleString('pt-BR')}`,
        },
      },
    },
    scales: {
      x: {
        grid: { display: false },
        ticks: {
          color: '#6b7280',
          font: { size: 12, weight: 500 },
        },
        border: { display: false },
      },
      y: {
        grid: {
          color: 'rgba(229, 231, 235, 0.5)',
        },
        ticks: {
          color: '#6b7280',
          font: { size: 11 },
          callback: (v) => `R$ ${Number(v).toLocaleString('pt-BR')}`,
          maxTicksLimit: 5,
        },
        border: { display: false },
      },
    },
  };

  delinquents = [
    { name: 'Carlos Mendes', property: 'Apto 302 — Centro', amount: 'R$ 1.800', days: 15 },
    { name: 'Ana Ribeiro', property: 'Casa 14 — Jardim', amount: 'R$ 1.100', days: 8 },
    { name: 'Pedro Lima', property: 'Apto 201 — Sul', amount: 'R$ 300', days: 5 },
  ];

  properties = [
    { name: 'Apartamento 302', address: 'Rua das Flores, 120 — Centro', occupied: true },
    { name: 'Casa 14', address: 'Alameda dos Ipês, 45 — Jardim', occupied: true },
    { name: 'Studio 07', address: 'Av. Paulista, 1500 — Bela Vista', occupied: false },
  ];

  payments = [
    { name: 'Carlos Mendes', date: '05/09/2025', value: 'R$ 1.800', status: 'Atrasado' },
    { name: 'Ana Ribeiro', date: '10/09/2025', value: 'R$ 1.100', status: 'Pendente' },
    { name: 'Marcos Souza', date: '05/09/2025', value: 'R$ 2.200', status: 'Pago' },
    { name: 'Juliana Costa', date: '05/09/2025', value: 'R$ 1.500', status: 'Pago' },
  ];

  constructor() {
    // Atualiza cores do gráfico quando o tema muda
    effect(() => {
      const isDark = this.theme.theme() === 'dark';
      this.updateChartTheme(isDark);
    });
  }

  ngOnInit() {
    setTimeout(() => this.loading.set(false), 900);
  }

  private updateChartTheme(isDark: boolean): void {
    const textColor = isDark ? '#94a3b8' : '#6b7280';
    const gridColor = isDark ? 'rgba(51, 65, 85, 0.5)' : 'rgba(229, 231, 235, 0.5)';

    this.barOptions = {
      ...this.barOptions,
      scales: {
        x: {
          ...(this.barOptions.scales?.['x'] || {}),
          ticks: { color: textColor, font: { size: 12, weight: 500 } },
          grid: { display: false },
          border: { display: false },
        },
        y: {
          ...(this.barOptions.scales?.['y'] || {}),
          ticks: {
            color: textColor,
            font: { size: 11 },
            callback: (v) => `R$ ${Number(v).toLocaleString('pt-BR')}`,
            maxTicksLimit: 5,
          },
          grid: { color: gridColor },
          border: { display: false },
        },
      },
    };
  }
}