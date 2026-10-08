import { Directive, ElementRef, Input, OnInit, OnDestroy, inject, PLATFORM_ID } from '@angular/core';
import { isPlatformBrowser } from '@angular/common';

@Directive({
  selector: '[appAnimateOnView]',
  standalone: true,
})
export class AnimateOnViewDirective implements OnInit, OnDestroy {
  @Input() appAnimateOnView: 'fade-in' | 'fade-up' | 'fade-down' | 'scale-in' | 'slide-right' | 'slide-left' = 'fade-up';
  @Input() delay: number = 0;
  @Input() duration: number = 300;
  @Input() threshold: number = 0.1;

  private el = inject(ElementRef);
  private platformId = inject(PLATFORM_ID);
  private observer?: IntersectionObserver;
  private originalOpacity: string = '1';
  private prefersReducedMotion = false;
  private isLowEndDevice = false;

  ngOnInit() {
    // SSR: não executa animações
    if (!isPlatformBrowser(this.platformId)) return;

    // Detecta preferência de movimento reduzido
    this.prefersReducedMotion = window.matchMedia('(prefers-reduced-motion: reduce)').matches;

    // Detecta dispositivo de baixa performance
    this.isLowEndDevice = this.detectLowEndDevice();

    // Se prefere movimento reduzido ou dispositivo fraco, mostra direto sem animação
    if (this.prefersReducedMotion || this.isLowEndDevice) {
      this.el.nativeElement.style.opacity = this.getOriginalOpacity();
      return;
    }

    // Captura opacidade original de forma otimizada
    this.originalOpacity = this.getOriginalOpacity();

    // Estado inicial
    this.el.nativeElement.style.opacity = '0';
    this.el.nativeElement.style.transform = this.getInitialTransform();

    // Observer com rootMargin para performance
    this.observer = new IntersectionObserver(
      (entries) => {
        entries.forEach((entry) => {
          if (entry.isIntersecting) {
            this.animate();
            this.observer?.unobserve(entry.target);
          }
        });
      },
      { 
        threshold: this.threshold,
        rootMargin: '50px' // Começa a animar um pouco antes de entrar no viewport
      }
    );

    this.observer.observe(this.el.nativeElement);
  }

  ngOnDestroy() {
    this.observer?.disconnect();
  }

  /**
   * Detecta opacidade original sem forçar reflow
   * Verifica classes Tailwind (opacity-*) em vez de getComputedStyle
   */
  private getOriginalOpacity(): string {
    const classList = Array.from(this.el.nativeElement.classList) as string[];
    const opacityClass = classList.find(c => c.startsWith('opacity-'));
    
    if (opacityClass) {
      const value = parseInt(opacityClass.split('-')[1]);
      return (value / 100).toString();
    }
    
    return '1';
  }

  /**
   * Detecta dispositivo de baixa performance
   */
  private detectLowEndDevice(): boolean {
    const nav = navigator as any;
    const cores = nav.hardwareConcurrency || 4;
    const memory = nav.deviceMemory || 4;
    const connection = nav.connection?.effectiveType || '4g';
    
    return cores <= 2 || memory <= 2 || connection === '2g' || connection === 'slow-2g';
  }

  private getInitialTransform(): string {
    const transforms = {
      'fade-in': 'none',
      'fade-up': 'translateY(20px)',
      'fade-down': 'translateY(-20px)',
      'scale-in': 'scale(0.95)',
      'slide-right': 'translateX(-20px)',
      'slide-left': 'translateX(20px)',
    };
    return transforms[this.appAnimateOnView];
  }

  /**
   * Animação otimizada com requestAnimationFrame
   */
  private animate() {
    setTimeout(() => {
      requestAnimationFrame(() => {
        this.el.nativeElement.style.transition = `opacity ${this.duration}ms ease-out, transform ${this.duration}ms ease-out`;
        this.el.nativeElement.style.opacity = this.originalOpacity;
        this.el.nativeElement.style.transform = 'none';

        // Cleanup após animação (remove inline styles)
        setTimeout(() => {
          requestAnimationFrame(() => {
            this.el.nativeElement.style.removeProperty('opacity');
            this.el.nativeElement.style.removeProperty('transform');
            this.el.nativeElement.style.removeProperty('transition');
          });
        }, this.duration + 50);
      });
    }, this.delay);
  }
}