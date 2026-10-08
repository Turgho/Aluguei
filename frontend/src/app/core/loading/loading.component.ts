import { Component, inject } from '@angular/core';
import { LoadingService } from './loading.service';

@Component({
  selector: 'app-loading',
  standalone: true,
  template: `
    @if (loading.isLoading()) {
      <div
        class="fixed inset-0 z-9999 flex items-center justify-center
               bg-black/40 backdrop-blur-sm animate-fade-in"
        role="status"
        aria-live="polite"s
        aria-label="Carregando conteúdo"
      >
        <!-- Container do spinner com sombra -->
        <div class="flex flex-col items-center gap-3">
          <!-- Spinner -->
          <div class="relative">
            <!-- Anel externo (decorativo) -->
            <div class="h-14 w-14 rounded-full border-4 border-brand-200/30 dark:border-brand-800/30"></div>
            <!-- Anel girando -->
            <div
              class="absolute inset-0 h-14 w-14 rounded-full border-4 border-transparent
                     border-t-brand-500 border-r-brand-500 animate-spin"
            ></div>
          </div>
          
          <!-- Texto opcional -->
          <p class="text-sm font-medium text-white drop-shadow-lg">
            Carregando...
          </p>
        </div>
      </div>
    }
  `,
})
export class LoadingComponent {
  protected readonly loading = inject(LoadingService);
}