import { Component, Input, Output, EventEmitter, forwardRef, signal } from '@angular/core';
import { ControlValueAccessor, NG_VALUE_ACCESSOR } from '@angular/forms';
import { IconComponent } from '../icon/icon.component';
import { IconName } from '../../shared/icons';

type InputMask = 'cpf' | 'phone' | 'cep';

function applyMask(value: string, mask: InputMask): string {
  if (mask === 'cpf') {
    let v = value.replace(/\D/g, '').slice(0, 11);
    if (v.length > 9)      return v.replace(/(\d{3})(\d{3})(\d{3})(\d{0,2})/, '$1.$2.$3-$4');
    if (v.length > 6)      return v.replace(/(\d{3})(\d{3})(\d{0,3})/,        '$1.$2.$3');
    if (v.length > 3)      return v.replace(/(\d{3})(\d{0,3})/,               '$1.$2');
    return v;
  }

  if (mask === 'phone') {
    let v = value.replace(/\D/g, '').slice(0, 11);
    if (v.length > 10)     return v.replace(/(\d{2})(\d{5})(\d{0,4})/,   '($1) $2-$3');
    if (v.length > 6)      return v.replace(/(\d{2})(\d{4,5})(\d{0,4})/, '($1) $2-$3');
    if (v.length > 2)      return v.replace(/(\d{2})(\d{0,5})/,          '($1) $2');
    return value.replace(/(\d{0,2})/, '($1');
  }

  if (mask === 'cep'){
    let v = value.replace(/\D/g, '').slice(0, 8);
    if (v.length > 5) return v.replace(/(\d{5})(\d{0,3})/, '$1-$2');
    return v;
  }

  return value;
}

@Component({
  selector: 'app-input',
  standalone: true,
  imports: [IconComponent],
  providers: [{ provide: NG_VALUE_ACCESSOR, useExisting: forwardRef(() => InputComponent), multi: true }],
  template: `
    <div class="w-full">
      @if (label) {
        <label class="my-1.5 block text-sm font-medium text-text-primary">{{ label }}</label>
      }

      <div class="relative">
        @if (icon) {
          <span class="absolute left-3.5 top-1/2 -translate-y-1/2 text-text-secondary pointer-events-none">
            <app-icon [name]="icon" size="sm" />
          </span>
        }

        <input
          [type]="isPasswordType ? (passwordVisible() ? 'text' : 'password') : type"
          [placeholder]="placeholder"
          [disabled]="disabled"
          [value]="value"
          (input)="onInput($event)"
          (blur)="onTouched()"
          [class]="inputClasses"
          [attr.aria-invalid]="!!error" />

        @if (isPasswordType) {
          <button
            type="button"
            (click)="togglePasswordVisibility()"
            [attr.aria-label]="passwordVisible() ? 'Ocultar senha' : 'Mostrar senha'"
            class="absolute right-3 top-1/2 -translate-y-1/2 p-1
                   text-text-secondary hover:text-text-primary
                   rounded-lg transition-colors
                   focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-brand-500">
            <app-icon [name]="passwordVisible() ? 'eyeOff' : 'eye'" size="sm" />
          </button>
        }
      </div>

      @if (error) {
        <p class="mt-1.5 text-xs text-red-500 dark:text-red-400 animate-fade-in">{{ error }}</p>
      }
    </div>
  `,
})
export class InputComponent implements ControlValueAccessor {
  @Input() label?: string;
  @Input() placeholder = '';
  @Input() type: string = 'text';
  @Input() icon?: IconName;
  @Input() error?: string;
  @Input() disabled = false;
  @Input() mask?: InputMask;

  @Output() valueChange = new EventEmitter<string>();

  value: string = '';
  passwordVisible = signal(false);
  onChange: (v: string) => void = () => {};
  onTouched: () => void = () => {};

  writeValue(v: string): void { this.value = v ?? ''; }
  registerOnChange(fn: any): void { this.onChange = fn; }
  registerOnTouched(fn: any): void { this.onTouched = fn; }
  setDisabledState(d: boolean): void { this.disabled = d; }

  onInput(e: Event): void {
    const input = e.target as HTMLInputElement;
    const masked = this.mask ? applyMask(input.value, this.mask) : input.value;
    if (this.mask) input.value = masked;
    this.value = masked;
    this.onChange(masked);
    this.valueChange.emit(masked);
  }

  togglePasswordVisibility(): void {
    this.passwordVisible.update(v => !v);
  }

  get isPasswordType(): boolean {
    return this.type === 'password';
  }

  private get isDateInput(): boolean {
    return ['date', 'time', 'datetime-local', 'month', 'week'].includes(this.type);
  }

  get inputClasses(): string {
    const base     = 'w-full rounded-xl border bg-bg-primary text-sm text-text-primary placeholder:text-text-secondary/50 transition-all focus:outline-none focus:ring-2 focus:ring-brand-500/20 focus:border-brand-500';
    const padLeft  = this.icon ? 'pl-10' : 'pl-3.5';
    const padRight = this.isPasswordType ? 'pr-11' : 'pr-3.5';
    const padding  = 'py-3';
    const state    = this.error
      ? 'border-red-400 focus:border-red-500 focus:ring-red-500/20'
      : 'border-border hover:border-brand-400';
    const disabled = this.disabled ? 'opacity-60 cursor-not-allowed' : 'cursor-text';
    const dateFix  = '';

    return `${base} ${padLeft} ${padRight} ${padding} ${state} ${disabled} ${dateFix}`.trim();
  }
}