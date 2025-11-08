import React, { createContext, useContext, useState, useEffect } from 'react';
import AsyncStorage from '@react-native-async-storage/async-storage';
import type { Owner } from '../types/api';

interface AuthContextType {
  isAuthenticated: boolean;
  owner: Owner | null;
  token: string | null;
  login: (token: string, owner: Owner) => Promise<void>;
  logout: () => Promise<void>;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export function AuthProvider({ children }: { children: React.ReactNode }) {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [owner, setOwner] = useState<Owner | null>(null);
  const [token, setToken] = useState<string | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    loadAuthData();
  }, []);

  const loadAuthData = async () => {
    try {
      const savedToken = await AsyncStorage.getItem('token');
      const savedOwner = await AsyncStorage.getItem('owner');
      
      if (savedToken && savedOwner) {
        setToken(savedToken);
        setOwner(JSON.parse(savedOwner));
        setIsAuthenticated(true);
      }
    } catch (error) {
      console.error('Erro ao carregar dados de autenticação:', error);
    } finally {
      setLoading(false);
    }
  };

  const login = async (newToken: string, ownerData: Owner) => {
    try {
      await AsyncStorage.setItem('token', newToken);
      await AsyncStorage.setItem('owner', JSON.stringify(ownerData));
      
      setToken(newToken);
      setOwner(ownerData);
      setIsAuthenticated(true);
    } catch (error) {
      console.error('Erro ao salvar dados de autenticação:', error);
    }
  };

  const logout = async () => {
    try {
      await AsyncStorage.removeItem('token');
      await AsyncStorage.removeItem('owner');
    } catch (error) {
      console.error('Erro ao limpar dados de autenticação:', error);
    }
    
    setToken(null);
    setOwner(null);
    setIsAuthenticated(false);
  };

  if (loading) {
    return null; // ou um componente de loading
  }

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