import { Injectable, signal, computed } from '@angular/core';

@Injectable({ providedIn: 'root' })
export class LoadingService {
  private _count = signal(0);
  private _timer: ReturnType<typeof setTimeout> | null = null;
  private _visible = signal(false);

  readonly isLoading = computed(() => this._visible());

  // Exibe o loading
  show(): void {
    this._count.update(v => v + 1);

    if (this._visible()) return;

    this._timer = setTimeout(() => {
      if (this._count() > 0) {
        this._visible.set(true);
      }
    }, 150);
  }

  // Oculta o loading
  hide(): void {
    this._count.update(v => Math.max(0, v - 1));

    if (this._count() === 0) {
      if (this._timer) clearTimeout(this._timer);
      this._timer = null;
      this._visible.set(false);
    }
  }

  // Reseta o loading
  reset(): void {
    this._count.set(0);
    this._visible.set(false);
    if (this._timer) clearTimeout(this._timer);
    this._timer = null;
  }
}
