import React, { useState } from 'react';
import { View } from 'react-native';
import { router } from 'expo-router';
import Input from '../../../components/forms/Input';
import Button from '../../../components/forms/Button';
import { apiService } from '../../../services/api';
import { useAuth } from '../../../contexts/AuthContext';

interface LoginFormProps {
  onStatusChange: (status: { type: 'success' | 'error' | null; message: string }) => void;
}

export default function LoginForm({ onStatusChange }: LoginFormProps) {
  const { login } = useAuth();
  const [form, setForm] = useState({
    email: '',
    password: ''
  });
  const [loading, setLoading] = useState(false);

  const updateForm = (field: keyof typeof form, value: string) => {
    setForm(prev => ({ ...prev, [field]: value }));
  };

  const handleLogin = async () => {
    if (!form.email || !form.password) {
      onStatusChange({ type: 'error', message: 'Por favor, preencha todos os campos' });
      return;
    }

    setLoading(true);
    onStatusChange({ type: null, message: '' });
    
    try {
      const data = await apiService.login(
        form.email.toLowerCase(),
        form.password
      );
      
      await login(data.token, data.owner);
      onStatusChange({ type: 'success', message: `Bem-vindo, ${data.owner.name}!` });
      
      setTimeout(() => {
        router.push('./dashboard');
      }, 1500);
    } catch (error: any) {
      onStatusChange({ type: 'error', message: error.message || 'Erro ao fazer login' });
    } finally {
      setLoading(false);
    }
  };

  return (
    <View>
      <Input
        label="Email"
        placeholder="seu@email.com"
        value={form.email}
        onChangeText={(value: string) => updateForm('email', value)}
        iconName="mail"
        keyboardType="email-address"
      />
      
      <Input
        label="Senha"
        placeholder="Sua senha"
        value={form.password}
        onChangeText={(value: string) => updateForm('password', value)}
        iconName="key"
        secureTextEntry
      />
      
      <Button onPress={handleLogin} loading={loading} title="Entrar" />
    </View>
  );
}