import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-skeleton',
  standalone: true,
  template: `<span [class]="classes"></span>`,
  styles: [`:host { display: block; } .sk { background: linear-gradient(90deg, var(--color-border) 25%, color-mix(in srgb, var(--color-border) 60%, transparent) 50%, var(--color-border) 75%); background-size: 200% 100%; animation: shimmer 1.4s infinite; border-radius: 0.5rem; } @keyframes shimmer { 0% { background-position: 200% 0; } 100% { background-position: -200% 0; } }`],
})
export class SkeletonComponent {
  @Input() width = '100%';
  @Input() height = '1rem';
  @Input() rounded: 'sm' | 'md' | 'lg' | 'full' = 'md';

  get classes(): string {
    const r = { sm: 'rounded', md: 'rounded-md', lg: 'rounded-lg', full: 'rounded-full' }[this.rounded];
    return `sk ${r} block`;
  }
}