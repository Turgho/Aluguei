import { Component } from '@angular/core';
import { RouterLink } from '@angular/router';

@Component({
  selector: 'app-forbidden',
  standalone: true,
  imports: [RouterLink],
  template: `
    <main
      class="min-h-dvh flex flex-col items-center justify-center gap-4
             bg-bg-secondary text-text-primary px-6 text-center"
    >
      <p class="text-6xl font-extrabold text-brand-600">403</p>
      <h1 class="text-2xl font-bold">Acesso negado</h1>
      <p class="text-text-secondary max-w-md">
        Você não tem permissão para acessar esta página.
      </p>
      <a
        routerLink="/dashboard"
        class="mt-2 px-5 py-2.5 rounded-xl bg-brand-600 text-white font-semibold
               hover:bg-brand-700 transition-colors"
      >
        Voltar ao painel
      </a>
    </main>
  `,
})
export class ForbiddenComponent {}
