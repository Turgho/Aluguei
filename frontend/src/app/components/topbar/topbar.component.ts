import { Component, inject, Output, EventEmitter } from '@angular/core';
import { Router } from '@angular/router';
import { IconComponent } from '../icon/icon.component';
import { ThemeToggleComponent } from '../theme-toggle/theme-toggle.component';
import { ButtonComponent } from '../button/button.component';

@Component({
  selector: 'app-topbar',
  standalone: true,
  imports: [IconComponent, ThemeToggleComponent, ButtonComponent],
  template: `
    <header class="sticky top-0 z-30 flex h-16 items-center justify-between border-b border-border bg-bg-primary/80 px-4 backdrop-blur lg:px-6">
      <div class="flex items-center gap-3">
        <button (click)="toggleMenu.emit()" class="rounded-lg p-2 text-text-secondary hover:bg-bg-secondary lg:hidden" aria-label="Abrir menu">
          <app-icon name="menu" />
        </button>
        <div class="relative hidden md:block">
          <span class="absolute left-3 top-1/2 -translate-y-1/2 text-text-secondary"><app-icon name="search" size="sm" /></span>
          <input type="text" placeholder="Buscar..." class="w-80 rounded-lg border border-border bg-bg-secondary py-2 pl-9 pr-3 text-sm text-text-primary placeholder:text-text-secondary focus:outline-none focus:ring-2 focus:ring-brand-500" />
        </div>
      </div>
      <div class="flex items-center gap-2">
        <app-theme-toggle />
        <button class="relative rounded-lg border border-border bg-bg-primary p-2 text-text-secondary hover:text-text-primary">
          <app-icon name="bell" />
          <span class="absolute top-1 right-1 h-2 w-2 rounded-full bg-red-500"></span>
        </button>
        <div class="ml-2 flex items-center gap-2">
          <div class="flex h-9 w-9 items-center justify-center rounded-full bg-brand-gradient text-white text-sm font-semibold">JP</div>
          <div class="hidden sm:block">
            <p class="text-sm font-medium text-text-primary leading-tight">João Pereira</p>
            <p class="text-xs text-text-secondary">Proprietário</p>
          </div>
        </div>
        <app-button variant="ghost" icon="logout" size="sm" (click)="logout()" />
      </div>
    </header>
  `,
})
export class TopbarComponent {
  @Output() toggleMenu = new EventEmitter<void>();
  router = inject(Router);
  logout() { this.router.navigate(['/login']); }
}