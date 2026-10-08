import { Component, Input, HostBinding } from '@angular/core';
import { IconComponent } from '../icon/icon.component';
import { IconName } from '../../shared/icons';

@Component({
  selector: 'app-alert',
  standalone: true,
  imports: [IconComponent],
  template: `
    <div [class]="wrapClasses" role="alert">
      <app-icon [name]="iconName" class="shrink-0 mt-0.5" />
      <div class="text-sm"><ng-content /></div>
    </div>
  `,
})
export class AlertComponent {
  @Input() variant: 'info' | 'success' | 'warning' | 'danger' = 'info';
  
  @HostBinding('class') get hostClass() {
    return 'block h-full';
  }

  get iconName(): IconName {
    return this.variant === 'success' ? 'check' : this.variant === 'warning' || this.variant === 'danger' ? 'warning' : 'info';
  }

  get wrapClasses(): string {
    const base = 'h-full flex items-start gap-3 rounded-lg border p-3';
    const v = {
      info: 'border-sky-200 bg-sky-50 text-sky-800 dark:border-sky-900/50 dark:bg-sky-900/30 dark:text-sky-200',
      success: 'border-emerald-200 bg-emerald-50 text-emerald-800 dark:border-emerald-900/50 dark:bg-emerald-900/30 dark:text-emerald-200',
      warning: 'border-amber-200 bg-amber-50 text-amber-800 dark:border-amber-900/50 dark:bg-amber-900/30 dark:text-amber-200',
      danger: 'border-red-200 bg-red-50 text-red-800 dark:border-red-900/50 dark:bg-red-900/30 dark:text-red-200',
    }[this.variant];
    return `${base} ${v}`;
  }
}