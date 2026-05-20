import {
  Component,
  Input,
  OnInit,
  OnDestroy,
  ChangeDetectionStrategy,
  ChangeDetectorRef,
} from '@angular/core';
import { CommonModule } from '@angular/common';

/**
 * BuildingIllustrationComponent
 * Detecta o tema pelo atributo de classe no <html>
 * e troca sol ↔ lua com transição suave.
 *
 * Uso no template:
 *   <app-building-illustration />
 *
 * Forçar tema:
 *   <app-building-illustration [dark]="true" />
 *
 * Classe customizada:
 *   <app-building-illustration className="w-full max-w-sm mx-auto" />
 */
@Component({
  selector: 'app-building-illustration',
  standalone: true,
  imports: [CommonModule],
  changeDetection: ChangeDetectionStrategy.OnPush,
  template: `
    <svg
      viewBox="0 0 360 230"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
      [class]="className || 'w-full'"
      style="max-width: 380px; transition: all 0.4s ease"
    >
      <!-- ── Estrelas (visíveis só no dark) ── -->
      <circle
        *ngFor="let star of stars; let i = index"
        [attr.cx]="star[0]"
        [attr.cy]="star[1]"
        [attr.r]="i % 3 === 0 ? 1.5 : 1"
        fill="rgba(255,255,255,1)"
        [style.opacity]="isDark ? (i % 2 === 0 ? 0.55 : 0.35) : 0"
        [style.transition]="'opacity 0.6s ease ' + i * 30 + 'ms'"
      />

      <!-- ── Sol (light mode) ── -->
      <g [style.opacity]="isDark ? 0 : 1" style="transition: opacity 0.5s ease">
        <circle cx="310" cy="28" r="22" fill="rgba(255,220,80,0.12)" />
        <circle cx="310" cy="28" r="16" fill="rgba(255,220,80,0.20)" />
        <circle cx="310" cy="28" r="11" fill="rgba(255,220,80,0.85)" />
        <line
          *ngFor="let ray of sunRays"
          [attr.x1]="ray.x1" [attr.y1]="ray.y1"
          [attr.x2]="ray.x2" [attr.y2]="ray.y2"
          stroke="rgba(255,220,80,0.60)"
          stroke-width="1.5"
          stroke-linecap="round"
        />
      </g>

      <!-- ── Lua (dark mode) ── -->
      <g [style.opacity]="isDark ? 1 : 0" style="transition: opacity 0.5s ease">
        <circle cx="310" cy="28" r="20" fill="rgba(200,210,255,0.07)" />
        <circle cx="310" cy="28" r="14" fill="rgba(200,210,255,0.10)" />
        <circle cx="310" cy="28" r="11" fill="rgba(220,225,255,0.90)" />
        <circle cx="316" cy="24" r="9"
          [attr.fill]="isDark ? 'rgba(15,23,42,1)' : 'rgba(255,255,255,0)'"
          style="transition: fill 0.4s ease"
        />
        <circle cx="304" cy="30" r="1.5" fill="rgba(180,190,240,0.40)" />
        <circle cx="308" cy="23" r="1"   fill="rgba(180,190,240,0.30)" />
      </g>

      <!-- ══ Edifício esquerdo ══ -->
      <rect x="15" y="55" width="52" height="175" rx="2"
        fill="rgba(255,255,255,0.07)" stroke="rgba(255,255,255,0.12)" stroke-width="0.8" />
      <line x1="41" y1="55" x2="41" y2="40" stroke="rgba(255,255,255,0.25)" stroke-width="1.5" />
      <circle cx="41" cy="38" r="2" fill="rgba(255,100,100,0.60)" />
      <rect *ngFor="let w of windowsLeft"
        [attr.x]="w.x" [attr.y]="w.y" width="9" height="7" rx="1"
        [attr.fill]="w.fill" />

      <!-- ══ Edifício central ══ -->
      <rect x="80" y="30" width="90" height="200" rx="3"
        fill="rgba(255,255,255,0.10)" stroke="rgba(255,255,255,0.20)" stroke-width="1" />
      <rect x="80" y="30" width="90" height="12" rx="2" fill="rgba(255,255,255,0.06)" />
      <rect *ngFor="let w of windowsCenter"
        [attr.x]="w.x" [attr.y]="w.y" width="14" height="9" rx="1"
        [attr.fill]="w.fill" />
      <rect x="112" y="208" width="26" height="22" rx="2"
        fill="rgba(255,255,255,0.08)" stroke="rgba(255,255,255,0.15)" stroke-width="0.8" />

      <!-- ══ Edifício direito médio ══ -->
      <rect x="188" y="75" width="60" height="155" rx="2"
        fill="rgba(255,255,255,0.08)" stroke="rgba(255,255,255,0.13)" stroke-width="0.8" />
      <line x1="218" y1="75" x2="218" y2="60" stroke="rgba(255,255,255,0.20)" stroke-width="1.2" />
      <circle cx="218" cy="58" r="1.8" fill="rgba(255,100,100,0.55)" />
      <rect *ngFor="let w of windowsMid"
        [attr.x]="w.x" [attr.y]="w.y" width="11" height="8" rx="1"
        [attr.fill]="w.fill" />

      <!-- ══ Edifício direito baixo ══ -->
      <rect x="266" y="110" width="48" height="120" rx="2"
        fill="rgba(255,255,255,0.06)" stroke="rgba(255,255,255,0.10)" stroke-width="0.8" />
      <rect *ngFor="let w of windowsSmall"
        [attr.x]="w.x" [attr.y]="w.y" width="9" height="7" rx="1"
        [attr.fill]="w.fill" />

      <!-- ══ Casa residencial ══ -->
      <rect x="318" y="155" width="36" height="75" rx="2"
        fill="rgba(255,255,255,0.12)" stroke="rgba(255,255,255,0.20)" stroke-width="0.8" />
      <polygon points="308,155 336,135 364,155"
        fill="rgba(255,255,255,0.15)" stroke="rgba(255,255,255,0.20)" stroke-width="0.8" />
      <rect x="348" y="138" width="6"  height="16" rx="1" fill="rgba(255,255,255,0.10)" />
      <rect x="325" y="163" width="10" height="9"  rx="1" fill="rgba(255,220,100,0.60)" />
      <rect x="337" y="163" width="10" height="9"  rx="1" fill="rgba(255,220,100,0.45)" />
      <rect x="329" y="195" width="14" height="35" rx="2" fill="rgba(255,255,255,0.10)" />
      <circle cx="340" cy="213" r="1.5" fill="rgba(255,220,100,0.70)" />

      <!-- ══ Calçada ══ -->
      <rect x="0" y="228" width="360" height="2" rx="1" fill="rgba(255,255,255,0.15)" />

      <!-- ══ Árvores ══ -->
      <rect x="68"  y="205" width="4" height="23" fill="rgba(255,255,255,0.12)" />
      <ellipse cx="70"  cy="200" rx="9" ry="11" fill="rgba(255,255,255,0.08)" />
      <rect x="304" y="210" width="3" height="18" fill="rgba(255,255,255,0.12)" />
      <ellipse cx="305" cy="206" rx="7" ry="9"   fill="rgba(255,255,255,0.07)" />
    </svg>
  `,
})
export class BuildingIllustrationComponent implements OnInit, OnDestroy {
  @Input() dark?: boolean;
  @Input() className?: string;

