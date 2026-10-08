import { Component, signal, OnInit } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { PageHeaderComponent } from '../../core/layout/page-header/page-header.component';
import { ButtonComponent } from '../../components/button/button.component';
import { CardComponent } from '../../components/card/card.component';
import { InputComponent } from '../../components/input/input.component';
import { SelectComponent } from '../../components/select/select.component';
import { StatCardComponent } from '../../components/stat-card/stat-card.component';
import { BadgeComponent } from '../../components/badge/badge.component';
import { EmptyStateComponent } from '../../components/empty-state/empty-state.component';
import { SkeletonComponent } from '../../components/skeleton/skeleton.component';
import { AnimateOnViewDirective } from '../../shared/directives/animate-on-view.directive';
import { StaggerDirective } from '../../shared/directives/stagger.directive';
import { IconComponent } from '../../components/icon/icon.component';

@Component({
  selector: 'app-reports-page',
  standalone: true,
  imports: [
    FormsModule,
    PageHeaderComponent,
    ButtonComponent,
    CardComponent,
    InputComponent,
    SelectComponent,
    StatCardComponent,
    BadgeComponent,
    EmptyStateComponent,
    SkeletonComponent,
    AnimateOnViewDirective,
    StaggerDirective,
    IconComponent,
  ],
  templateUrl: './reports.page.html',
})
export class ReportsPage implements OnInit {
  loading = signal(true);
  filters: any = { property: 'all', tenant: 'all', status: 'all' };
  rows = [
    { id: 1, date: '05/09/2025', tenant: 'Carlos Mendes', property: 'Apto 302', value: 'R$ 1.800', status: 'Atrasado' },
    { id: 2, date: '10/09/2025', tenant: 'Ana Ribeiro', property: 'Casa 14', value: 'R$ 2.500', status: 'Pendente' },
    { id: 3, date: '05/09/2025', tenant: 'Marcos Souza', property: 'Casa 22', value: 'R$ 2.200', status: 'Pago' },
    { id: 4, date: '05/09/2025', tenant: 'Juliana Costa', property: 'Apto 101', value: 'R$ 1.300', status: 'Pago' },
    { id: 5, date: '03/09/2025', tenant: 'Pedro Lima', property: 'Studio 07', value: 'R$ 1.500', status: 'Pago' },
  ];
  ngOnInit() { setTimeout(() => this.loading.set(false), 700); }
  apply() { this.loading.set(true); setTimeout(() => this.loading.set(false), 500); }
  export(format: string) { alert(`Exportando ${format.toUpperCase()}...`); }
}