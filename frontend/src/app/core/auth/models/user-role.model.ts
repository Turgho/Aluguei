/** Papéis alinhados ao backend (entities.Role) */
export type UserRole = 'owner' | 'tenant';

export const USER_ROLE_OPTIONS: { value: UserRole; label: string }[] = [
  { value: 'owner', label: 'Proprietário' },
  { value: 'tenant', label: 'Inquilino' },
];
