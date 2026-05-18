import { Component, inject } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { LoadingComponent } from './core/loading/loading.component';
import { ThemeService } from './core/theme/theme.service';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet, LoadingComponent],
  templateUrl: './app.html',
})
export class App {
  // Garante que o tema do localStorage seja aplicado em todas as rotas (ex.: login)
  private readonly _theme = inject(ThemeService);
}
