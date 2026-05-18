import { Injectable, signal, effect, computed } from '@angular/core';
import { getStorageItem, setStorageItem } from '../utils/storage.util';

export type Theme = 'light' | 'dark';

const THEME_KEY = 'theme';

@Injectable({ providedIn: 'root' })
export class ThemeService {
  private _theme = signal<Theme>(this.loadTheme());

  readonly theme = this._theme.asReadonly();
  readonly isDark = computed(() => this._theme() === 'dark');

  constructor() {
    // Aplica imediatamente — o effect roda depois do primeiro render
    this.applyToDocument(this._theme());

    effect(() => {
      this.applyToDocument(this._theme());
      setStorageItem(localStorage, THEME_KEY, this._theme());
    });
  }

  // Alterna o tema
  toggle(): void {
    this._theme.update(t => (t === 'light' ? 'dark' : 'light'));
  }

  // Define o tema
  set(theme: Theme): void {
    this._theme.set(theme);
  }

  // Aplica o tema ao documento
  private applyToDocument(theme: Theme): void {
    document.documentElement.classList.toggle('dark', theme === 'dark');
  }

  // Carrega o tema salvo no localStorage ou o tema do sistema
  private loadTheme(): Theme {
    const saved = getStorageItem(localStorage, THEME_KEY) as Theme | null;
    if (saved === 'light' || saved === 'dark') return saved;
    if (typeof window !== 'undefined' && window.matchMedia) {
      return window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light';
    }
    return 'light';
  }
}