  isDark = false;

  private observer?: MutationObserver;

  readonly stars: number[][] = [
    [50, 20], [80, 10], [200, 15], [250, 8],
    [140, 25], [30, 35], [100, 6], [170, 18],
    [290, 12], [320, 22], [60, 40], [220, 5],
  ];

  readonly sunRays = Array.from({ length: 8 }, (_, i) => {
    const angle = (i * 45 * Math.PI) / 180;
    return {
      x1: +(310 + Math.cos(angle) * 13).toFixed(2),
      y1: +(28  + Math.sin(angle) * 13).toFixed(2),
      x2: +(310 + Math.cos(angle) * 18).toFixed(2),
      y2: +(28  + Math.sin(angle) * 18).toFixed(2),
    };
  });

  readonly windowsLeft = this.buildWindows(
    [65, 80, 95, 110, 125],
    [22, 35, 48],
    (row, col) => (row + col) % 2 === 0
      ? `rgba(255,220,100,${(0.30 + col * 0.10).toFixed(2)})`
      : 'rgba(255,255,255,0.12)'
  );

  readonly windowsCenter = this.buildWindows(
    [48, 63, 78, 93, 108, 123, 138],
    [88, 107, 126, 145],
    (row, col) => (row + col) % 2 === 0
      ? `rgba(255,220,100,${(0.35 + col * 0.06).toFixed(2)})`
      : 'rgba(255,255,255,0.15)'
  );

  readonly windowsMid = this.buildWindows(
    [85, 99, 113, 127, 141],
    [196, 212, 228],
    (row, col) => (row + col) % 2 === 0
      ? `rgba(255,220,100,${(0.30 + col * 0.08).toFixed(2)})`
      : 'rgba(255,255,255,0.12)'
  );

  readonly windowsSmall = this.buildWindows(
    [118, 131, 144, 157],
    [272, 285, 298],
    (row, col) => (row + col) % 2 === 0
      ? `rgba(255,220,100,${(0.28 + col * 0.06).toFixed(2)})`
      : 'rgba(255,255,255,0.10)'
  );

  constructor(private cdr: ChangeDetectorRef) {}

  ngOnInit(): void {
    // Se a prop `dark` for passada, usa ela diretamente
    if (this.dark !== undefined) {
      this.isDark = this.dark;
      return;
    }

    this.isDark = document.documentElement.classList.contains('dark');

    // Observa mudanças de classe no <html>
    this.observer = new MutationObserver(() => {
      this.isDark = document.documentElement.classList.contains('dark');
      this.cdr.markForCheck();
    });

    this.observer.observe(document.documentElement, {
      attributes: true,
      attributeFilter: ['class'],
    });
  }

  ngOnDestroy(): void {
    this.observer?.disconnect();
  }

  private buildWindows(
    rows: number[],
    cols: number[],
    fillFn: (row: number, col: number) => string
  ): { x: number; y: number; fill: string }[] {
    return rows.flatMap((y, row) =>
      cols.map((x, col) => ({ x, y, fill: fillFn(row, col) }))
    );
  }
}