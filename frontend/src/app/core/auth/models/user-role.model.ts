/** Papéis alinhados ao backend (entities.Role) */
export type UserRole = 'owner' | 'tenant';

export const USER_ROLE_OPTIONS = [
  { value: 'owner'  as UserRole, label: 'Proprietário' },
  { value: 'tenant' as UserRole, label: 'Inquilino'    },
] as const;