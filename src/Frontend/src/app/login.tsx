import React, { useState } from 'react';
import { View, Text, TouchableOpacity } from 'react-native';
import { SafeAreaView, useSafeAreaInsets } from 'react-native-safe-area-context';
import { Ionicons } from '@expo/vector-icons';
import { router } from 'expo-router';

import Input from '../components/shared/Input';
import Button from '../components/shared/Button';
import StatusBanner from '../components/shared/StatusBanner';
import { apiService } from '../services/api';
import { useAuth } from '../contexts/AuthContext';

export default function LoginScreen() {
  const { login } = useAuth();
  const insets = useSafeAreaInsets();
  const [form, setForm] = useState({
    email: '',
    password: ''
  });
  const [loading, setLoading] = useState(false);
  const [status, setStatus] = useState<{ type: 'success' | 'error' | null; message: string }>({ type: null, message: '' });

  const updateForm = (field: keyof typeof form, value: string) => {
    setForm(prev => ({ ...prev, [field]: value }));
  };

  const handleLogin = async () => {
    // Validação básica dos campos
    if (!form.email || !form.password) {
      setStatus({ type: 'error', message: 'Por favor, preencha todos os campos' });
      setTimeout(() => setStatus({ type: null, message: '' }), 3000);
      return;
    }

    setLoading(true);
    setStatus({ type: null, message: '' });
    
    try {
      const data = await apiService.login(
        form.email.toLowerCase(),
        form.password
      );
      
      // Salva autenticação no contexto (memória apenas)
      login(data.token, data.owner);
      
      setStatus({ type: 'success', message: `Bem-vindo, ${data.owner.name}!` });
      // Navega para dashboard após login bem-sucedido
      setTimeout(() => {
        router.push('./dashboard');
      }, 1500);
    } catch (error: any) {
      setStatus({ type: 'error', message: error.message || 'Erro ao fazer login' });
      // Auto-limpa mensagem de erro após 3s
      setTimeout(() => setStatus({ type: null, message: '' }), 3000);
      console.error('Login error:', error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <SafeAreaView className="flex-1 bg-white">
      {/* Main Container */}
      <View className="flex-1 px-8 justify-center">
        <View className="w-full max-w-md mx-auto">
          {/* Top Section */}
          <View>
            
            {/* Welcome Text */}
            <View className="items-center mb-4">
              <Text className="text-orange-800 text-2xl font-semibold text-center mb-1">
                Bem-vindo de volta!
              </Text>
              <Text className="text-orange-600 text-center text-sm">
                Acesse sua conta para continuar
              </Text>
            </View>
            
            {/* Form */}
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
              
              {/* Forgot Password */}
              <View className="items-end">
                <TouchableOpacity>
                  <Text className="text-orange-600 text-sm font-medium">Esqueceu a senha?</Text>
                </TouchableOpacity>
              </View>
              
              <Button onPress={handleLogin} loading={loading} title="Entrar" />
            </View>
          </View>
          
          {/* Bottom Section */}
          <View className="mt-8">
            {/* Create Account Link */}
            <View className="items-center mb-6">
              <TouchableOpacity onPress={() => router.push('./register')}>
                <Text className="text-orange-600 text-sm font-medium">
                  Não tem uma conta? <Text className="font-bold">Criar conta</Text>
                </Text>
              </TouchableOpacity>
            </View>
            
            {/* Social Login */}
            <View className="items-center">
              <View className="flex-row items-center mb-4">
                <View className="flex-1 h-0.5 w-20 bg-gray-500"></View>
                <Text className="text-gray-500 text-xs mx-4">Ou entre com</Text>
                <View className="flex-1 h-0.5 w-20 bg-gray-500"></View>
              </View>
              <View className="flex-row gap-3">
                {/* Google Icon */}
                <TouchableOpacity className="w-12 h-12 bg-gray-100 rounded-full items-center justify-center">
                  <Ionicons name="logo-google" size={18} color="#DB4437" />
                </TouchableOpacity>
                {/* Apple Icon */}
                <TouchableOpacity className="w-12 h-12 bg-gray-100 rounded-full items-center justify-center">
                  <Ionicons name="logo-apple" size={18} color="#000" />
                </TouchableOpacity>
                {/* Facebook Icon */}
                <TouchableOpacity className="w-12 h-12 bg-gray-100 rounded-full items-center justify-center">
                  <Ionicons name="logo-facebook" size={18} color="#1877F2" />
                </TouchableOpacity>
              </View>
            </View>
          </View>
        </View>
      </View>
      
      {status.type && (
        <View 
          className="absolute bottom-0 left-0 right-0"
          style={{ paddingBottom: insets.bottom }}
        >
          <StatusBanner type={status.type} message={status.message} />
        </View>
      )}
    </SafeAreaView>
  );
}