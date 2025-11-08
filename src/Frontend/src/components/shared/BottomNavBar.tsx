import React from 'react';
import { View, TouchableOpacity, Text } from 'react-native';
import { useSafeAreaInsets } from 'react-native-safe-area-context';
import { Ionicons } from '@expo/vector-icons';

export default function BottomNavBar() {
  const insets = useSafeAreaInsets();

  return (
    <View 
      className="absolute bottom-0 left-0 right-0 bg-white border-t border-gray-200"
      style={{ 
        paddingBottom: insets.bottom,
        shadowColor: '#000',
        shadowOffset: { width: 0, height: -2 },
        shadowOpacity: 0.1,
        shadowRadius: 8,
        elevation: 10
      }}
    >
      <View className="flex-row justify-around py-3">
        {/* Nova Propriedade */}
        <TouchableOpacity className="items-center p-2">
          <View className="bg-orange-500 w-10 h-10 rounded-full items-center justify-center mb-1">
            <Ionicons name="add-circle" size={20} color="white" />
          </View>
          <Text className="text-xs text-gray-600">Nova</Text>
        </TouchableOpacity>

        {/* Novo Contrato */}
        <TouchableOpacity className="items-center p-2">
          <View className="bg-blue-500 w-10 h-10 rounded-full items-center justify-center mb-1">
            <Ionicons name="document-text" size={20} color="white" />
          </View>
          <Text className="text-xs text-gray-600">Contrato</Text>
        </TouchableOpacity>

        {/* Registrar Pagamento */}
        <TouchableOpacity className="items-center p-2">
          <View className="bg-green-500 w-10 h-10 rounded-full items-center justify-center mb-1">
            <Ionicons name="cash" size={20} color="white" />
          </View>
          <Text className="text-xs text-gray-600">Pagamento</Text>
        </TouchableOpacity>

        {/* Relatórios */}
        <TouchableOpacity className="items-center p-2">
          <View className="bg-purple-500 w-10 h-10 rounded-full items-center justify-center mb-1">
            <Ionicons name="bar-chart" size={20} color="white" />
          </View>
          <Text className="text-xs text-gray-600">Relatórios</Text>
        </TouchableOpacity>
      </View>
    </View>
  );
}