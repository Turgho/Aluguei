import { ChartData, ChartOptions } from 'chart.js';

// SummaryCard
export interface SummaryCard {
  label: string;
  value: string;
  iconPath: string;
  bg: string;
  darkBg: string;
}

// Property
export interface Property {
  name: string;
  address: string;
  status: 'Alugada' | 'Disponível';
  type: string;
  rent: string;
}

// Payment
export interface Payment {
  tenant: string;
  date: string;
  amount: string;
  method: string;
}

// Alert
export interface Alert {
  type: 'danger' | 'warning';
  title: string;
  description: string;
}

// DashboardData
export interface DashboardData {
  summaryCards: SummaryCard[];
  barData: ChartData<'bar'>;
  barOptions: ChartOptions<'bar'>;
  properties: Property[];
  payments: Payment[];
  alerts: Alert[];
}
