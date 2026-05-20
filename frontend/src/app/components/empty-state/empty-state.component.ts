import { Component, Input } from '@angular/core';
import { IconComponent } from '../icon/icon.component';
import { IconName } from '../../shared/icons';

@Component({
  selector: 'app-empty-state',
  standalone: true,
  imports: [IconComponent],
  template: `
    <div class="flex flex-col items-center justify-center rounded-xl border border-dashed border-border bg-bg-primary py-12 px-6 text-center">
      <div class="mb-4 flex h-14 w-14 items-center justify-center rounded-full bg-brand-100 text-brand-600 dark:bg-brand-900/40 dark:text-brand-300">
        <app-icon [name]="icon" size="lg" />
      </div>
      <h3 class="text-base font-semibold text-text-primary">{{ title }}</h3>
      <p class="mt-1 max-w-sm text-sm text-text-secondary">{{ description }}</p>
      <div class="mt-4"><ng-content /></div>
    </div>
  `,
})
export class EmptyStateComponent {
  @Input() icon: IconName = 'folder';
  @Input() title = 'Nenhum item encontrado';
  @Input() description = 'Ainda não há dados para exibir aqui.';
}