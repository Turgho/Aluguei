import { Component, Input } from '@angular/core';
import { IconComponent } from '../../../components/icon/icon.component';
import { IconName } from '../../../shared/icons';

interface FooterLink {
  label: string;
  href: string;
  external?: boolean;
}

interface FooterSection {
  title: string;
  links: FooterLink[];
}

@Component({
  selector: 'app-footer',
  standalone: true,
  imports: [IconComponent],
  template: `
    @if (variant === 'slim') {
      <!-- ═══ VERSÃO SLIM (Dashboard) ═══ -->
      <footer class="mt-8 border-t border-border bg-bg-primary">
        <div class="px-4 py-5 sm:px-6 lg:px-8">
          <div class="flex flex-col gap-3 sm:flex-row sm:items-center sm:justify-between">
            <!-- Esquerda: Logo + Copyright -->
            <div class="flex flex-col gap-1 sm:flex-row sm:items-center sm:gap-3">
              <div class="flex items-center gap-2">
                <div class="flex h-7 w-7 items-center justify-center rounded-lg bg-brand-gradient text-white">
                  <app-icon name="building" size="sm" />
                </div>
                <span class="text-sm font-bold text-text-primary">
                  Alugu<span class="text-brand-500">Pro</span>
                </span>
              </div>
              <span class="text-xs text-text-secondary">
                © {{ currentYear }} Aluguei Tecnologia LTDA · CNPJ 12.345.678/0001-90
              </span>
            </div>

            <!-- Direita: Links legais -->
            <div class="flex flex-wrap items-center gap-x-4 gap-y-2 text-xs">
              <a href="#" class="text-text-secondary hover:text-text-primary transition-colors">Termos</a>
              <a href="#" class="text-text-secondary hover:text-text-primary transition-colors">Privacidade</a>
              <a href="#" class="text-text-secondary hover:text-text-primary transition-colors">Cookies</a>
              <a href="#" class="text-text-secondary hover:text-text-primary transition-colors">Suporte</a>
            </div>
          </div>

          <!-- Badges de segurança -->
          <div class="mt-3 flex flex-wrap items-center gap-3 border-t border-border pt-3">
            <div class="flex items-center gap-1.5 text-[10px] text-text-muted">
              <app-icon name="shieldCheck" size="sm" class="text-emerald-500" />
              <span>SSL 256-bit</span>
            </div>
            <div class="flex items-center gap-1.5 text-[10px] text-text-muted">
              <app-icon name="shieldCheck" size="sm" class="text-emerald-500" />
              <span>LGPD Compliant</span>
            </div>
            <div class="flex items-center gap-1.5 text-[10px] text-text-muted">
              <app-icon name="shieldCheck" size="sm" class="text-emerald-500" />
              <span>ISO 27001</span>
            </div>
            <div class="flex items-center gap-1.5 text-[10px] text-text-muted">
              <app-icon name="shieldCheck" size="sm" class="text-emerald-500" />
              <span>PCI DSS</span>
            </div>
          </div>
        </div>
      </footer>
    } @else {
      <!-- ═══ VERSÃO COMPLETA (Público) ═══ -->
      <footer class="bg-bg-primary border-t border-border">
        <!-- Newsletter -->
        <div class="border-b border-border bg-brand-gradient-soft dark:bg-brand-900/20">
          <div class="max-w-7xl mx-auto px-5 sm:px-6 lg:px-8 py-8 sm:py-10">
            <div class="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
              <div>
                <h3 class="text-lg font-bold text-text-primary mb-1">Fique por dentro das novidades</h3>
                <p class="text-sm text-text-secondary">Dicas, atualizações e conteúdo exclusivo sobre gestão de imóveis.</p>
              </div>
              <form class="flex w-full sm:w-auto gap-2" (ngSubmit)="onSubscribe($event)">
                <input
                  type="email"
                  placeholder="seu@email.com"
                  class="flex-1 sm:w-64 px-4 py-2.5 rounded-lg border border-border bg-bg-primary text-sm text-text-primary placeholder:text-text-secondary/50 focus:outline-none focus:ring-2 focus:ring-brand-500/20 focus:border-brand-500" />
                <button
                  type="submit"
                  class="px-5 py-2.5 bg-brand-600 hover:bg-brand-700 text-white text-sm font-semibold rounded-lg shadow-brand transition-all active:scale-[0.98]">
                  Assinar
                </button>
              </form>
            </div>
          </div>
        </div>

        <!-- Seções principais -->
        <div class="max-w-7xl mx-auto px-5 sm:px-6 lg:px-8 py-10 sm:py-12">
          <div class="grid gap-8 sm:grid-cols-2 lg:grid-cols-6">
            <!-- Coluna 1: Marca + Descrição -->
            <div class="lg:col-span-2">
              <div class="flex items-center gap-2 mb-4">
                <div class="flex h-10 w-10 items-center justify-center rounded-xl bg-brand-gradient text-white shadow-brand">
                  <app-icon name="building" />
                </div>
                <span class="text-xl font-bold text-text-primary">
                  Alugu<span class="text-brand-500">EI!</span>
                </span>
              </div>
              <p class="text-sm text-text-secondary leading-relaxed mb-4 max-w-xs">
                A plataforma mais completa para proprietários gerenciarem seus imóveis, contratos e inquilinos em um só lugar.
              </p>

              <!-- Contato -->
              <div class="space-y-2 text-sm">
                <a href="tel:+551140028922" class="flex items-center gap-2 text-text-secondary hover:text-text-primary transition-colors">
                  <app-icon name="phoneCall" size="sm" class="text-brand-500" />
                  <span>(11) 4002-8922</span>
                </a>
                <a href="mailto:contato@aluguei.com.br" class="flex items-center gap-2 text-text-secondary hover:text-text-primary transition-colors">
                  <app-icon name="mail" size="sm" class="text-brand-500" />
                  <span>contato@aluguei.com.br</span>
                </a>
                <div class="flex items-start gap-2 text-text-secondary">
                  <app-icon name="mapPin" size="sm" class="text-brand-500 mt-0.5 shrink-0" />
                  <span>Av. Paulista, 1000 - Bela Vista<br>São Paulo - SP, 01310-100</span>
                </div>
              </div>
            </div>

            <!-- Colunas de links -->
            @for (section of sections; track section.title) {
              <div>
                <h4 class="text-sm font-semibold text-text-primary mb-4">{{ section.title }}</h4>
                <ul class="space-y-2.5">
                  @for (link of section.links; track link.label) {
                    <li>
                      <a
                        [href]="link.href"
                        [target]="link.external ? '_blank' : '_self'"
                        [rel]="link.external ? 'noopener noreferrer' : ''"
                        class="text-sm text-text-secondary hover:text-brand-600 dark:hover:text-brand-400 transition-colors">
                        {{ link.label }}
                      </a>
                    </li>
                  }
                </ul>
              </div>
            }
          </div>
        </div>

        <!-- Bottom: Copyright + Social + Legal -->
        <div class="border-t border-border">
          <div class="max-w-7xl mx-auto px-5 sm:px-6 lg:px-8 py-6">
            <div class="flex flex-col gap-4 md:flex-row md:items-center md:justify-between">
              <div class="text-xs text-text-secondary text-center md:text-left">
                © {{ currentYear }} Aluguei Tecnologia LTDA · CNPJ 12.345.678/0001-90 · Todos os direitos reservados.
              </div>

              <!-- Redes sociais -->
              <div class="flex items-center justify-center gap-2">
                @for (social of socials; track social.name) {
                  <a
                    [href]="social.href"
                    target="_blank"
                    rel="noopener noreferrer"
                    [attr.aria-label]="social.name"
                    class="flex h-9 w-9 items-center justify-center rounded-lg border border-border bg-bg-primary text-text-secondary hover:text-brand-600 hover:border-brand-500 transition-all">
                    <app-icon [name]="social.icon" size="sm" />
                  </a>
                }
              </div>

              <!-- Badges -->
              <div class="flex items-center justify-center gap-3">
                <div class="flex items-center gap-1.5 text-xs text-text-muted">
                  <app-icon name="shieldCheck" size="sm" class="text-emerald-500" />
                  <span>LGPD</span>
                </div>
                <div class="flex items-center gap-1.5 text-xs text-text-muted">
                  <app-icon name="shieldCheck" size="sm" class="text-emerald-500" />
                  <span>SSL</span>
                </div>
                <div class="flex items-center gap-1.5 text-xs text-text-muted">
                  <app-icon name="shieldCheck" size="sm" class="text-emerald-500" />
                  <span>PCI DSS</span>
                </div>
              </div>
            </div>

            <!-- Feito com amor -->
            <div class="mt-4 text-center text-xs text-text-muted">
              Feito com
              <svg
                class="inline h-3.5 w-3.5 text-red-500 mx-0.5 -mt-0.5"
                viewBox="0 0 24 24"
                fill="currentColor"
                aria-hidden="true">
                <path d="M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 0 0 0-7.78z"/>
              </svg>
              em São Paulo, Brasil
            </div>
          </div>
          
        </div>
      </footer>
    }
  `,
})
export class FooterComponent {
  @Input() variant: 'slim' | 'full' = 'full';

