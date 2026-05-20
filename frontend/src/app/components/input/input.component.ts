import { Component, Input, Output, EventEmitter, forwardRef } from '@angular/core';
import { ControlValueAccessor, NG_VALUE_ACCESSOR } from '@angular/forms';
import { IconComponent } from '../icon/icon.component';
import { IconName } from '../../shared/icons';

@Component({
  selector: 'app-input',
  standalone: true,
  imports: [IconComponent],
  providers: [{ provide: NG_VALUE_ACCESSOR, useExisting: forwardRef(() => InputComponent), multi: true }],
  template: `
    <div class="w-full">
      @if (label) {
        <label class="mb-2 block text-sm font-medium text-text-primary">{{ label }}</label>
      }
      <div class="relative">
        @if (icon) {
          <span class="absolute left-3.5 top-1/2 -translate-y-1/2 text-text-secondary pointer-events-none">
            <app-icon [name]="icon" size="sm" />
          </span>
        }
        <input
          [type]="type"
          [placeholder]="placeholder"
          [disabled]="disabled"
          [value]="value"
          (input)="onInput($event)"
          (blur)="onTouched()"
          [class]="inputClasses"
          [attr.aria-invalid]="!!error" />
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

  @Output() valueChange = new EventEmitter<string>();

  value: string = '';
  onChange: (v: string) => void = () => {};
  onTouched: () => void = () => {};

  writeValue(v: string): void { this.value = v ?? ''; }
  registerOnChange(fn: any): void { this.onChange = fn; }
  registerOnTouched(fn: any): void { this.onTouched = fn; }
  setDisabledState(d: boolean): void { this.disabled = d; }

  onInput(e: Event) {
    const val = (e.target as HTMLInputElement).value;
    this.value = val;
    this.onChange(val);
    this.valueChange.emit(val);
  }

  private get isDateInput(): boolean {
    return ['date', 'time', 'datetime-local', 'month', 'week'].includes(this.type);
  }

  get inputClasses(): string {
    const base = 'w-full rounded-xl border bg-bg-primary text-sm text-text-primary placeholder:text-text-secondary/50 transition-all focus:outline-none focus:ring-2 focus:ring-brand-500/20 focus:border-brand-500';
    
    // Padding esquerdo: com ícone = pl-10, sem ícone = pl-3.5
    const padLeft = this.icon ? 'pl-10' : 'pl-3.5';
    
    // Padding direito: inputs de data/time precisam de mais espaço para o ícone nativo
    const padRight = this.isDateInput ? 'pr-12' : 'pr-3.5';
    
    // Altura/padding vertical
    const padding = 'py-3';
    
    // Estado de erro ou normal
    const state = this.error
      ? 'border-red-400 focus:border-red-500 focus:ring-red-500/20'
      : 'border-border hover:border-brand-400';
    
    // Disabled
    const disabled = this.disabled ? 'opacity-60 cursor-not-allowed' : 'cursor-text';
    
    // Classe especial para inputs de data (remove UI nativa inconsistente)
    const dateFix = this.isDateInput ? 'date-input-fix' : '';
    
    return `${base} ${padLeft} ${padRight} ${padding} ${state} ${disabled} ${dateFix}`.trim();
  }
}