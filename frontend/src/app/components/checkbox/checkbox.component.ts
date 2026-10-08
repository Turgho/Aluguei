import { Component, Input, forwardRef } from '@angular/core';
import { ControlValueAccessor, NG_VALUE_ACCESSOR } from '@angular/forms';

@Component({
  selector: 'app-checkbox',
  standalone: true,
  providers: [{ provide: NG_VALUE_ACCESSOR, useExisting: forwardRef(() => CheckboxComponent), multi: true }],
  template: `
    <label class="flex items-center gap-3 cursor-pointer select-none group"
           [class.opacity-50]="disabled"
           [class.pointer-events-none]="disabled">

      <!-- Caixa customizada -->
      <div [class]="boxClasses">
        <input
          type="checkbox"
          class="sr-only"
          [checked]="value"
          [disabled]="disabled"
          (change)="onChangeEvent($event)" />

        <!-- Checkmark -->
        @if (value) {
          <svg class="w-3 h-3 text-white" viewBox="0 0 12 12" fill="none"
               stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
            <path d="M2 6l3 3 5-5"/>
          </svg>
        }
      </div>

      <!-- Label -->
      <span class="text-sm text-text-primary leading-none">
        <ng-content />
      </span>

    </label>
  `,
})
export class CheckboxComponent implements ControlValueAccessor {
  @Input() disabled = false;

  value = false;
  onChange: (v: boolean) => void = () => {};
  onTouched: () => void = () => {};

  writeValue(v: boolean): void { this.value = !!v; }
  registerOnChange(fn: any): void { this.onChange = fn; }
  registerOnTouched(fn: any): void { this.onTouched = fn; }

  onChangeEvent(e: Event): void {
    const v = (e.target as HTMLInputElement).checked;
    this.value = v;
    this.onChange(v);
    this.onTouched();
  }

  get boxClasses(): string {
    const base = 'relative flex h-5 w-5 shrink-0 items-center justify-center rounded-md border-2 transition-all duration-150';
    return this.value
      ? `${base} border-brand-600 bg-brand-600`
      : `${base} border-border bg-bg-primary group-hover:border-brand-400`;
  }
}