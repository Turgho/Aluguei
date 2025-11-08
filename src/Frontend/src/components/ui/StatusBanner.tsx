import React from 'react';
import { View, Text } from 'react-native';
import { Ionicons } from '@expo/vector-icons';

interface StatusBannerProps {
  type: 'success' | 'error' | null;
  message: string;
}

export default function StatusBanner({ type, message }: StatusBannerProps) {
  if (!type) return null;

  return (
    <View className={`p-4 ${
      type === 'success' ? 'bg-green-500' : 'bg-red-500'
    }`}>
      <View className="flex-row items-center justify-center">
        <Ionicons 
          name={type === 'success' ? 'checkmark-circle' : 'alert-circle'} 
          size={20} 
          color="white" 
        />
        <Text className="text-white font-medium ml-2">
          {message}
        </Text>
      </View>
    </View>
  );
}