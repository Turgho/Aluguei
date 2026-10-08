import { Component, Input, forwardRef } from '@angular/core';
import { ControlValueAccessor, NG_VALUE_ACCESSOR } from '@angular/forms';

@Component({
  selector: 'app-select',
  standalone: true,
  providers: [{ provide: NG_VALUE_ACCESSOR, useExisting: forwardRef(() => SelectComponent), multi: true }],
  template: `
    <div class="w-full">
      @if (label) {
        <label class="my-1.5 block text-sm font-medium text-text-primary">{{ label }}</label>
      }
      <div class="relative">
        <select
          [disabled]="disabled"
          (change)="onInput($event)"
          (blur)="onTouched()"
          [class]="selectClasses">
          @if (placeholder) {
            <option value="" disabled [selected]="!value">{{ placeholder }}</option>
          }
          @for (opt of options; track opt.value) {
            <option [value]="opt.value" [selected]="opt.value === value">{{ opt.label }}</option>
          }
        </select>
        
        <!-- Seta customizada -->
        <div class="pointer-events-none absolute inset-y-0 right-0 flex items-center pr-3 text-text-secondary">
          <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true">
            <path d="m6 9 6 6 6-6"/>
          </svg>
        </div>
      </div>
      
      @if (error) { 
        <p class="mt-1.5 text-xs text-red-500 dark:text-red-400 animate-fade-in">{{ error }}</p> 
      }
    </div>
  `,
})
export class SelectComponent implements ControlValueAccessor {
  @Input() label?: string;
  @Input() placeholder?: string;
  @Input() options: { value: string | number; label: string }[] = [];
  @Input() error?: string;
  @Input() disabled = false;

  value: any = '';
  onChange: (v: any) => void = () => {};
  onTouched: () => void = () => {};

  writeValue(v: any): void { this.value = v ?? ''; }
  registerOnChange(fn: any): void { this.onChange = fn; }
  registerOnTouched(fn: any): void { this.onTouched = fn; }
  setDisabledState(d: boolean): void { this.disabled = d; }

  onInput(e: Event) {
    const val = (e.target as HTMLSelectElement).value;
    this.value = val;
    this.onChange(val);
  }

  get selectClasses(): string {
    const base = 'w-full appearance-none rounded-xl border bg-bg-primary pl-3 pr-10 py-3 text-sm text-text-primary transition-all focus:outline-none focus:ring-2 focus:ring-brand-500/20 focus:border-brand-500 cursor-pointer';
    const state = this.error 
      ? 'border-red-400 focus:border-red-500 focus:ring-red-500/20' 
      : 'border-border hover:border-brand-400';
    const disabled = this.disabled ? 'opacity-60 cursor-not-allowed' : '';
    return `${base} ${state} ${disabled}`.trim();
  }
}