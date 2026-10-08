import { Component, Input, Output, EventEmitter, OnChanges, SimpleChanges, inject, HostListener } from '@angular/core';
import { DOCUMENT } from '@angular/common';
import { IconComponent } from '../icon/icon.component';

@Component({
  selector: 'app-modal',
  standalone: true,
  imports: [IconComponent],
  template: `
    @if (open) {
      <div
        class="fixed inset-0 z-50 flex items-end sm:items-center justify-center
               sm:p-4 animate-fade-in"
        role="dialog"
        aria-modal="true"
        [attr.aria-label]="title">

        <!-- Backdrop -->
        <div
          class="absolute inset-0 bg-black/50 backdrop-blur-sm"
          (click)="onClose()">
        </div>

        <!-- Modal -->
        <div class="relative w-full sm:max-w-lg flex flex-col
                    max-h-[92dvh] sm:max-h-[85dvh]
                    rounded-t-2xl sm:rounded-xl
                    border border-border bg-bg-primary shadow-xl">

          <!-- Header -->
          <div class="flex shrink-0 items-center justify-between
                      border-b border-border px-5 py-4">
            <h3 class="text-base font-semibold text-text-primary">{{ title }}</h3>
            <button
              type="button"
              (click)="onClose()"
              class="rounded-lg p-1.5 text-text-secondary
                     hover:bg-bg-secondary transition-colors"
              aria-label="Fechar">
              <app-icon name="close" size="sm" />
            </button>
          </div>

          <!-- Body com scroll -->
          <div class="flex-1 overflow-y-auto px-5 py-5">
            <ng-content />
          </div>

          <!-- Footer fixo — use <div footer> no template pai -->
          <ng-content select="[footer]" />

        </div>
      </div>
    }
  `,
})
export class ModalComponent implements OnChanges {
  @Input() open = false;
  @Input() title = '';
  @Output() closed = new EventEmitter<void>();

  private doc = inject(DOCUMENT);

  ngOnChanges(changes: SimpleChanges): void {
    if (changes['open']) {
      this.doc.body.style.overflow = this.open ? 'hidden' : '';
    }
  }

  // Fecha com Escape quando o modal está aberto
  @HostListener('document:keydown.escape')
  onEscapeKey(): void {
    if (this.open) this.onClose();
  }

  onClose(): void {
    this.closed.emit();
  }
}