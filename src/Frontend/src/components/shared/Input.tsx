import React from 'react';
import { View, Text, TextInput } from 'react-native';
import { Ionicons } from '@expo/vector-icons';

interface InputProps {
  label: string;
  placeholder: string;
  value: string;
  onChangeText: (text: string) => void;
  iconName: keyof typeof Ionicons.glyphMap;
  secureTextEntry?: boolean;
  keyboardType?: 'default' | 'email-address' | 'phone-pad' | 'numeric';
  mask?: 'phone' | 'cpf' | 'date';
}

const applyMask = (text: string, mask?: string) => {
  if (!mask) return text;
  
  const numbers = text.replace(/\D/g, '');
  
  switch (mask) {
    case 'phone':
      return numbers
        .replace(/(\d{2})(\d)/, '($1) $2')
        .replace(/(\d{5})(\d)/, '$1-$2')
        .substring(0, 15);
    
    case 'cpf':
      return numbers
        .replace(/(\d{3})(\d)/, '$1.$2')
        .replace(/(\d{3})(\d)/, '$1.$2')
        .replace(/(\d{3})(\d{1,2})/, '$1-$2')
        .substring(0, 14);
    
    case 'date':
      return numbers
        .replace(/(\d{2})(\d)/, '$1/$2')
        .replace(/(\d{2})\/(\d{2})(\d)/, '$1/$2/$3')
        .substring(0, 10);
    
    default:
      return text;
  }
};

export default function Input({
  label,
  placeholder,
  value,
  onChangeText,
  iconName,
  secureTextEntry = false,
  keyboardType = 'default',
  mask
}: InputProps) {
  
  const handleTextChange = (text: string) => {
    const maskedText = applyMask(text, mask);
    onChangeText(maskedText);
  };
  return (
    <View className="mb-4">
      <Text className="text-orange-800 mb-3 font-semibold text-base">{label}</Text>
      <View className="relative">
        <View className="absolute left-4 top-1/2 -translate-y-1/2 z-10">
          <Ionicons name={iconName} size={22} color="#EA580C" />
        </View>
        <TextInput
          className="w-full border-2 border-orange-200 rounded-2xl pl-14 pr-4 py-5 text-orange-900 bg-orange-50/50 focus:bg-white focus:border-orange-400"
          placeholder={placeholder}
          placeholderTextColor="#FB923C"
          value={value}
          onChangeText={handleTextChange}
          secureTextEntry={secureTextEntry}
          keyboardType={keyboardType}
          autoCapitalize="none"
          autoCorrect={false}
        />
      </View>
    </View>
  );
}