import { Component, Input } from '@angular/core';
import { IconComponent } from '../icon/icon.component';
import { IconName } from '../../shared/icons';

@Component({
  selector: 'app-stat-card',
  standalone: true,
  imports: [IconComponent],
  template: `
    <div class="h-full rounded-xl border border-border bg-bg-primary p-4 sm:p-5 transition-shadow hover:shadow-md">
      <div class="flex items-start justify-between h-full">
        <div class="min-w-0 flex-1">
          <p class="text-xs sm:text-sm text-text-secondary truncate">{{ label }}</p>
          <p class="mt-1.5 sm:mt-2 text-xl sm:text-2xl font-bold text-text-primary truncate">{{ value }}</p>
          @if (delta) {
            <p class="mt-1 text-[10px] sm:text-xs" [class]="delta.startsWith('-') ? 'text-red-500' : 'text-emerald-500'">{{ delta }}</p>
          }
        </div>
        <div [class]="iconClasses">
          <app-icon [name]="icon" size="md" />
        </div>
      </div>
    </div>
  `,
})
export class StatCardComponent {
  @Input() label = '';
  @Input() value = '';
  @Input() delta?: string;
  @Input() icon: IconName = 'chart';
  @Input() color: 'brand' | 'emerald' | 'sky' | 'amber' | 'red' = 'brand';

  get iconClasses(): string {
    const c = {
      brand: 'bg-brand-100 text-brand-600 dark:bg-brand-900/40 dark:text-brand-300',
      emerald: 'bg-emerald-100 text-emerald-600 dark:bg-emerald-900/40 dark:text-emerald-300',
      sky: 'bg-sky-100 text-sky-600 dark:bg-sky-900/40 dark:text-sky-300',
      amber: 'bg-amber-100 text-amber-600 dark:bg-amber-900/40 dark:text-amber-300',
      red: 'bg-red-100 text-red-600 dark:bg-red-900/40 dark:text-red-300',
    }[this.color];
    return `flex h-10 w-10 sm:h-11 sm:w-11 items-center justify-center rounded-lg shrink-0 ${c}`;
  }
}