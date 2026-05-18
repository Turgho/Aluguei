import { AbstractControl, ValidationErrors, ValidatorFn } from '@angular/forms';

/** Remove caracteres não numéricos */
export function onlyDigits(value: string): string {
  return (value ?? '').replace(/\D/g, '');
}

/** Valida CPF com 11 dígitos (formato; dígitos verificadores simplificados) */
export function cpfValidator(): ValidatorFn {
  return (control: AbstractControl): ValidationErrors | null => {
    const cpf = onlyDigits(control.value);
    if (!cpf) return null;
    if (cpf.length !== 11 || /^(\d)\1{10}$/.test(cpf)) {
      return { cpf: true };
    }
    return null;
  };
}

/** Máscara visual: 000.000.000-00 */
export function formatCpf(value: string): string {
  const d = onlyDigits(value).slice(0, 11);
  return d
    .replace(/(\d{3})(\d)/, '$1.$2')
    .replace(/(\d{3})(\d)/, '$1.$2')
    .replace(/(\d{3})(\d{1,2})$/, '$1-$2');
}

/** Máscara visual de telefone BR */
export function formatPhone(value: string): string {
  const d = onlyDigits(value).slice(0, 11);
  if (d.length <= 10) {
    return d
      .replace(/(\d{2})(\d)/, '($1) $2')
      .replace(/(\d{4})(\d)/, '$1-$2');
  }
  return d
    .replace(/(\d{2})(\d)/, '($1) $2')
    .replace(/(\d{5})(\d)/, '$1-$2');
}
