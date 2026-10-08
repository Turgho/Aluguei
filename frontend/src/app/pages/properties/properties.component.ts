import { Component, signal, OnInit, inject } from '@angular/core';
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
import { IconComponent } from '../../components/icon/icon.component';
import { CheckboxComponent } from '../../components/checkbox/checkbox.component';
import { AnimateOnViewDirective } from '../../shared/directives/animate-on-view.directive';
import { StaggerDirective } from '../../shared/directives/stagger.directive';

// ─── Tipos espelhados do backend ──────────────────────────────────────────────

type PropertyType =
  | 'house' | 'apartment' | 'studio' | 'loft'
  | 'business' | 'commercial' | 'office' | 'store' | 'land';

type PropertyStatus = 'available' | 'rented' | 'inactive' | 'sold';

// ─── Form model ───────────────────────────────────────────────────────────────

interface PropertyAddress {
  zipCode:      string;
  street:       string;
  number:       string;
  complement:   string;
  neighborhood: string;
  city:         string;
  state:        string;
}

interface PropertyForm {
  title:               string;
  description:         string;
  type:                PropertyType;
  status:              PropertyStatus;
  // Exibidos em R$ — convertidos para centavos no submit
  priceCents:          number;
  condominiumFeeCents: number;
  iptuCents:           number;
  // Características (int no backend)
  areaM2:              number;
  bedrooms:            number;
  bathrooms:           number;
  suites:              number;
  parkingSpaces:       number;
  // Comodidades
  furnished:           boolean;
  petFriendly:         boolean;
  hasBalcony:          boolean;
  hasElevator:         boolean;
  // Endereço
  address:             PropertyAddress;
}

const emptyForm = (): PropertyForm => ({
  title:               '',
  description:         '',
  type:                'apartment',
  status:              'available',
  priceCents:          0,
  condominiumFeeCents: 0,
  iptuCents:           0,
  areaM2:              0,
  bedrooms:            0,
  bathrooms:           0,
  suites:              0,
  parkingSpaces:       0,
  furnished:           false,
  petFriendly:         false,
  hasBalcony:          false,
  hasElevator:         false,
  address: {
    zipCode:      '',
    street:       '',
    number:       '',
    complement:   '',
    neighborhood: '',
    city:         '',
    state:        '',
  },
});

// ─── Dados de listagem (substituir por service futuramente) ───────────────────

interface PropertyListItem {
  id:       number;
  name:     string;
  address:  string;
  rent:     string;
  occupied: boolean;
  icon:     string;
}

// ─── Component ────────────────────────────────────────────────────────────────

@Component({
  selector: 'app-properties-page',
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
    IconComponent,
    CheckboxComponent,
    AnimateOnViewDirective,
    StaggerDirective,
  ],
  templateUrl: './properties.page.html',
})
export class PropertiesPage implements OnInit {

  // ── Estado ──────────────────────────────────────────────────────────────────

  loading   = signal(true);
  modalOpen = signal(false);
  search    = '';
  form: PropertyForm = emptyForm();

  // ── Opções dos selects ────────────────────────────────────────────────────

  readonly propertyTypes: { value: PropertyType; label: string }[] = [
    { value: 'apartment',  label: 'Apartamento' },
    { value: 'house',      label: 'Casa'        },
    { value: 'studio',     label: 'Studio'      },
    { value: 'loft',       label: 'Loft'        },
    { value: 'business',   label: 'Comercial'   },
    { value: 'commercial', label: 'Galpão'      },
    { value: 'office',     label: 'Escritório'  },
    { value: 'store',      label: 'Loja'        },
    { value: 'land',       label: 'Terreno'     },
  ];

  readonly propertyStatuses: { value: PropertyStatus; label: string }[] = [
    { value: 'available', label: 'Disponível' },
    { value: 'rented',    label: 'Alugado'    },
    { value: 'inactive',  label: 'Inativo'    },
    { value: 'sold',      label: 'Vendido'    },
  ];

