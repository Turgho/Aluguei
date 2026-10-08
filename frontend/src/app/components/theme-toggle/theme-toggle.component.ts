import { Component, inject } from '@angular/core';
import { IconComponent } from '../icon/icon.component';
import { ThemeService } from '../../core/theme/theme.service';

@Component({
  selector: 'app-theme-toggle',
  standalone: true,
  imports: [IconComponent],
  template: `
    <button (click)="theme.toggle()" class="rounded-lg border border-border bg-bg-primary p-2 text-text-secondary hover:text-text-primary transition-colors" aria-label="Alternar tema">
      @if (theme.theme() === 'light') { <app-icon name="moon" /> } @else { <app-icon name="sun" /> }
    </button>
  `,
})
export class ThemeToggleComponent { theme = inject(ThemeService); }