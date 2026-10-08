import { Component, inject, computed, signal, OnInit } from '@angular/core';
import { BaseChartDirective } from 'ng2-charts';
import { CommonModule } from '@angular/common';
import { Router } from '@angular/router';
import { AuthService } from '../../core/auth/auth.service';
import { ThemeService } from '../../core/theme/theme.service';
import { SkeletonComponent } from '../../shared/ui/skeleton/skeleton.component';
import { DashboardService } from './dashboard.service';
import { DashboardData } from './dashboard.models';
import { Icons } from '../../core/icons/icons';

@Component({
  selector: 'app-dashboard',
  standalone: true,
  imports: [CommonModule, BaseChartDirective, SkeletonComponent],
  templateUrl: './dashboard.component.html',
})
export class DashboardComponent implements OnInit {
  private router = inject(Router);
  private auth = inject(AuthService);
  private dashboardService = inject(DashboardService);

  protected theme = inject(ThemeService);
  protected data = signal<DashboardData | null>(null);
  protected pageLoading = signal(true);
  protected readonly Icons = Icons;

  // Nome do usuário
  protected userName = computed(() => {
    const name = this.auth.user()?.name ?? '';
    return name.split(' ')[0] || 'Usuário';
  });

  // Carrega os dados do dashboard
  ngOnInit(): void {
    this.dashboardService.load().subscribe(dashboard => {
      this.data.set(dashboard);
      this.pageLoading.set(false);
    });
  }

  // Logout do usuário
  async logout(): Promise<void> {
    await this.auth.logout();
    this.router.navigate(['/login']);
  }

  // Alterna o tema
  toggleTheme(): void {
    this.theme.toggle();
  }
}
