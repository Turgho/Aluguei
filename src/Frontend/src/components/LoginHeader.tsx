import React from 'react';
import { View, Text } from 'react-native';
import { Ionicons } from '@expo/vector-icons';

export default function LoginHeader() {
  return (
    <View className="items-center mb-12">
      <View className="items-center justify-center bg-gradient-to-r from-orange-500 to-amber-500 p-4 rounded-2xl mb-4">
        <Ionicons name="home" size={48} color="orange" />
      </View>
      <Text className="text-4xl font-black text-orange-600 text-center mb-3 italic">
        ALUGUEI!
      </Text>
      <Text className="text-orange-700 text-center text-lg font-medium">
        Sua casa, seu lar 🏠
      </Text>
    </View>
  );
}