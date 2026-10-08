import { Component, inject, signal, OnInit, computed } from '@angular/core';
import { Router, RouterLink } from '@angular/router';
import { FormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';

// Componentes
import { IconComponent } from '../../components/icon/icon.component';
import { ButtonComponent } from '../../components/button/button.component';
import { InputComponent } from '../../components/input/input.component';
import { CheckboxComponent } from '../../components/checkbox/checkbox.component';
import { ThemeToggleComponent } from '../../components/theme-toggle/theme-toggle.component';
import { BuildingIllustrationComponent } from '../../components/illustrations/building-illustration.component';
import { FooterComponent } from '../../core/layout/footer/footer.component';

// Diretivas e Pipes
import { AnimateOnViewDirective } from '../../shared/directives/animate-on-view.directive';
import { StaggerDirective } from '../../shared/directives/stagger.directive';
import { SafeHtmlPipe } from '../../shared/pipes/safe-html.pipe';

// Serviços e Utils
import { ThemeService } from '../../core/theme/theme.service';
import { AuthService } from '../../core/auth/auth.service';
import { GRID_PATTERN, DOTS_PATTERN } from '../../shared/patterns/patterns';

// ─── Types ─────────────────────────────────────────────────────────────────────

type LoginFormField = 'email' | 'password' | 'rememberMe';

interface Stat {
  value: string;
  label: string;
}

interface Avatar {
  initial: string;
  color: string;
}

// ─── Constants ─────────────────────────────────────────────────────────────────

const STATS: Stat[] = [
  { value: '1.2k+', label: 'Imóveis' },
  { value: '98%', label: 'Ocupação' },
  { value: '4.8★', label: 'Avaliação' },
];

const AVATARS: Avatar[] = [
  { initial: 'M', color: 'bg-brand-100 text-brand-700' },
  { initial: 'R', color: 'bg-brand-200 text-brand-800' },
  { initial: 'C', color: 'bg-brand-300 text-brand-900' },
];

const ERROR_MESSAGES = {
  email: {
    required: 'O e-mail é obrigatório',
    email: 'Informe um e-mail válido',
  },
  password: {
    required: 'A senha é obrigatória',
    minlength: (requiredLength: number) => `Mínimo ${requiredLength} caracteres`,
  },
} as const;

// ─── Component ─────────────────────────────────────────────────────────────────

@Component({
  selector: 'app-login-page',
  standalone: true,
  imports: [
    ReactiveFormsModule,
    RouterLink,
    IconComponent,
    ButtonComponent,
    InputComponent,
    CheckboxComponent,
    ThemeToggleComponent,
    SafeHtmlPipe,
    BuildingIllustrationComponent,
    StaggerDirective,
    AnimateOnViewDirective,
    FooterComponent,
  ],
  templateUrl: './login.page.html',
})
export class LoginPage implements OnInit {
  // Dependencies
  private readonly fb = inject(FormBuilder);
  private readonly router = inject(Router);
  private readonly theme = inject(ThemeService);
  private readonly auth = inject(AuthService);

  // Constants
  protected readonly gridPattern = GRID_PATTERN;
  protected readonly dotsPattern = DOTS_PATTERN;
  protected readonly stats = STATS;
  protected readonly avatars = AVATARS;

  // State
  protected readonly isLoading = signal(false);
  protected readonly loginError = signal('');

  // Form
  protected readonly form = this.createForm();

  // Computed errors
  protected readonly emailError = computed(() => this.getFieldError('email'));
  protected readonly passwordError = computed(() => this.getFieldError('password'));

  // ─── Lifecycle ────────────────────────────────────────────────────────────

  ngOnInit(): void {
    this.theme.init();
  }

  // ─── Public Methods ───────────────────────────────────────────────────────

  /**
   * Submete o formulário de login.
   */
  protected async onSubmit(): Promise<void> {
    if (this.form.invalid) {
      this.form.markAllAsTouched();
      return;
    }

    this.isLoading.set(true);
    this.loginError.set('');

    try {
      const { email, password } = this.form.getRawValue();

      await this.auth.login({ email, password });
      await this.router.navigate(['/dashboard']);
    } catch (error) {
      console.error('Login error:', error);
      this.loginError.set('E-mail ou senha incorretos. Tente novamente.');
    } finally {
      this.isLoading.set(false);
    }
  }

  // ─── Private Methods ──────────────────────────────────────────────────────

  /**
   * Cria o formulário de login com validações.
   */
  private createForm() {
    return this.fb.nonNullable.group({
      email: ['', [Validators.required, Validators.email]],
      password: ['', [Validators.required, Validators.minLength(8)]],
      rememberMe: [false],
    });
  }

  /**
   * Retorna a mensagem de erro de um campo específico.
   */
  private getFieldError(field: LoginFormField): string {
    const control = this.form.get(field);
    if (!control?.touched || !control?.invalid) return '';

    if (field === 'email') {
      if (control.errors?.['required']) return ERROR_MESSAGES.email.required;
      if (control.errors?.['email']) return ERROR_MESSAGES.email.email;
    }

    if (field === 'password') {
      if (control.errors?.['required']) return ERROR_MESSAGES.password.required;
      if (control.errors?.['minlength']) {
        return ERROR_MESSAGES.password.minlength(control.errors['minlength'].requiredLength);
      }
    }

    return '';
  }
}