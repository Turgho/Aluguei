import React from 'react';
import { View, Text } from 'react-native';

export default function WelcomeHeader() {
  return (
    <View className="items-center mb-2">
      <Text className="text-3xl font-black text-orange-600 text-center mb-1 italic">
        ALUGUEI!
      </Text>
      <Text className="text-orange-700 text-center text-base font-medium">
        Sua casa, seu lar 🏠
      </Text>
    </View>
  );
}