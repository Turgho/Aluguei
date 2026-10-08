import { Component, Input, HostBinding } from '@angular/core';

@Component({
  selector: 'app-card',
  standalone: true,
  template: `<div [class]="classes"><ng-content /></div>`,
})
export class CardComponent {
  @Input() padding: 'none' | 'sm' | 'md' | 'lg' = 'md';
  @Input() hover = false;
  @Input() fullHeight = false;

  @HostBinding('class') get hostClass() {
    return this.fullHeight ? 'block h-full' : 'block';
  }

  get classes(): string {
    const base = 'rounded-xl border border-border bg-bg-primary shadow-sm';
    const pad = { none: '', sm: 'p-3', md: 'p-5', lg: 'p-6' }[this.padding];
    const h = this.hover ? 'transition-shadow hover:shadow-md' : '';
    const height = this.fullHeight ? 'h-full' : '';
    return `${base} ${pad} ${h} ${height}`;
  }
}