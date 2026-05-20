import { Component, Input, forwardRef } from '@angular/core';
import { ControlValueAccessor, NG_VALUE_ACCESSOR } from '@angular/forms';

@Component({
  selector: 'app-checkbox',
  standalone: true,
  providers: [{ provide: NG_VALUE_ACCESSOR, useExisting: forwardRef(() => CheckboxComponent), multi: true }],
  template: `
    <label class="flex items-center gap-2.5 cursor-pointer select-none">
      <input
        type="checkbox"
        [checked]="value"
        [disabled]="disabled"
        (change)="onChangeEvent($event)"
        class="w-4 h-4 rounded accent-brand-600 cursor-pointer" />
      <span class="text-sm text-text-secondary">
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

  onChangeEvent(e: Event) {
    const v = (e.target as HTMLInputElement).checked;
    this.value = v;
    this.onChange(v);
    this.onTouched();
  }
}