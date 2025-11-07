import React, { useState } from 'react';
import { View, Text } from 'react-native';
import { SafeAreaView, useSafeAreaInsets } from 'react-native-safe-area-context';
import { Ionicons } from '@expo/vector-icons';
import LoginHeader from '../components/LoginHeader';
import LoginInput from '../components/LoginInput';
import LoginButton from '../components/LoginButton';
import StatusBanner from '../components/StatusBanner';
import type { LoginRequest, LoginResponse, ApiError } from '../types/api';

export default function LoginScreen() {
  const insets = useSafeAreaInsets();
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [loading, setLoading] = useState(false);
  const [status, setStatus] = useState<{ type: 'success' | 'error' | null; message: string }>({ type: null, message: '' });

  const handleLogin = async () => {
    if (!email || !password) {
      setStatus({ type: 'error', message: 'Por favor, preencha todos os campos' });
      setTimeout(() => setStatus({ type: null, message: '' }), 3000);
      return;
    }

    setLoading(true);
    setStatus({ type: null, message: '' });
    
    try {
      const response = await fetch('http://192.168.1.6:8080/api/v1/auth/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          email: email.toLowerCase(),
          password,
        }),
      });

      if (response.ok) {
        const data: LoginResponse = await response.json();
        setStatus({ type: 'success', message: `Bem-vindo, ${data.owner.name}!` });
        // TODO: Salvar token e navegar para dashboard
        console.log('Token:', data.token);
        console.log('Owner:', data.owner);
      } else {
        const errorData: ApiError = await response.json();
        setStatus({ type: 'error', message: errorData.error || 'Erro ao fazer login' });
        setTimeout(() => setStatus({ type: null, message: '' }), 3000);
      }
    } catch (error) {
      setStatus({ type: 'error', message: 'Erro de conexão. Verifique sua internet.' });
      setTimeout(() => setStatus({ type: null, message: '' }), 3000);
      console.error('Login error:', error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <SafeAreaView className="flex-1 bg-gradient-to-br from-orange-50 via-amber-50 to-yellow-50">
      {/* Background Pattern */}
      <View className="absolute inset-0 opacity-5">
        <View className="flex-row justify-around items-center pt-20">
          <Ionicons name="home" size={40} color="#F97316" />
          <Ionicons name="business" size={35} color="#EA580C" />
          <Ionicons name="home-outline" size={45} color="#FB923C" />
        </View>
        <View className="flex-row justify-around items-center pt-32">
          <Ionicons name="business-outline" size={38} color="#F97316" />
          <Ionicons name="home" size={42} color="#EA580C" />
        </View>
      </View>

      {/* Main Container */}
      <View className="flex-1 justify-center px-8">
        <View className="w-full max-w-md mx-auto">
          <LoginHeader />
          
          {/* Form Card */}
          <View className="bg-white/90 backdrop-blur-sm rounded-3xl p-8 shadow-2xl border border-orange-100">
            <LoginInput
              label="Email"
              placeholder="seu@email.com"
              value={email}
              onChangeText={setEmail}
              iconName="mail"
              keyboardType="email-address"
            />
            
            <LoginInput
              label="Senha"
              placeholder="Sua senha"
              value={password}
              onChangeText={setPassword}
              iconName="key"
              secureTextEntry
            />
            
            <LoginButton onPress={handleLogin} loading={loading} />
          </View>

          {/* Footer */}
          <View className="items-center mt-8">
            <Text className="text-orange-600 text-sm font-medium text-center">
              Gerencie seus imóveis com facilidade
            </Text>
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