  currentYear = new Date().getFullYear();
  buildNumber = 'a7f3k9';

  sections: FooterSection[] = [
    {
      title: 'Produto',
      links: [
        { label: 'Recursos', href: '#' },
        { label: 'Preços', href: '#' },
        { label: 'Integrações', href: '#' },
        { label: 'API', href: '#' },
        { label: 'Changelog', href: '#' },
      ],
    },
    {
      title: 'Empresa',
      links: [
        { label: 'Sobre nós', href: '#' },
        { label: 'Carreiras', href: '#' },
        { label: 'Blog', href: '#' },
        { label: 'Imprensa', href: '#' },
        { label: 'Parceiros', href: '#' },
      ],
    },
    {
      title: 'Suporte',
      links: [
        { label: 'Central de ajuda', href: '#' },
        { label: 'Contato', href: '#' },
        { label: 'Status do sistema', href: '#' },
        { label: 'Documentação', href: '#' },
        { label: 'Comunidade', href: '#' },
      ],
    },
    {
      title: 'Legal',
      links: [
        { label: 'Termos de uso', href: '#' },
        { label: 'Privacidade', href: '#' },
        { label: 'Cookies', href: '#' },
        { label: 'LGPD', href: '#' },
        { label: 'Segurança', href: '#' },
      ],
    },
  ];

  socials: { name: string; icon: IconName; href: string }[] = [
    { name: 'Instagram', icon: 'instagram', href: 'https://instagram.com/aluguei' },
    { name: 'Twitter', icon: 'twitter', href: 'https://twitter.com/aluguei' },
    { name: 'LinkedIn', icon: 'linkedin', href: 'https://linkedin.com/company/aluguei' },
    { name: 'YouTube', icon: 'youtube', href: 'https://youtube.com/@aluguei' },
    { name: 'GitHub', icon: 'github', href: 'https://github.com/aluguei' },
  ];

  onSubscribe(e: Event) {
    e.preventDefault();
    alert('Obrigado por se inscrever! 🎉');
  }
}