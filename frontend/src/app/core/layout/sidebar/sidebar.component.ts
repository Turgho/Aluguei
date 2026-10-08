import { Component, Input } from '@angular/core';
import { RouterLink, RouterLinkActive } from '@angular/router';
import { IconComponent } from '../../../components/icon/icon.component';
import { IconName } from '../../../shared/icons';

interface SidebarItem {
  label: string;
  route: string;
  icon: IconName;
  badge?: number;
  badgeVariant?: 'danger' | 'warning' | 'info';
}

interface SidebarSection {
  title: string;
  items: SidebarItem[];
}

@Component({
  selector: 'app-sidebar',
  standalone: true,
  imports: [RouterLink, RouterLinkActive, IconComponent],
  template: `
    <aside
      class="fixed inset-y-0 left-0 z-40 flex h-screen w-64 flex-col border-r border-border bg-bg-primary transition-transform duration-300 lg:sticky lg:top-0 lg:translate-x-0 overflow-hidden"
      [class.-translate-x-full]="!open">
      
      <!-- ═══ Header com logo + status ═══ -->
      <div class="relative shrink-0">
        <!-- Sutil gradiente decorativo no topo -->
        <div class="absolute inset-x-0 top-0 h-32 bg-linar-to-b from-brand-500/5 to-transparent pointer-events-none" aria-hidden="true"></div>
        
        <div class="relative flex h-16 items-center justify-between px-5 border-b border-border">
          <div class="flex items-center gap-3">
            <!-- Logo elaborado -->
            <div class="relative">
              <div class="flex h-10 w-10 items-center justify-center rounded-xl bg-brand-gradient text-white shadow-brand shadow-brand-500/20">
                <app-icon name="building" size="md" />
              </div>
              <!-- Indicador online -->
              <span class="absolute -bottom-0.5 -right-0.5 flex h-3 w-3">
                <span class="absolute inline-flex h-full w-full animate-ping rounded-full bg-emerald-400 opacity-75"></span>
                <span class="relative inline-flex h-3 w-3 rounded-full bg-emerald-500 border-2 border-bg-primary"></span>
              </span>
            </div>
            
            <div>
              <span class="block text-base font-bold text-text-primary leading-tight">
                Alugu<span class="text-brand-500">EI!</span>
              </span>
              <span class="block text-[10px] text-text-muted font-medium uppercase tracking-wider">
                Gestão Imobiliária
              </span>
            </div>
          </div>
        </div>
      </div>

      <!-- ═══ Navegação agrupada por seções ═══ -->
      <nav class="flex-1 overflow-y-auto px-3 py-4 sidebar-scroll">
        @for (section of sections; track section.title; let isFirst = $first; let isLast = $last) {
          <!-- Título da seção -->
          <div class="mb-2 mt-4 first:mt-0 px-3">
            <p class="text-[10px] font-semibold text-text-muted uppercase tracking-wider">
              {{ section.title }}
            </p>
          </div>
          
          <!-- Itens da seção -->
          @for (item of section.items; track item.label) {
            <a
              [routerLink]="item.route"
              routerLinkActive="active-item"
              class="group relative mb-0.5 flex items-center gap-3 rounded-lg px-3 py-2.5 text-sm font-medium text-text-secondary hover:bg-bg-secondary hover:text-text-primary hover:translate-x-0.5 transition-all duration-200">
              
              <!-- Barra indicadora do item ativo -->
              <span class="absolute left-0 top-1/2 h-6 w-1 -translate-y-1/2 rounded-r-full bg-brand-500 scale-y-0 transition-transform duration-200"></span>
              
              <!-- Ícone com container -->
              <div class="relative flex h-8 w-8 shrink-0 items-center justify-center rounded-lg bg-bg-secondary group-hover:bg-brand-50 dark:group-hover:bg-brand-900/20 transition-colors">
                <app-icon [name]="item.icon" size="sm" />
                
                <!-- Badge de notificação -->
                @if (item.badge) {
                  <span class="absolute -top-1 -right-1 flex h-4 min-w-4 items-center justify-center rounded-full bg-red-500 px-1 text-[10px] font-bold text-white ring-2 ring-bg-primary">
                    {{ item.badge }}
                  </span>
                }
              </div>
              
              <span class="flex-1">{{ item.label }}</span>
              
              <!-- Seta sutil no hover -->
              <app-icon name="chevronRight" size="sm" class="opacity-0 group-hover:opacity-100 transition-opacity text-text-muted" />
            </a>
          }
          
          <!-- Separador entre seções (não aparece após a última) -->
          @if (!isLast) {
            <div class="my-3 mx-3 border-t border-border/60"></div>
          }
        }
      </nav>

      <!-- ═══ Footer: Card Premium + Info ═══ -->
      <div class="shrink-0 border-t border-border p-4 space-y-3">
        <!-- Card upgrade -->
        <div class="relative overflow-hidden rounded-xl bg-brand-gradient p-4 text-white">
          <!-- Padrão decorativo -->
          <div class="absolute inset-0 opacity-10" aria-hidden="true">
            <svg class="w-full h-full" viewBox="0 0 100 100" preserveAspectRatio="none">
              <defs>
                <pattern id="sidebar-pattern" width="20" height="20" patternUnits="userSpaceOnUse">
                  <circle cx="10" cy="10" r="1" fill="currentColor"/>
                </pattern>
              </defs>
              <rect width="100" height="100" fill="url(#sidebar-pattern)"/>
            </svg>
          </div>
          
          <div class="relative">
            <div class="flex items-center gap-2 mb-2">
              <div class="flex h-8 w-8 items-center justify-center rounded-lg bg-white/20 backdrop-blur">
                <app-icon name="shield" size="sm" />
              </div>
              <div>
                <p class="text-sm font-bold">Plano Premium</p>
                <p class="text-[10px] text-white/80">Expanda seu negócio</p>
              </div>
            </div>
            
            <p class="text-xs text-white/90 mb-3 leading-relaxed">
              Relatórios ilimitados, integrações avançadas e suporte prioritário.
            </p>
            
            <button class="w-full flex items-center justify-center gap-2 bg-white text-brand-700 text-xs font-semibold py-2 rounded-lg hover:bg-white/90 transition-colors active:scale-[0.98]">
              <app-icon name="arrowRight" size="sm" />
              <span>Fazer upgrade</span>
            </button>
          </div>
        </div>
        
        <!-- Info do sistema -->
        <div class="flex items-center justify-between px-2 py-1.5 text-[10px] text-text-muted">
          <div class="flex items-center gap-1.5">
            <span class="h-1.5 w-1.5 rounded-full bg-emerald-500"></span>
            <span>Sistema operacional</span>
          </div>
          <span class="font-mono">v2.4.1</span>
        </div>
      </div>
    </aside>
  `,
  styles: [`
    /* Apenas estilos que não dependem do tema */
    .sidebar-scroll::-webkit-scrollbar {
      width: 4px;
    }
    .sidebar-scroll::-webkit-scrollbar-track {
      background: transparent;
    }
    .sidebar-scroll::-webkit-scrollbar-thumb {
      background: var(--color-border);
      border-radius: 999px;
    }
    .sidebar-scroll::-webkit-scrollbar-thumb:hover {
      background: var(--color-text-muted);
    }
    .sidebar-scroll {
      scrollbar-width: thin;
      scrollbar-color: var(--color-border) transparent;
    }
  `]
})
export class SidebarComponent {
  @Input() open = false;
  
  sections: SidebarSection[] = [
    {
      title: 'Principal',
      items: [
        { label: 'Dashboard', route: '/dashboard', icon: 'dashboard' },
        { label: 'Propriedades', route: '/properties', icon: 'building' },
      ],
    },
    {
      title: 'Gestão',
      items: [
        { label: 'Inquilinos', route: '/tenants', icon: 'users' },
        { label: 'Contratos', route: '/contracts', icon: 'file', badge: 3, badgeVariant: 'warning' },
      ],
    },
    {
      title: 'Análises',
      items: [
        { label: 'Relatórios', route: '/reports', icon: 'chart' },
      ],
    },
  ];
}