import React from 'react';
import { View, Text, TextInput } from 'react-native';
import { Ionicons } from '@expo/vector-icons';

interface LoginInputProps {
  label: string;
  placeholder: string;
  value: string;
  onChangeText: (text: string) => void;
  iconName: keyof typeof Ionicons.glyphMap;
  secureTextEntry?: boolean;
  keyboardType?: 'default' | 'email-address';
}

export default function LoginInput({
  label,
  placeholder,
  value,
  onChangeText,
  iconName,
  secureTextEntry = false,
  keyboardType = 'default'
}: LoginInputProps) {
  return (
    <View className="mb-6">
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
          onChangeText={onChangeText}
          secureTextEntry={secureTextEntry}
          keyboardType={keyboardType}
          autoCapitalize="none"
          autoCorrect={false}
        />
      </View>
    </View>
  );
}