import { Directive, ElementRef, Input, OnInit, OnDestroy, inject, PLATFORM_ID } from '@angular/core';
import { isPlatformBrowser } from '@angular/common';

interface ChildState {
  element: HTMLElement;
  originalOpacity: string;
}

@Directive({
  selector: '[appStagger]',
  standalone: true,
})
export class StaggerDirective implements OnInit, OnDestroy {
  @Input() appStagger: 'fade-up' | 'fade-in' | 'scale-in' = 'fade-up';
  @Input() staggerDelay: number = 60;
  @Input() staggerDuration: number = 300;

  private el = inject(ElementRef);
  private platformId = inject(PLATFORM_ID);
  private observer?: IntersectionObserver;
  private childrenStates: ChildState[] = [];
  private prefersReducedMotion = false;
  private isLowEndDevice = false;

  ngOnInit() {
    if (!isPlatformBrowser(this.platformId)) return;

    this.prefersReducedMotion = window.matchMedia('(prefers-reduced-motion: reduce)').matches;
    this.isLowEndDevice = this.detectLowEndDevice();

    const children = Array.from(this.el.nativeElement.children) as HTMLElement[];

    // Se prefere movimento reduzido ou dispositivo fraco, mostra tudo direto
    if (this.prefersReducedMotion || this.isLowEndDevice) {
      children.forEach((child) => {
        child.style.opacity = this.getOriginalOpacity(child);
      });
      return;
    }

    // Captura opacidade original de cada filho (sem getComputedStyle)
    this.childrenStates = children.map((child) => {
      const originalOpacity = this.getOriginalOpacity(child);
      child.style.opacity = '0';
      child.style.transform = this.getInitialTransform();
      return { element: child, originalOpacity };
    });

    this.observer = new IntersectionObserver(
      (entries) => {
        entries.forEach((entry) => {
          if (entry.isIntersecting) {
            this.animateChildren();
            this.observer?.unobserve(entry.target);
          }
        });
      },
      { 
        threshold: 0.1,
        rootMargin: '50px'
      }
    );

    this.observer.observe(this.el.nativeElement);
  }

  ngOnDestroy() {
    this.observer?.disconnect();
  }

  /**
   * Detecta opacidade via classes CSS (rápido, sem reflow)
   */
  private getOriginalOpacity(element: HTMLElement): string {
    const classList = Array.from(element.classList) as string[];
    const opacityClass = classList.find(c => c.startsWith('opacity-'));
    
    if (opacityClass) {
      const value = parseInt(opacityClass.split('-')[1]);
      return (value / 100).toString();
    }
    
    return '1';
  }

  private detectLowEndDevice(): boolean {
    const nav = navigator as any;
    const cores = nav.hardwareConcurrency || 4;
    const memory = nav.deviceMemory || 4;
    const connection = nav.connection?.effectiveType || '4g';
    
    return cores <= 2 || memory <= 2 || connection === '2g' || connection === 'slow-2g';
  }

  private getInitialTransform(): string {
    return this.appStagger === 'scale-in' ? 'scale(0.95)' : 'translateY(20px)';
  }

  /**
   * Animação stagger otimizada com requestAnimationFrame
   */
  private animateChildren() {
    this.childrenStates.forEach(({ element, originalOpacity }, index) => {
      setTimeout(() => {
        requestAnimationFrame(() => {
          element.style.transition = `opacity ${this.staggerDuration}ms ease-out, transform ${this.staggerDuration}ms ease-out`;
          element.style.opacity = originalOpacity;
          element.style.transform = 'none';

          // Cleanup
          setTimeout(() => {
            requestAnimationFrame(() => {
              element.style.removeProperty('opacity');
              element.style.removeProperty('transform');
              element.style.removeProperty('transition');
            });
          }, this.staggerDuration + 50);
        });
      }, index * this.staggerDelay);
    });
  }
}