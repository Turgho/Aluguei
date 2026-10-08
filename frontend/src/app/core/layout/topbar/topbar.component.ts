import { Component, inject, Output, EventEmitter, signal } from '@angular/core';
import { Router } from '@angular/router';
import { IconComponent } from '../../../components/icon/icon.component';
import { ThemeToggleComponent } from '../../../components/theme-toggle/theme-toggle.component';
import { ButtonComponent } from '../../../components/button/button.component';
import { AuthService } from '../../auth/auth.service';

@Component({
  selector: 'app-topbar',
  standalone: true,
  imports: [IconComponent, ThemeToggleComponent, ButtonComponent],
  template: `
    <header class="sticky top-0 z-30 flex h-16 items-center justify-between border-b border-border bg-bg-primary/80 px-4 backdrop-blur-md lg:px-6">
      
      <!-- ═══ Esquerda: Menu Mobile + Busca ═══ -->
      <div class="flex items-center gap-3">
        <!-- Botão menu (mobile) -->
        <button 
          (click)="toggleMenu.emit()" 
          class="rounded-lg p-2 text-text-secondary hover:bg-bg-secondary hover:text-text-primary transition-colors lg:hidden" 
          aria-label="Abrir menu">
          <app-icon name="menu" />
        </button>

        <!-- Busca (desktop) -->
        <div class="relative hidden md:block">
          <span class="absolute left-3 top-1/2 -translate-y-1/2 text-text-muted pointer-events-none">
            <app-icon name="search" size="sm" />
          </span>
          <input 
            type="text" 
            placeholder="Buscar propriedades, inquilinos..." 
            class="w-64 lg:w-80 rounded-lg border border-border bg-bg-secondary py-2 pl-9 pr-3 text-sm text-text-primary placeholder:text-text-muted transition-all focus:outline-none focus:ring-2 focus:ring-brand-500/20 focus:border-brand-500 focus:bg-bg-primary" 
            aria-label="Buscar no sistema" />
        </div>

        <!-- Busca (mobile - ícone) -->
        <button 
          class="rounded-lg p-2 text-text-secondary hover:bg-bg-secondary hover:text-text-primary transition-colors md:hidden"
          aria-label="Buscar">
          <app-icon name="search" />
        </button>
      </div>

      <!-- ═══ Direita: Ações + Usuário ═══ -->
      <div class="flex items-center gap-1 sm:gap-2">
        <!-- Theme toggle -->
        <app-theme-toggle />

        <!-- Notificações -->
        <button 
          class="relative rounded-lg p-2 text-text-secondary hover:bg-bg-secondary hover:text-text-primary transition-colors"
          aria-label="Notificações - 3 não lidas">
          <app-icon name="bell" />
          <span class="absolute top-1.5 right-1.5 flex h-2 w-2">
            <span class="absolute inline-flex h-full w-full animate-ping rounded-full bg-red-400 opacity-75"></span>
            <span class="relative inline-flex h-2 w-2 rounded-full bg-red-500"></span>
          </span>
        </button>

        <!-- Separador -->
        <div class="hidden sm:block h-8 w-px bg-border mx-1" aria-hidden="true"></div>

        <!-- Perfil do usuário -->
        <button 
          class="flex items-center gap-2.5 rounded-lg p-1.5 hover:bg-bg-secondary transition-colors"
          aria-label="Menu do perfil">
          <div class="flex h-8 w-8 items-center justify-center rounded-full bg-brand-gradient text-white text-xs font-bold shrink-0">
            JP
          </div>
          <div class="hidden lg:block text-left">
            <p class="text-sm font-medium text-text-primary leading-tight">João Pereira</p>
            <p class="text-xs text-text-muted leading-tight">Proprietário</p>
          </div>
        </button>

        <!-- Logout -->
        <app-button 
          variant="ghost" 
          icon="logout" 
          size="sm"
          [loading]="isLoggingOut()"
          [disabled]="isLoggingOut()"
          [attr.aria-label]="isLoggingOut() ? 'Saindo...' : 'Sair da conta'"
          (click)="logout()" />
      </div>
    </header>
  `,
})
export class TopbarComponent {
  @Output() toggleMenu = new EventEmitter<void>();

  private readonly router = inject(Router);
  private readonly auth = inject(AuthService);

  protected readonly isLoggingOut = signal(false);

  /**
   * Realiza logout do usuário e redireciona para a tela de login.
   */
  async logout(): Promise<void> {
    if (this.isLoggingOut()) return;

    this.isLoggingOut.set(true);

    try {
      await this.auth.logout();
    } catch (error) {
      console.error('Erro ao fazer logout:', error);
    } finally {
      await this.router.navigate(['/login']);
      this.isLoggingOut.set(false);
    }
  }
}