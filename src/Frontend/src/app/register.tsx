import React, { useState } from 'react';
import { View, Text, TouchableOpacity, ScrollView } from 'react-native';
import { SafeAreaView, useSafeAreaInsets } from 'react-native-safe-area-context';
import { router } from 'expo-router';

import Input from '../components/shared/Input';
import Button from '../components/shared/Button';
import StatusBanner from '../components/shared/StatusBanner';
import { apiService } from '../services/api';

export default function RegisterScreen() {
  const insets = useSafeAreaInsets();

  const [form, setForm] = useState({
    name: '',
    email: '',
    password: '',
    phone: '',
    cpf: '',
    birthDate: ''
  });
  const [loading, setLoading] = useState(false);
  const [status, setStatus] = useState<{ type: 'success' | 'error' | null; message: string }>({ type: null, message: '' });

  const updateForm = (field: keyof typeof form, value: string) => {
    // Remove espaços no final do nome durante digitação
    const cleanValue = field === 'name' ? value.trimEnd() : value;
    setForm(prev => ({ ...prev, [field]: cleanValue }));
  };

  // Converte data DD/MM/AAAA para formato ISO AAAA-MM-DD
  const convertDateToBackend = (dateStr: string) => {
    if (!dateStr) return undefined;
    const [day, month, year] = dateStr.split('/');
    return year && month && day ? `${year}-${month.padStart(2, '0')}-${day.padStart(2, '0')}` : undefined;
  };

  const handleRegister = async () => {
    if (!form.name || !form.email || !form.password || !form.phone || !form.cpf) {
      setStatus({ type: 'error', message: 'Por favor, preencha todos os campos obrigatórios' });
      setTimeout(() => setStatus({ type: null, message: '' }), 3000);
      return;
    }

    setLoading(true);
    setStatus({ type: null, message: '' });
    
    try {
      await apiService.createOwner({
        name: form.name.trim(),
        email: form.email.toLowerCase(),
        password: form.password,
        phone: '+55' + form.phone.replace(/\D/g, ''), // Adiciona código do Brasil
        cpf: form.cpf.replace(/\D/g, ''), // Remove máscaras
        birth_date: convertDateToBackend(form.birthDate),
      });
      
      setStatus({ type: 'success', message: 'Conta criada com sucesso!' });
      setTimeout(() => router.push('./login'), 2000);
    } catch (error: any) {
      setStatus({ type: 'error', message: error.message || 'Erro ao criar conta' });
      setTimeout(() => setStatus({ type: null, message: '' }), 3000);
      console.error('Register error:', error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <SafeAreaView className="flex-1 bg-white">
      <ScrollView className="flex-1 px-8" showsVerticalScrollIndicator={false}>
        <View className="w-full max-w-md mx-auto py-8">
          {/* Welcome Text */}
          <View className="items-center mb-6">
            <Text className="text-orange-800 text-2xl font-semibold text-center mb-1">
              Criar Conta
            </Text>
            <Text className="text-orange-600 text-center text-sm">
              Preencha seus dados para começar
            </Text>
          </View>
          
          {/* Form */}
          <View>
            <Input
              label="Nome Completo"
              placeholder="Seu nome completo"
              value={form.name}
              onChangeText={(value: string) => updateForm('name', value)}
              iconName="person"
            />
            
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
              placeholder="Mínimo 6 caracteres"
              value={form.password}
              onChangeText={(value: string) => updateForm('password', value)}
              iconName="key"
              secureTextEntry
            />
            
            <Input
              label="Telefone"
              placeholder="(11) 99999-9999"
              value={form.phone}
              onChangeText={(value: string) => updateForm('phone', value)}
              iconName="call"
              keyboardType="phone-pad"
              mask="phone"
            />
            
            <Input
              label="CPF"
              placeholder="000.000.000-00"
              value={form.cpf}
              onChangeText={(value: string) => updateForm('cpf', value)}
              iconName="card"
              keyboardType="numeric"
              mask="cpf"
            />
            
            <Input
              label="Data de Nascimento (Opcional)"
              placeholder="DD/MM/AAAA"
              value={form.birthDate}
              onChangeText={(value: string) => updateForm('birthDate', value)}
              iconName="calendar"
              keyboardType="numeric"
              mask="date"
            />
            
            <Button onPress={handleRegister} loading={loading} title="Criar Conta" />
          </View>

          {/* Login Link */}
          <View className="items-center mt-6">
            <TouchableOpacity onPress={() => router.push('./login')}>
              <Text className="text-orange-600 text-sm font-medium">
                Já tem uma conta? <Text className="font-bold">Fazer login</Text>
              </Text>
            </TouchableOpacity>
          </View>
        </View>
      </ScrollView>
      
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