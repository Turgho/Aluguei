import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-badge',
  standalone: true,
  template: `<span [class]="classes"><ng-content /></span>`,
})
export class BadgeComponent {
  @Input() variant: 'success' | 'warning' | 'danger' | 'info' | 'neutral' | 'brand' = 'neutral';

  get classes(): string {
    const base = 'inline-flex items-center gap-1 rounded-full px-2.5 py-0.5 text-xs font-medium';
    const v = {
      success: 'bg-emerald-100 text-emerald-700 dark:bg-emerald-900/40 dark:text-emerald-300',
      warning: 'bg-amber-100 text-amber-700 dark:bg-amber-900/40 dark:text-amber-300',
      danger:  'bg-red-100 text-red-700 dark:bg-red-900/40 dark:text-red-300',
      info:    'bg-sky-100 text-sky-700 dark:bg-sky-900/40 dark:text-sky-300',
      neutral: 'bg-bg-secondary text-text-secondary border border-border',
      brand:   'bg-brand-100 text-brand-800 dark:bg-brand-900/40 dark:text-brand-200',
    };
    return `${base} ${v[this.variant]}`;
  }
}