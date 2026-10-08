import { AbstractControl, ValidationErrors, ValidatorFn } from "@angular/forms";

export function cpfValidator(): ValidatorFn {
  return (control: AbstractControl): ValidationErrors | null => {
    const raw = control.value?.replace(/\D/g, '') ?? '';
    if (!raw) return null;
    if (raw.length !== 11 || /^(\d)\1+$/.test(raw)) return { cpf: true };

    const calc = (len: number) => {
      const sum = raw
        .slice(0, len)
        .split('')
        .reduce((acc: number, d: string, i: number) => acc + +d * (len + 1 - i), 0);
      const rem = (sum * 10) % 11;
      return rem >= 10 ? 0 : rem;
    };

    return calc(9) === +raw[9] && calc(10) === +raw[10] ? null : { cpf: true };
  };
}

export function phoneValidator(): ValidatorFn {
  return (control: AbstractControl): ValidationErrors | null => {
    const raw = control.value?.replace(/\D/g, '') ?? '';
    if (!raw) return null;
    return raw.length < 10 || raw.length > 11 ? { phone: true } : null;
  };
}

export function passwordMatchValidator(): ValidatorFn {
  return (group: AbstractControl): ValidationErrors | null => {
    const pw = group.get('password')?.value;
    const confirm = group.get('confirmPassword')?.value;
    return pw && confirm && pw !== confirm ? { passwordMismatch: true } : null;
  };
}

export function passwordStrengthValidator(): ValidatorFn {
  return (control: AbstractControl): ValidationErrors | null => {
    const value: string = control.value ?? '';

    const rules = {
      minLength:   value.length >= 8,
      hasUpper:    /[A-Z]/.test(value),
      hasLower:    /[a-z]/.test(value),
      hasNumber:   /[0-9]/.test(value),
      hasSpecial:  /[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]/.test(value),
    };

    const passed = Object.values(rules).filter(Boolean).length;

    if (passed === 5) return null; // senha forte — sem erro

    return { passwordStrength: rules }; // devolve quais regras falharam
  };
}