import { Component, inject, signal } from '@angular/core';
import { FormBuilder, ReactiveFormsModule, Validators, AbstractControl } from '@angular/forms';
import { Router, RouterLink } from '@angular/router';
import { AuthService } from '../../../core/auth/auth';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [ReactiveFormsModule, RouterLink],
  templateUrl: './login.html',
  styleUrl: './login.scss',
})
export class LoginComponent {
  private fb     = inject(FormBuilder);
  private router = inject(Router);
  private auth   = inject(AuthService);

  isLoading    = signal(false);
  showPassword = signal(false);
  loginError   = signal<string | null>(null);

  form = this.fb.group({
    email:      ['', [Validators.required, Validators.email]],
    password:   ['', [Validators.required, Validators.minLength(8)]],
    rememberMe: [false],
  });

  get email(): AbstractControl    { return this.form.get('email')!; }
  get password(): AbstractControl { return this.form.get('password')!; }

  togglePassword(): void {
    this.showPassword.update(v => !v);
  }

  async onSubmit(): Promise<void> {
    if (this.form.invalid) {
      this.form.markAllAsTouched();
      return;
    }

    this.isLoading.set(true);
    this.loginError.set(null);

    try {
      await this.auth.login({
        email:    this.form.value.email!,
        password: this.form.value.password!,
      });
      this.router.navigate(['/dashboard']);
    } catch {
      this.loginError.set('E-mail ou senha incorretos. Tente novamente.');
    } finally {
      this.isLoading.set(false);
    }
  }
}