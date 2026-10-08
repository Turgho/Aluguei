import { Component, inject, signal, OnInit, computed } from '@angular/core';
import { Router, RouterLink } from '@angular/router';
import { FormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { toSignal } from '@angular/core/rxjs-interop';

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
import { UserRole } from '../../core/auth/models/user-role.model';
import { 
  cpfValidator, 
  passwordMatchValidator, 
  passwordStrengthValidator, 
  phoneValidator 
} from '../../core/utils/validators.util';
import { GRID_PATTERN, DOTS_PATTERN } from '../../shared/patterns/patterns';

// ─── Types ─────────────────────────────────────────────────────────────────────

export interface RegisterFormModel {
  firstName: string;
  lastName: string;
  email: string;
  cpf: string;
  phone: string;
  password: string;
  confirmPassword: string;
  terms: boolean;
}

type FormField = keyof RegisterFormModel;

interface PasswordRule {
  key: string;
  label: string;
}

// ─── Constants ─────────────────────────────────────────────────────────────────

const ERROR_MESSAGES: Record<string, Record<string, string>> = {
  firstName: {
    required: 'O nome é obrigatório',
    minlength: 'Mínimo 2 caracteres',
  },
  lastName: {
    required: 'O sobrenome é obrigatório',
    minlength: 'Mínimo 2 caracteres',
  },
  email: {
    required: 'O e-mail é obrigatório',
    email: 'Informe um e-mail válido',
  },
  cpf: {
    required: 'O CPF é obrigatório',
    cpf: 'CPF inválido',
  },
  phone: {
    phone: 'Telefone inválido',
  },
  password: {
    required: 'A senha é obrigatória',
    passwordStrength: 'A senha não atende aos requisitos',
  },
  confirmPassword: {
    required: 'Confirme sua senha',
    passwordMismatch: 'As senhas não conferem',
  },
  terms: {
    required: 'Aceite os termos para continuar',
  },
};

const PASSWORD_RULES: PasswordRule[] = [
  { key: 'minLength', label: 'Mínimo 8 caracteres' },
  { key: 'hasUpper', label: 'Letra maiúscula' },
  { key: 'hasLower', label: 'Letra minúscula' },
  { key: 'hasNumber', label: 'Número' },
  { key: 'hasSpecial', label: 'Caractere especial' },
];

const BENEFITS = [
  'Controle total de pagamentos',
  'Contratos digitais organizados',
  'Relatórios financeiros detalhados',
  'Notificações automáticas',
];

// ─── Component ─────────────────────────────────────────────────────────────────

@Component({
  selector: 'app-register-page',
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
    AnimateOnViewDirective,
    StaggerDirective,
    FooterComponent,
  ],
  templateUrl: './register.page.html',
})
export class RegisterPage implements OnInit {
  // Dependencies
  private readonly fb = inject(FormBuilder);
  private readonly router = inject(Router);
  private readonly theme = inject(ThemeService);
  private readonly auth = inject(AuthService);

  // Constants
  protected readonly gridPattern = GRID_PATTERN;
  protected readonly dotsPattern = DOTS_PATTERN;
  protected readonly benefits = BENEFITS;
  protected readonly passwordRules = PASSWORD_RULES;

  // State
  protected readonly isLoading = signal(false);
  protected readonly registerError = signal('');

  // Form
  protected readonly form = this.createForm();

  // Reactive password validation
  private readonly passwordValue = toSignal(
    this.form.controls.password.valueChanges,
    { initialValue: '' }
  );

  protected readonly passwordPassed = computed(() => {
    this.passwordValue();
    const errors = this.form.controls.password.errors?.['passwordStrength'];
    if (!errors) return PASSWORD_RULES.length;

    // errors[key] = true quando passou, false quando falhou
    return PASSWORD_RULES.filter(rule => errors[rule.key] === true).length;
  });

  // ─── Lifecycle ────────────────────────────────────────────────────────────

  ngOnInit(): void {
    this.theme.init();
  }

  // ─── Public Methods ───────────────────────────────────────────────────────

  /**
   * Retorna a cor da barra de progresso baseado na força da senha.
   */
  protected getStrengthColor(strength: number): string {
    if (strength <= 2) return 'bg-red-400';
    if (strength === 3) return 'bg-yellow-400';
    if (strength === 4) return 'bg-blue-400';
    return 'bg-emerald-400';
  }

  /**
   * Retorna a cor do texto do label de força.
   */
  protected getStrengthTextColor(strength: number): string {
    if (strength <= 2) return 'text-red-500';
    if (strength === 3) return 'text-yellow-500';
    if (strength === 4) return 'text-blue-500';
    return 'text-emerald-500';
  }

  /**
   * Retorna o label textual da força da senha.
   */
  protected getStrengthLabel(strength: number): string {
    if (strength <= 2) return 'Senha fraca';
    if (strength === 3) return 'Senha razoável';
    if (strength === 4) return 'Senha boa';
    return 'Senha forte';
  }

  /**
   * Retorna a primeira mensagem de erro de um campo.
   */
  protected getError(field: FormField): string {
    const control = this.form.get(field);
    if (!control?.touched) return '';

    // Erros no próprio campo
    if (control.errors) {
      const errorKey = Object.keys(control.errors)[0];
      return ERROR_MESSAGES[field]?.[errorKey] ?? '';
    }

    // Erro de grupo (passwordMismatch) exibido no confirmPassword
    if (field === 'confirmPassword' && control?.touched && this.form.errors?.['passwordMismatch']) {
      return ERROR_MESSAGES['confirmPassword']['passwordMismatch'];
    }

    return '';
  }

  /**
   * Submete o formulário de registro.
   */
  protected async onSubmit(): Promise<void> {
    if (this.form.invalid) {
      this.form.markAllAsTouched();
      return;
    }

    this.isLoading.set(true);
    this.registerError.set('');

    try {
      const { firstName, lastName, email, cpf, phone, password } = this.form.getRawValue();

      await this.auth.register({
        first_name: firstName,
        last_name: lastName,
        email: email,
        cpf: cpf.replace(/\D/g, ''), // tira a mask
        phone: phone || undefined,
        password: password,
        role: 'owner' as UserRole, // role padrão
      });

      setTimeout(async () => {
        await this.router.navigate(['/dashboard']);
      }, 1000);
    } catch (error) {
      this.registerError.set('Erro ao criar conta. Verifique os dados e tente novamente.');
    } finally {
      this.isLoading.set(false);
    }
  }

  // ─── Private Methods ──────────────────────────────────────────────────────

  /**
   * Cria o formulário de registro com validações.
   */
  private createForm() {
    return this.fb.nonNullable.group(
      {
        firstName:       ['', [Validators.required, Validators.minLength(2)]],
        lastName:        ['', [Validators.required, Validators.minLength(2)]],
        email:           ['', [Validators.required, Validators.email]],
        cpf:             ['', [Validators.required, cpfValidator()]],
        phone:           ['', [phoneValidator()]],
        password:        ['', [Validators.required, passwordStrengthValidator()]],
        confirmPassword: ['', [Validators.required]],
        terms:           [false, [Validators.requiredTrue]],
      },
      { validators: passwordMatchValidator() },
    );
  }
}