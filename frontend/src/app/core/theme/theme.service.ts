import { Injectable, signal } from '@angular/core';

@Injectable({ providedIn: 'root' })
export class ThemeService {
  theme = signal<'light' | 'dark'>(
    (typeof localStorage !== 'undefined' && localStorage.getItem('theme') as any) || 'light'
  );
  toggle() { this.set(this.theme() === 'light' ? 'dark' : 'light'); }
  set(t: 'light' | 'dark') {
    this.theme.set(t);
    if (typeof document !== 'undefined') {
      document.documentElement.classList.toggle('dark', t === 'dark');
      localStorage.setItem('theme', t);
    }
  }
  init() { this.set(this.theme()); }
}