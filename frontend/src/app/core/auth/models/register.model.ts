import { UserRole } from './user-role.model';

// RegisterPayload
export interface RegisterPayload {
  first_name: string;
  last_name: string;
  email: string;
  cpf: string;
  phone?: string;
  password: string;
  role: UserRole;
}
