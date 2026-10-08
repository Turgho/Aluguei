import { Component, Input, forwardRef } from '@angular/core';
import { ControlValueAccessor, NG_VALUE_ACCESSOR } from '@angular/forms';

@Component({
  selector: 'app-textarea',
  standalone: true,
  providers: [{ provide: NG_VALUE_ACCESSOR, useExisting: forwardRef(() => TextareaComponent), multi: true }],
  template: `
    <div class="w-full">
      @if (label) {
        <label class="mb-1.5 block text-sm font-medium text-text-primary">{{ label }}</label>
      }
      <textarea
        rows="4"
        [placeholder]="placeholder"
        [disabled]="disabled"
        (input)="onInput($event)"
        (blur)="onTouched()"
        class="w-full rounded-lg border border-border bg-bg-primary px-3 py-2 text-sm text-text-primary placeholder:text-text-secondary transition-colors focus:outline-none focus:ring-2 focus:ring-brand-500 focus:border-transparent resize-y"></textarea>
      @if (error) { <p class="mt-1 text-xs text-red-500">{{ error }}</p> }
    </div>
  `,
})
export class TextareaComponent implements ControlValueAccessor {
  @Input() label?: string;
  @Input() placeholder = '';
  @Input() error?: string;
  @Input() disabled = false;

  value = '';
  onChange: (v: string) => void = () => {};
  onTouched: () => void = () => {};

  writeValue(v: string): void { this.value = v ?? ''; }
  registerOnChange(fn: any): void { this.onChange = fn; }
  registerOnTouched(fn: any): void { this.onTouched = fn; }

  onInput(e: Event) {
    const val = (e.target as HTMLTextAreaElement).value;
    this.value = val;
    this.onChange(val);
  }
}