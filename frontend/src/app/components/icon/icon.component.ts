import { Component, Input, inject } from '@angular/core';
import { DomSanitizer, SafeHtml } from '@angular/platform-browser';
import { ICONS, IconName } from '../../shared/icons';

@Component({
  selector: 'app-icon',
  standalone: true,
  template: `
    <svg
      [attr.viewBox]="'0 0 24 24'"
      [class]="svgClass + ' transition-colors'"
      [attr.fill]="brand ? null : 'none'"
      [attr.stroke]="brand ? null : 'currentColor'"
      [innerHTML]="safePath">
    </svg>
  `,
})
export class IconComponent {
  @Input({ required: true }) name!: IconName;
  @Input() size: 'sm' | 'md' | 'lg' = 'md';
  @Input() brand = false;

  private sanitizer = inject(DomSanitizer);

  get safePath(): SafeHtml {
    const raw = ICONS[this.name] ?? '';
    return this.sanitizer.bypassSecurityTrustHtml(raw);
  }

  get svgClass(): string {
    const base = 'shrink-0';
    return `${base} ${this.size === 'sm' ? 'h-4 w-4' : this.size === 'lg' ? 'h-6 w-6' : 'h-5 w-5'}`;
  }
}