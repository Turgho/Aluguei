import { Component, inject, signal } from '@angular/core';
import {
  AbstractControl,
  FormBuilder,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { Router, RouterLink } from '@angular/router';
import { AuthService } from '../../../core/auth/auth.service';
import { RegisterPayload } from '../../../core/auth/models/register.model';
import { USER_ROLE_OPTIONS, UserRole } from '../../../core/auth/models/user-role.model';
import {
  cpfValidator,
  formatCpf,
  formatPhone,
  onlyDigits,
} from './register.validators';

@Component({
  selector: 'app-register',
  standalone: true,
  imports: [ReactiveFormsModule, RouterLink],
  templateUrl: './register.component.html',
})
export class RegisterComponent {
  private fb = inject(FormBuilder);
  private router = inject(Router);
  private auth = inject(AuthService);

  readonly roleOptions = USER_ROLE_OPTIONS;

  isLoading = signal(false);
  showPassword = signal(false);
  registerError = signal<string | null>(null);
  registerSuccess = signal(false);

  form = this.fb.group({
    firstName: ['', [Validators.required, Validators.minLength(2)]],
    lastName: ['', [Validators.required, Validators.minLength(2)]],
    cpf: ['', [Validators.required, cpfValidator()]],
    email: ['', [Validators.required, Validators.email]],
    phone: [''],
    password: ['', [Validators.required, Validators.minLength(8)]],
    role: ['' as UserRole | '', Validators.required],
  });

  get firstName(): AbstractControl {
    return this.form.get('firstName')!;
  }

  get lastName(): AbstractControl {
    return this.form.get('lastName')!;
  }

  get cpf(): AbstractControl {
    return this.form.get('cpf')!;
  }

  get email(): AbstractControl {
    return this.form.get('email')!;
  }

  get phone(): AbstractControl {
    return this.form.get('phone')!;
  }

  get password(): AbstractControl {
    return this.form.get('password')!;
  }

  get role(): AbstractControl {
    return this.form.get('role')!;
  }

  // Alterna a visibilidade da senha
  togglePassword(): void {
    this.showPassword.update(v => !v);
  }

  // Formata o CPF
  onCpfInput(event: Event): void {
    const input = event.target as HTMLInputElement;
    const formatted = formatCpf(input.value);
    input.value = formatted;
    this.cpf.setValue(formatted, { emitEvent: false });
  }

  // Formata o telefone
  onPhoneInput(event: Event): void {
    const input = event.target as HTMLInputElement;
    const formatted = formatPhone(input.value);
    input.value = formatted;
    this.phone.setValue(formatted, { emitEvent: false });
  }

  // Submete o formulário de registro
  async onSubmit(): Promise<void> {
    if (this.form.invalid) {
      this.form.markAllAsTouched();
      return;
    }

    this.isLoading.set(true);
    this.registerError.set(null);
    this.registerSuccess.set(false);

    const payload: RegisterPayload = {
      first_name: this.firstName.value.trim(),
      last_name: this.lastName.value.trim(),
      cpf: onlyDigits(this.cpf.value),
      email: this.email.value.trim(),
      password: this.password.value,
      role: this.role.value as UserRole,
    };

    // Se o telefone tem dígitos, adiciona ao payload
    const phoneDigits = onlyDigits(this.phone.value);
    if (phoneDigits) {
      payload.phone = phoneDigits;
    }

    try {
      await this.auth.register(payload);
      this.registerSuccess.set(true);
      // Redireciona para o login
      setTimeout(() => this.router.navigate(['/login']), 1500);
    } catch {
      // Exibe mensagem de erro
      this.registerError.set(
        'Não foi possível criar a conta. Verifique os dados ou tente outro e-mail/CPF.',
      );
    } finally {
      // Oculta o loading
      this.isLoading.set(false);
    }
  }
}
