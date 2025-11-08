import React from 'react';
import { View, Text, TextInput } from 'react-native';
import { Ionicons } from '@expo/vector-icons';
import { useTheme } from '../../contexts/ThemeContext';

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
  const { isDark } = useTheme();
  const [isFocused, setIsFocused] = React.useState(false);
  
  const handleTextChange = (text: string) => {
    const maskedText = applyMask(text, mask);
    onChangeText(maskedText);
  };
  return (
    <View className="mb-4">
      <Text className={`${isDark ? 'text-orange-400' : 'text-orange-800'} mb-3 font-semibold text-base`}>{label}</Text>
      <View className="relative">
        <View className="absolute left-4 top-1/2 -translate-y-1/2 z-10">
          <Ionicons name={iconName} size={22} color="#EA580C" />
        </View>
        <TextInput
          className={`w-full border-2 rounded-2xl pl-14 pr-4 py-5 ${
            isFocused
              ? isDark
                ? 'border-orange-400 bg-dark-card text-dark-text'
                : 'border-orange-400 bg-white text-orange-900'
              : isDark
                ? 'border-dark-border bg-dark-surface text-dark-text'
                : 'border-orange-200 bg-orange-50 text-orange-900'
          }`}
          placeholder={placeholder}
          placeholderTextColor={isDark ? '#a3a3a3' : '#FB923C'}
          value={value}
          onChangeText={handleTextChange}
          onFocus={() => setIsFocused(true)}
          onBlur={() => setIsFocused(false)}
          secureTextEntry={secureTextEntry}
          keyboardType={keyboardType}
          autoCapitalize="none"
          autoCorrect={false}
        />
      </View>
    </View>
  );
}