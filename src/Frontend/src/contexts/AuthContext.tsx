import React, { createContext, useContext, useState } from 'react';
import type { Owner } from '../types/api';

interface AuthContextType {
  isAuthenticated: boolean;
  owner: Owner | null;
  token: string | null;
  login: (token: string, owner: Owner) => void;
  logout: () => void;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

// Provedor de autenticação que mantém estado em memória (não persiste)
export function AuthProvider({ children }: { children: React.ReactNode }) {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [owner, setOwner] = useState<Owner | null>(null);
  const [token, setToken] = useState<string | null>(null);

  const login = (newToken: string, ownerData: Owner) => {
    setToken(newToken);
    setOwner(ownerData);
    setIsAuthenticated(true);
  };

  const logout = () => {
    setToken(null);
    setOwner(null);
    setIsAuthenticated(false);
  };

  return (
    <AuthContext.Provider value={{
      isAuthenticated,
      owner,
      token,
      login,
      logout
    }}>
      {children}
    </AuthContext.Provider>
  );
}

export function useAuth() {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
}