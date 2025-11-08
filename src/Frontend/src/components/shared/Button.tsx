import React from 'react';
import { View, Text, TouchableOpacity } from 'react-native';
import { Ionicons } from '@expo/vector-icons';

interface ButtonProps {
  onPress: () => void;
  loading: boolean;
  title: string;
}

export default function Button({ onPress, loading, title }: ButtonProps) {
  return (
    <TouchableOpacity
      className={`rounded-2xl py-5 mt-8 ${
        loading 
          ? 'bg-orange-300' 
          : 'bg-orange-500 active:bg-orange-600'
      }`}
      onPress={onPress}
      disabled={loading}
    >
      <View className="flex-row items-center justify-center">
        {loading ? (
          <View className="mr-3">
            <Ionicons name="refresh" size={24} color="white" />
          </View>
        ) : (
          <View className="mr-3">
            <Ionicons name="log-in" size={24} color="white" />
          </View>
        )}
        <Text className="text-white text-center font-bold text-xl">
          {loading ? 'Carregando...' : title}
        </Text>
      </View>
    </TouchableOpacity>
  );
}