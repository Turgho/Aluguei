import { Component, Input, Output, EventEmitter } from '@angular/core';
import { IconComponent } from '../icon/icon.component';

@Component({
  selector: 'app-modal',
  standalone: true,
  imports: [IconComponent],
  template: `
    @if (open) {
      <div class="fixed inset-0 z-50 flex items-center justify-center p-4 animate-fade-in">
        <div class="absolute inset-0 bg-black/50 backdrop-blur-sm" (click)="onClose()"></div>
        <div class="relative w-full max-w-lg rounded-xl border border-border bg-bg-primary shadow-xl">
          <div class="flex items-center justify-between border-b border-border px-5 py-4">
            <h3 class="text-base font-semibold text-text-primary">{{ title }}</h3>
            <button (click)="onClose()" class="rounded-md p-1 text-text-secondary hover:bg-bg-secondary" aria-label="Fechar">
              <app-icon name="close" />
            </button>
          </div>
          <div class="px-5 py-4"><ng-content /></div>
          @if (footer) {
            <div class="flex items-center justify-end gap-2 border-t border-border px-5 py-3">
              <ng-content select="[footer]" />
            </div>
          }
        </div>
      </div>
    }
  `,
})
export class ModalComponent {
  @Input() open = false;
  @Input() title = '';
  @Input() footer = true;
  @Output() closed = new EventEmitter<void>();
  onClose() { this.closed.emit(); }
}