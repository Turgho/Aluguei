import React, { useState } from 'react';
import { View, ScrollView } from 'react-native';
import { router } from 'expo-router';
import Input from '../../../components/forms/Input';
import Button from '../../../components/forms/Button';
import { apiService } from '../../../services/api';
import { useAuth } from '../../../contexts/AuthContext';
import { BRAZILIAN_STATES } from '../../../lib/constants/states';
import { usePropertyForm } from '../hooks/usePropertyForm';

interface PropertyFormProps {
  onStatusChange: (status: { type: 'success' | 'error' | null; message: string }) => void;
}

export default function PropertyForm({ onStatusChange }: PropertyFormProps) {
  const { owner } = useAuth();
  const { form, updateForm, loading, handleSubmit } = usePropertyForm(onStatusChange);

  return (
    <ScrollView className="flex-1" showsVerticalScrollIndicator={false}>
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
  );
}