import { Component, signal } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { PageHeaderComponent } from '../../core/layout/page-header/page-header.component';
import { ButtonComponent } from '../../components/button/button.component';
import { CardComponent } from '../../components/card/card.component';
import { InputComponent } from '../../components/input/input.component';
import { SelectComponent } from '../../components/select/select.component';
import { TextareaComponent } from '../../components/text-area/text-area.component';
import { AlertComponent } from '../../components/alert/alert.component';
import { AnimateOnViewDirective } from '../../shared/directives/animate-on-view.directive';
import { StaggerDirective } from '../../shared/directives/stagger.directive';
import { IconComponent } from '../../components/icon/icon.component';

@Component({
  selector: 'app-contracts-page',
  standalone: true,
  imports: [
    FormsModule,
    PageHeaderComponent,
    ButtonComponent,
    CardComponent,
    InputComponent,
    SelectComponent,
    TextareaComponent,
    AlertComponent,
    AnimateOnViewDirective,
    StaggerDirective,
    IconComponent,
  ],
  templateUrl: './contracts.page.html',

})
export class ContractsPage {
  currentStep = signal(1);
  steps = [
    { id: 1, label: 'Seleção' },
    { id: 2, label: 'Datas e valores' },
    { id: 3, label: 'Reajuste' },
    { id: 4, label: 'Regras' },
    { id: 5, label: 'Resumo' },
  ];
  contract: any = {};
  properties = [
    { value: '1', label: 'Apartamento 302 — Centro' },
    { value: '2', label: 'Casa 14 — Jardim' },
    { value: '3', label: 'Studio 07 — Paulista' },
  ];
  tenants = [
    { value: '1', label: 'Carlos Mendes' },
    { value: '2', label: 'Ana Ribeiro' },
  ];
  indexes = [
    { value: 'igpm', label: 'IGP-M' },
    { value: 'ipca', label: 'IPCA' },
  ];
  guarantees = [
    { value: 'caucao', label: 'Caução' },
    { value: 'fiador', label: 'Fiador' },
    { value: 'seguro', label: 'Seguro-fiança' },
  ];

  summary() {
    return [
      { label: 'Propriedade', value: this.properties.find(p => p.value === this.contract.propertyId)?.label || '—' },
      { label: 'Inquilino', value: this.tenants.find(t => t.value === this.contract.tenantId)?.label || '—' },
      { label: 'Início', value: this.contract.startDate || '—' },
      { label: 'Término', value: this.contract.endDate || '—' },
      { label: 'Valor do aluguel', value: this.contract.rent ? `R$ ${this.contract.rent}` : '—' },
      { label: 'Garantia', value: this.guarantees.find(g => g.value === this.contract.guarantee)?.label || '—' },
    ];
  }
  stepClass(id: number): string {
    if (id === this.currentStep()) return 'text-brand-600 dark:text-brand-400';
    if (id < this.currentStep()) return 'text-emerald-600';
    return 'text-text-secondary';
  }
  next() { if (this.currentStep() < this.steps.length) this.currentStep.update(s => s + 1); }
  prev() { if (this.currentStep() > 1) this.currentStep.update(s => s - 1); }
  goTo(id: number) { this.currentStep.set(id); }
  save() { alert('Contrato salvo!'); }
}