import { BackendUser, StoredUser, User } from './models/user.model';

// Mapeia o usuário do backend para o modelo do usuário
export function mapBackendUser(backendUser: BackendUser): User {
  return {
    id: backendUser.id,
    name: `${backendUser.first_name} ${backendUser.last_name}`.trim(),
    email: backendUser.email,
    role: backendUser.role,
    phone: backendUser.phone,
  };
}

// Mapeia o usuário para o modelo armazenado
export function toStoredUser(user: User): StoredUser {
  return {
    id: user.id,
    name: user.name,
    email: user.email,
    role: user.role,
  };
}

// Mapeia o usuário armazenado para o modelo do usuário
export function fromStoredUser(stored: StoredUser): User {
  return { ...stored };
}