  readonly brazilStates = [
    { value: 'AC', label: 'AC' }, { value: 'AL', label: 'AL' },
    { value: 'AP', label: 'AP' }, { value: 'AM', label: 'AM' },
    { value: 'BA', label: 'BA' }, { value: 'CE', label: 'CE' },
    { value: 'DF', label: 'DF' }, { value: 'ES', label: 'ES' },
    { value: 'GO', label: 'GO' }, { value: 'MA', label: 'MA' },
    { value: 'MT', label: 'MT' }, { value: 'MS', label: 'MS' },
    { value: 'MG', label: 'MG' }, { value: 'PA', label: 'PA' },
    { value: 'PB', label: 'PB' }, { value: 'PR', label: 'PR' },
    { value: 'PE', label: 'PE' }, { value: 'PI', label: 'PI' },
    { value: 'RJ', label: 'RJ' }, { value: 'RN', label: 'RN' },
    { value: 'RS', label: 'RS' }, { value: 'RO', label: 'RO' },
    { value: 'RR', label: 'RR' }, { value: 'SC', label: 'SC' },
    { value: 'SP', label: 'SP' }, { value: 'SE', label: 'SE' },
    { value: 'TO', label: 'TO' },
  ];

  // ── Dados mockados (substituir por service) ──────────────────────────────

  readonly properties: PropertyListItem[] = [
    { id: 1, name: 'Apartamento 302',  address: 'Rua das Flores, 120 — Centro',    rent: 'R$ 1.800', occupied: true,  icon: '🏢' },
    { id: 2, name: 'Casa 14',          address: 'Alameda dos Ipês, 45 — Jardim',   rent: 'R$ 2.500', occupied: true,  icon: '🏡' },
    { id: 3, name: 'Studio 07',        address: 'Av. Paulista, 1500 — Bela Vista', rent: 'R$ 1.500', occupied: false, icon: '🏬' },
    { id: 4, name: 'Casa 22',          address: 'Rua Verde, 88 — Vila Nova',       rent: 'R$ 2.200', occupied: true,  icon: '🏡' },
    { id: 5, name: 'Apto 101',         address: 'Rua Norte, 33 — Sul',             rent: 'R$ 1.300', occupied: false, icon: '🏢' },
  ];

  // ── Computed ─────────────────────────────────────────────────────────────

  get filtered(): PropertyListItem[] {
    const s = this.search.toLowerCase();
    return this.properties.filter(p =>
      p.name.toLowerCase().includes(s) ||
      p.address.toLowerCase().includes(s)
    );
  }

  /** Terrenos não possuem cômodos — backend valida a mesma regra */
  get isLand(): boolean {
    return this.form.type === 'land';
  }

  // ── Lifecycle ─────────────────────────────────────────────────────────────

  ngOnInit(): void {
    setTimeout(() => this.loading.set(false), 800);
  }

  // ── Modal ─────────────────────────────────────────────────────────────────

  openModal(): void {
    this.form = emptyForm();
    this.modalOpen.set(true);
  }

  closeModal(): void {
    this.modalOpen.set(false);
    this.form = emptyForm();
  }

  // ── CEP (ViaCEP) ──────────────────────────────────────────────────────────

  fetchingCep = signal(false);

  async onCepBlur(): Promise<void> {
    const cep = this.form.address.zipCode.replace(/\D/g, '');
    if (cep.length !== 8) return;

    this.fetchingCep.set(true);
    try {
      const res  = await fetch(`https://viacep.com.br/ws/${cep}/json/`);
      const data = await res.json();
      if (data.erro) return;

      this.form.address.street       = data.logradouro  ?? '';
      this.form.address.neighborhood = data.bairro      ?? '';
      this.form.address.city         = data.localidade  ?? '';
      this.form.address.state        = data.uf          ?? '';
    } catch {
      // CEP inválido ou sem conexão — usuário preenche manualmente
    } finally {
      this.fetchingCep.set(false);
    }
  }

  // ── Submit ────────────────────────────────────────────────────────────────

  save(): void {
    const payload = {
      ...this.form,
      // Converte R$ → centavos para o backend
      priceCents:          Math.round(this.form.priceCents * 100),
      condominiumFeeCents: Math.round(this.form.condominiumFeeCents * 100),
      iptuCents:           Math.round(this.form.iptuCents * 100),
      // Terreno zera cômodos (regra espelhada do backend)
      bedrooms:     this.isLand ? 0 : this.form.bedrooms,
      bathrooms:    this.isLand ? 0 : this.form.bathrooms,
      suites:       this.isLand ? 0 : this.form.suites,
    };

    console.log('payload →', payload);
    this.closeModal();
  }
}