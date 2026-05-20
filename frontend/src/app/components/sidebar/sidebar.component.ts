import { Component, Input } from '@angular/core';
import { RouterLink, RouterLinkActive } from '@angular/router';
import { IconComponent } from '../icon/icon.component';
import { IconName } from '../../shared/icons';

@Component({
  selector: 'app-sidebar',
  standalone: true,
  imports: [RouterLink, RouterLinkActive, IconComponent],
  template: `
    <aside
      class="fixed inset-y-0 left-0 z-40 flex h-screen w-64 flex-col border-r border-border bg-bg-primary transition-transform duration-300 lg:sticky lg:top-0 lg:translate-x-0"
      [class.-translate-x-full]="!open">
      
      <!-- Header fixo no topo -->
      <div class="flex h-16 shrink-0 items-center gap-2 border-b border-border px-5">
        <div class="flex h-9 w-9 items-center justify-center rounded-lg bg-brand-gradient text-white font-bold">A</div>
        <span class="text-base font-bold text-text-primary">Aluga<span class="text-brand-500">Pro</span></span>
      </div>

      <!-- Navegação ocupa todo o espaço disponível -->
      <nav class="flex-1 overflow-y-auto px-3 py-4">
        @for (item of items; track item.label) {
          <a
            [routerLink]="item.route"
            routerLinkActive="bg-brand-100 text-brand-800 dark:bg-brand-900/40 dark:text-brand-200"
            class="mb-1 flex items-center gap-3 rounded-lg px-3 py-2 text-sm font-medium text-text-secondary hover:bg-bg-secondary hover:text-text-primary transition-colors">
            <app-icon [name]="item.icon" />
            <span>{{ item.label }}</span>
          </a>
        }
      </nav>

      <!-- Footer fixo na base -->
      <div class="shrink-0 border-t border-border p-4">
        <div class="rounded-lg bg-brand-gradient dark:bg-brand-900/20 p-3">
          <p class="text-xs font-semibold text-brand-800 dark:text-brand-200">Plano Premium</p>
          <p class="mt-1 text-xs text-text-secondary">Relatórios ilimitados</p>
        </div>
      </div>
    </aside>
  `,
})
export class SidebarComponent {
  @Input() open = false;
  items: { label: string; route: string; icon: IconName }[] = [
    { label: 'Dashboard', route: '/dashboard', icon: 'dashboard' },
    { label: 'Propriedades', route: '/properties', icon: 'building' },
    { label: 'Inquilinos', route: '/tenants', icon: 'users' },
    { label: 'Contratos', route: '/contracts', icon: 'file' },
    { label: 'Relatórios', route: '/reports', icon: 'chart' },
  ];
}