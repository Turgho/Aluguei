import { Component, Input } from '@angular/core';
import { IconComponent } from '../icon/icon.component';
import { IconName } from '../../shared/icons';

type ButtonColor = 'default' | 'brand' | 'success' | 'warning' | 'danger' | 'info';

@Component({
  selector: 'app-button',
  standalone: true,
  imports: [IconComponent],
  template: `
    <button
      [type]="type"
      [disabled]="disabled || loading"
      [class]="classes"
      (click)="onClick($event)">
      @if (loading) {
        <span class="inline-block h-4 w-4 animate-spin rounded-full border-2 border-current border-r-transparent"></span>
      }
      @if (icon && !loading) {
        <app-icon [name]="icon" size="sm" [class]="iconColorClass" />
      }
      @if (label) { <span>{{ label }}</span> }
      <ng-content />
    </button>
  `,
})
export class ButtonComponent {
  @Input() label?: string;
  @Input() variant: 'primary' | 'secondary' | 'outline' | 'ghost' | 'danger' = 'primary';
  @Input() size: 'sm' | 'md' | 'lg' = 'md';
  @Input() icon?: IconName;
  @Input() type: 'button' | 'submit' | 'reset' = 'button';
  @Input() disabled = false;
  @Input() loading = false;
  @Input() full = false;
  @Input() color: ButtonColor = 'default';
  @Input() iconColor: ButtonColor = 'default';

  onClick(_: MouseEvent) {}

  // Paleta compartilhada para outline e iconColor
  private readonly palette: Record<Exclude<ButtonColor, 'default'>, string> = {
    brand:   'border-brand-500   text-brand-600   dark:text-brand-400   hover:bg-brand-50   dark:hover:bg-brand-900/20   focus-visible:ring-brand-500',
    success: 'border-emerald-400 text-emerald-600 dark:text-emerald-400 hover:bg-emerald-50 dark:hover:bg-emerald-900/20 focus-visible:ring-emerald-400',
    warning: 'border-amber-400   text-amber-600   dark:text-amber-400   hover:bg-amber-50   dark:hover:bg-amber-900/20   focus-visible:ring-amber-400',
    danger:  'border-red-400     text-red-600     dark:text-red-400     hover:bg-red-50     dark:hover:bg-red-900/20     focus-visible:ring-red-400',
    info:    'border-sky-400     text-sky-600     dark:text-sky-400     hover:bg-sky-50     dark:hover:bg-sky-900/20     focus-visible:ring-sky-400',
  };

  get iconColorClass(): string {
    if (this.iconColor === 'default') return '';
    const map: Record<Exclude<ButtonColor, 'default'>, string> = {
      brand:   'text-brand-500 dark:text-brand-400',
      success: 'text-emerald-500 dark:text-emerald-400',
      warning: 'text-amber-500 dark:text-amber-400',
      danger:  'text-red-500 dark:text-red-400',
      info:    'text-sky-500 dark:text-sky-400',
    };
    return map[this.iconColor] ?? '';
  }

  private get outlineClasses(): string {
    const base = 'bg-bg-primary active:scale-[0.98]';
    if (this.color === 'default') {
      return `${base} border border-border text-text-primary hover:bg-bg-secondary`;
    }
    return `${base} border ${this.palette[this.color]}`;
  }

  get classes(): string {
    const base = [
      'inline-flex cursor-pointer items-center justify-center gap-2',
      'rounded-xl font-semibold transition-all duration-150',
      'focus:outline-none focus-visible:ring-2 focus-visible:ring-offset-2',
      'disabled:opacity-60 disabled:cursor-not-allowed',
    ].join(' ');

    const sizes = {
      sm: 'px-4 py-2 text-sm',
      md: 'px-5 py-2.5 text-sm',
      lg: 'px-6 py-3.5 text-sm',
    };

    const variants = {
      primary:   'bg-brand-600 text-white shadow-brand hover:bg-brand-700 active:scale-[0.98]',
      secondary: 'bg-brand-100 text-brand-800 dark:bg-brand-900/40 dark:text-brand-200 hover:bg-brand-200 dark:hover:bg-brand-900/60',
      outline:   this.outlineClasses,
      ghost:     'text-text-secondary hover:bg-bg-secondary hover:text-text-primary',
      danger:    'bg-red-600 text-white hover:bg-red-700 active:scale-[0.98]',
    };

    return [base, sizes[this.size], variants[this.variant], this.full ? 'w-full' : ''].join(' ');
  }
}