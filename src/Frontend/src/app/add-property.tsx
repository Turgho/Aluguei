import React, { useState, useEffect } from 'react';
import { View, Text, ScrollView, TouchableOpacity } from 'react-native';
import { SafeAreaView, useSafeAreaInsets } from 'react-native-safe-area-context';
import { Ionicons } from '@expo/vector-icons';
import { router } from 'expo-router';

import ScreenTransition from '../components/shared/ScreenTransition';
import { apiService } from '../services/api';
import { useAuth } from '../contexts/AuthContext';
import { useTheme } from '../contexts/ThemeContext';
import { Button, Input, StatusBanner } from '../components';

export default function AddPropertyScreen() {
  const { owner, token } = useAuth();
  const { isDark } = useTheme();
  const insets = useSafeAreaInsets();

  // Verifica autenticação
  useEffect(() => {
    if (!token || !owner?.id) {
      router.replace('./welcome');
      return;
    }
  }, [token, owner]);

  const [form, setForm] = useState({
    title: '',
    description: '',
    address: '',
    city: '',
    state: '',
    zipCode: '',
    bedrooms: '',
    bathrooms: '',
    area: '',
    rentAmount: ''
  });
  const [loading, setLoading] = useState(false);
  const [loadingCep, setLoadingCep] = useState(false);
  const [status, setStatus] = useState<{ type: 'success' | 'error' | null; message: string }>({ type: null, message: '' });

  const validStates = [
    'AC', 'AL', 'AP', 'AM', 'BA', 'CE', 'DF', 'ES', 'GO', 'MA',
    'MT', 'MS', 'MG', 'PA', 'PB', 'PR', 'PE', 'PI', 'RJ', 'RN',
    'RS', 'RO', 'RR', 'SC', 'SP', 'SE', 'TO'
  ];

  const fetchAddressByCep = async (cep: string) => {
    const cleanCep = cep.replace(/\D/g, '');
    if (cleanCep.length !== 8) return;

    setLoadingCep(true);
    try {
      const response = await fetch(`https://viacep.com.br/ws/${cleanCep}/json/`);
      const data = await response.json();
      
      if (!data.erro) {
        setForm(prev => ({
          ...prev,
          address: data.logradouro || prev.address,
          city: data.localidade || prev.city,
          state: data.uf || prev.state
        }));
      }
    } catch (error) {
      console.log('Erro ao buscar CEP:', error);
    } finally {
      setLoadingCep(false);
    }
  };

  const updateForm = (field: keyof typeof form, value: string) => {
    if (field === 'state') {
      const upperValue = value.toUpperCase();
      if (upperValue.length <= 2 && (upperValue === '' || validStates.some(state => state.startsWith(upperValue)))) {
        setForm(prev => ({ ...prev, [field]: upperValue }));
      }
    } else {
      setForm(prev => ({ ...prev, [field]: value }));
      
      if (field === 'zipCode') {
        const cleanCep = value.replace(/\D/g, '');
        if (cleanCep.length === 8) {
          fetchAddressByCep(cleanCep);
        }
      }
    }
  };

  const handleSubmit = async () => {
    if (!form.title || !form.address || !form.city || !form.state || !form.rentAmount) {
      setStatus({ type: 'error', message: 'Por favor, preencha todos os campos obrigatórios' });
      setTimeout(() => setStatus({ type: null, message: '' }), 3000);
      return;
    }

    if (!validStates.includes(form.state)) {
      setStatus({ type: 'error', message: 'Estado inválido. Use apenas siglas dos estados brasileiros (ex: SP, RJ)' });
      setTimeout(() => setStatus({ type: null, message: '' }), 3000);
      return;
    }

    if (!owner?.id) {
      setStatus({ type: 'error', message: 'Usuário não autenticado' });
      return;
    }

    setLoading(true);
    setStatus({ type: null, message: '' });
    
    try {
      await apiService.createProperty({
        owner_id: owner.id,
        title: form.title.trim(),
        description: form.description.trim(),
        address: form.address.trim(),
        city: form.city.trim(),
        state: form.state.trim().toUpperCase(),
        zip_code: form.zipCode.replace(/\D/g, ''),
        bedrooms: form.bedrooms ? parseInt(form.bedrooms) : 0,
        bathrooms: form.bathrooms ? parseInt(form.bathrooms) : 0,
        area: form.area ? parseInt(form.area) : 0,
        rent_amount: parseFloat(form.rentAmount.replace(/[^\d,]/g, '').replace(',', '.')),
        status: 'available'
      });
      
      setStatus({ type: 'success', message: 'Propriedade criada com sucesso!' });
      setTimeout(() => router.replace('./dashboard'), 2000);
    } catch (error: any) {
      setStatus({ type: 'error', message: error.message || 'Erro ao criar propriedade' });
      setTimeout(() => setStatus({ type: null, message: '' }), 3000);
      console.error('Create property error:', error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <SafeAreaView className={`flex-1 ${isDark ? 'bg-dark-bg' : 'bg-white'}`}>
      <ScreenTransition showLoading={true}>
      {/* Header */}
      <View className={`px-6 py-4 border-b ${isDark ? 'border-dark-border' : 'border-gray-100'}`}>
        <View className="flex-row items-center">
          <TouchableOpacity onPress={() => router.replace('./dashboard')} className="mr-4">
            <Ionicons name="arrow-back" size={24} color={isDark ? '#f5f5f5' : '#374151'} />
          </TouchableOpacity>
          <Text className={`${isDark ? 'text-dark-text' : 'text-gray-800'} text-xl font-semibold`}>
            Nova Propriedade
          </Text>
        </View>
      </View>

      <ScrollView className="flex-1 px-6" showsVerticalScrollIndicator={false}>
        <View className="py-6">
          <Input
            label="Título *"
            placeholder="Ex: Apartamento 2 quartos Centro"
            value={form.title}
            onChangeText={(value: string) => updateForm('title', value)}
            iconName="home"
          />
          
          <Input
            label="Descrição"
            placeholder="Descrição da propriedade"
            value={form.description}
            onChangeText={(value: string) => updateForm('description', value)}
            iconName="document-text"
          />
          
          <View className="relative">
            <Input
              label="CEP * (preencha primeiro para autocompletar)"
              placeholder="00000-000"
              value={form.zipCode}
              onChangeText={(value: string) => updateForm('zipCode', value)}
              iconName="mail"
              keyboardType="numeric"
              mask="cep"
            />
            {loadingCep && (
              <View className="absolute right-3 top-12">
                <Ionicons name="refresh" size={16} color={isDark ? '#9ca3af' : '#6b7280'} className="animate-spin" />
              </View>
            )}
          </View>
          
          <Input
            label="Endereço *"
            placeholder="Rua, número, complemento"
            value={form.address}
            onChangeText={(value: string) => updateForm('address', value)}
            iconName="location"
          />
          
          <View className="flex-row gap-4">
            <View className="flex-1">
              <Input
                label="Cidade *"
                placeholder="São Paulo"
                value={form.city}
                onChangeText={(value: string) => updateForm('city', value)}
                iconName="business"
              />
            </View>
            <View className="w-24">
              <Input
                label="Estado *"
                placeholder="SP"
                value={form.state}
                onChangeText={(value: string) => updateForm('state', value)}
                iconName="flag"
              />
            </View>
          </View>
          
          <View className="flex-row gap-4">
            <View className="flex-1">
              <Input
                label="Quartos"
                placeholder="2"
                value={form.bedrooms}
                onChangeText={(value: string) => updateForm('bedrooms', value)}
                iconName="bed"
                keyboardType="numeric"
              />
            </View>
            <View className="flex-1">
              <Input
                label="Banheiros"
                placeholder="1"
                value={form.bathrooms}
                onChangeText={(value: string) => updateForm('bathrooms', value)}
                iconName="water"
                keyboardType="numeric"
              />
            </View>
          </View>
          
          <Input
            label="Área (m²)"
            placeholder="70"
            value={form.area}
            onChangeText={(value: string) => updateForm('area', value)}
            iconName="resize"
            keyboardType="numeric"
          />
          
          <Input
            label="Valor do Aluguel *"
            placeholder="R$ 1.500,00"
            value={form.rentAmount}
            onChangeText={(value: string) => updateForm('rentAmount', value)}
            iconName="cash"
            keyboardType="numeric"
            mask="currency"
          />
          
          <Button onPress={handleSubmit} loading={loading} title="Criar Propriedade" />
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
      </ScreenTransition>
    </SafeAreaView>
  );
}