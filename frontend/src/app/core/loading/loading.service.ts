import { Injectable, signal, computed } from '@angular/core';

@Injectable({ providedIn: 'root' })
export class LoadingService {
  private _count = signal(0);
  private _timer: ReturnType<typeof setTimeout> | null = null;
  private _visible = signal(false);

  // Delay mínimo antes de mostrar (evita flash em respostas rápidas)
  private readonly SHOW_DELAY = 150;
  // Tempo mínimo visível (evita flicker quando esconde muito rápido)
  private readonly MIN_VISIBLE_TIME = 300;
  private _shownAt: number | null = null;

  readonly isLoading = computed(() => this._visible());

  /** Exibe o loading global. Suporta múltiplas chamadas simultâneas. */
  show(): void {
    this._count.update(v => v + 1);

    // Se já está visível ou tem timer ativo, não faz nada
    if (this._visible() || this._timer) return;

    this._timer = setTimeout(() => {
      if (this._count() > 0) {
        this._visible.set(true);
        this._shownAt = Date.now();
      }
      this._timer = null;
    }, this.SHOW_DELAY);
  }

  /** Oculta o loading. Só esconde quando todas as requisições terminarem. */
  hide(): void {
    this._count.update(v => Math.max(0, v - 1));

    if (this._count() > 0) return;

    // Cancela timer pendente
    if (this._timer) {
      clearTimeout(this._timer);
      this._timer = null;
    }

    // Garante tempo mínimo visível para evitar flicker
    if (this._visible() && this._shownAt) {
      const elapsed = Date.now() - this._shownAt;
      const remaining = this.MIN_VISIBLE_TIME - elapsed;

      if (remaining > 0) {
        setTimeout(() => {
          this._visible.set(false);
          this._shownAt = null;
        }, remaining);
      } else {
        this._visible.set(false);
        this._shownAt = null;
      }
    } else {
      this._visible.set(false);
      this._shownAt = null;
    }
  }

  /** Reseta completamente o estado (use com cautela). */
  reset(): void {
    this._count.set(0);
    this._visible.set(false);
    this._shownAt = null;
    if (this._timer) {
      clearTimeout(this._timer);
      this._timer = null;
    }
  }
}