import { Component, inject, OnInit, signal } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { SidebarComponent } from '../../components/sidebar/sidebar.component';
import { TopbarComponent } from '../../components/topbar/topbar.component';
import { ThemeService } from '../theme/theme.service';

@Component({
  selector: 'app-main-layout',
  standalone: true,
  imports: [RouterOutlet, SidebarComponent, TopbarComponent],
  template: `
    <div class="flex min-h-screen bg-bg-secondary">
      <!-- Sidebar -->
      <app-sidebar [open]="sidebarOpen()" />
      
      <!-- Overlay mobile -->
      @if (sidebarOpen()) {
        <div 
          class="fixed inset-0 z-30 bg-black/40 backdrop-blur-sm lg:hidden" 
          (click)="sidebarOpen.set(false)"></div>
      }
      
      <!-- Conteúdo principal -->
      <div class="flex flex-1 flex-col min-w-0">
        <app-topbar (toggleMenu)="sidebarOpen.update(v => !v)" />
        <main class="flex-1 p-4 lg:p-6 overflow-x-hidden">
          <router-outlet />
        </main>
      </div>
    </div>
  `,
})
export class MainLayoutComponent implements OnInit {
  theme = inject(ThemeService);
  sidebarOpen = signal(false);
  ngOnInit() { this.theme.init(); }
}