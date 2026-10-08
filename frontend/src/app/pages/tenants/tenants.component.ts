import { Component, signal, OnInit } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { PageHeaderComponent } from '../../core/layout/page-header/page-header.component';
import { ButtonComponent } from '../../components/button/button.component';
import { CardComponent } from '../../components/card/card.component';
import { InputComponent } from '../../components/input/input.component';
import { SelectComponent } from '../../components/select/select.component';
import { TextareaComponent } from '../../components/text-area/text-area.component';
import { BadgeComponent } from '../../components/badge/badge.component';
import { ModalComponent } from '../../components/modal/modal.component';
import { EmptyStateComponent } from '../../components/empty-state/empty-state.component';
import { SkeletonComponent } from '../../components/skeleton/skeleton.component';
import { AnimateOnViewDirective } from '../../shared/directives/animate-on-view.directive';
import { StaggerDirective } from '../../shared/directives/stagger.directive';
import { IconComponent } from '../../components/icon/icon.component';

@Component({
  selector: 'app-tenants-page',
  standalone: true,
  imports: [
    FormsModule,
    PageHeaderComponent,
    ButtonComponent,
    CardComponent,
    InputComponent,
    SelectComponent,
    TextareaComponent,
    BadgeComponent,
    ModalComponent,
    EmptyStateComponent,
    SkeletonComponent,
    AnimateOnViewDirective,
    StaggerDirective,
    IconComponent,
  ],
  templateUrl: './tenants.page.html',
})
export class TenantsPage implements OnInit {
  loading = signal(true);
  modalOpen = signal(false);
  search = '';
  statusFilter = 'all';
  form: any = {};
  
  tenants = [
    { id: 1, initials: 'CM', name: 'Carlos Mendes', cpf: '123.456.789-00', email: 'carlos@email.com', phone: '(11) 98765-4321', property: 'Apto 302', status: 'Ativo' },
    { id: 2, initials: 'AR', name: 'Ana Ribeiro', cpf: '987.654.321-00', email: 'ana@email.com', phone: '(11) 91234-5678', property: 'Casa 14', status: 'Ativo' },
    { id: 3, initials: 'MS', name: 'Marcos Souza', cpf: '456.789.123-00', email: 'marcos@email.com', phone: '(11) 97777-8888', property: 'Casa 22', status: 'Inativo' },
  ];
  
  get filtered() {
    const searchLower = this.search.toLowerCase();
    const statusMap: Record<string, string> = {
      'active': 'Ativo',
      'inactive': 'Inativo'
    };
    
    return this.tenants.filter(t => {
      const matchesSearch = 
        t.name.toLowerCase().includes(searchLower) ||
        t.cpf.includes(searchLower) ||
        t.email.toLowerCase().includes(searchLower);
      
      const matchesStatus = 
        this.statusFilter === 'all' || 
        t.status === statusMap[this.statusFilter];
      
      return matchesSearch && matchesStatus;
    });
  }
  
  ngOnInit() { 
    setTimeout(() => this.loading.set(false), 700); 
  }
  
  openModal() { 
    this.modalOpen.set(true); 
  }
  
  closeModal() { 
    this.modalOpen.set(false); 
    this.form = {}; 
  }
  
  save() { 
    this.closeModal(); 
  }
}