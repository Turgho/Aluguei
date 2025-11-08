import React from 'react';
import { View, Text, TouchableOpacity, ScrollView } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { Ionicons } from '@expo/vector-icons';
import { router } from 'expo-router';
import { useAuth } from '../contexts/AuthContext';
import { useTheme } from '../contexts/ThemeContext';
import ScreenTransition from '../components/shared/ScreenTransition';

export default function ProfileScreen() {
  const { owner, logout, token } = useAuth();
  const { isDark } = useTheme();

  // Verifica autenticação
  React.useEffect(() => {
    if (!token || !owner?.id) {
      router.replace('./welcome');
      return;
    }
  }, [token, owner]);

  const handleLogout = async () => {
    await logout();
    router.replace('./welcome');
  };

  return (
    <SafeAreaView className={`flex-1 ${isDark ? 'bg-dark-bg' : 'bg-white'}`}>
      <ScreenTransition showLoading={false}>
        {/* Header */}
        <View className={`px-6 py-4 border-b ${isDark ? 'border-dark-border' : 'border-gray-100'}`}>
          <View className="flex-row items-center">
            <TouchableOpacity onPress={() => router.replace('./dashboard')} className="mr-4">
              <Ionicons name="arrow-back" size={24} color={isDark ? '#f5f5f5' : '#374151'} />
            </TouchableOpacity>
            <Text className={`${isDark ? 'text-dark-text' : 'text-gray-800'} text-xl font-semibold`}>
              Meu Perfil
            </Text>
          </View>
        </View>

        <ScrollView className="flex-1 px-6" showsVerticalScrollIndicator={false}>
          <View className="py-6">
            {/* Avatar e Info Principal */}
            <View className="items-center mb-8">
              <View className={`w-24 h-24 ${isDark ? 'bg-dark-surface' : 'bg-orange-100'} rounded-full items-center justify-center mb-4`}>
                <Ionicons name="person" size={48} color={isDark ? '#fb923c' : '#ea580c'} />
              </View>
              <Text className={`${isDark ? 'text-dark-text' : 'text-gray-800'} text-2xl font-bold mb-1`}>
                {owner?.name}
              </Text>
              <Text className={`${isDark ? 'text-dark-muted' : 'text-gray-600'} text-base`}>
                Proprietário
              </Text>
            </View>

            {/* Informações Pessoais */}
            <View className={`${isDark ? 'bg-dark-surface' : 'bg-gray-50'} rounded-xl p-4 mb-6`}>
              <Text className={`${isDark ? 'text-dark-text' : 'text-gray-800'} text-lg font-semibold mb-4`}>
                Informações Pessoais
              </Text>
              
              <View className="space-y-3">
                <View className="flex-row items-center">
                  <Ionicons name="mail" size={20} color={isDark ? '#9ca3af' : '#6b7280'} />
                  <Text className={`${isDark ? 'text-dark-text' : 'text-gray-700'} ml-3 flex-1`}>
                    {owner?.email}
                  </Text>
                </View>
                
                <View className="flex-row items-center">
                  <Ionicons name="call" size={20} color={isDark ? '#9ca3af' : '#6b7280'} />
                  <Text className={`${isDark ? 'text-dark-text' : 'text-gray-700'} ml-3 flex-1`}>
                    {owner?.phone || 'Não informado'}
                  </Text>
                </View>
              </View>
            </View>

            {/* Opções do Perfil */}
            <View className={`${isDark ? 'bg-dark-surface' : 'bg-gray-50'} rounded-xl mb-6`}>
              <TouchableOpacity className="flex-row items-center p-4 border-b border-gray-200 dark:border-dark-border">
                <Ionicons name="create" size={20} color={isDark ? '#fb923c' : '#ea580c'} />
                <Text className={`${isDark ? 'text-dark-text' : 'text-gray-700'} ml-3 flex-1`}>
                  Editar Perfil
                </Text>
                <Ionicons name="chevron-forward" size={16} color={isDark ? '#9ca3af' : '#6b7280'} />
              </TouchableOpacity>
              
              <TouchableOpacity className="flex-row items-center p-4 border-b border-gray-200 dark:border-dark-border">
                <Ionicons name="key" size={20} color={isDark ? '#fb923c' : '#ea580c'} />
                <Text className={`${isDark ? 'text-dark-text' : 'text-gray-700'} ml-3 flex-1`}>
                  Alterar Senha
                </Text>
                <Ionicons name="chevron-forward" size={16} color={isDark ? '#9ca3af' : '#6b7280'} />
              </TouchableOpacity>
              
              <TouchableOpacity className="flex-row items-center p-4">
                <Ionicons name="notifications" size={20} color={isDark ? '#fb923c' : '#ea580c'} />
                <Text className={`${isDark ? 'text-dark-text' : 'text-gray-700'} ml-3 flex-1`}>
                  Notificações
                </Text>
                <Ionicons name="chevron-forward" size={16} color={isDark ? '#9ca3af' : '#6b7280'} />
              </TouchableOpacity>
            </View>

            {/* Botão de Logout */}
            <TouchableOpacity 
              onPress={handleLogout}
              className="flex-row items-center justify-center bg-red-500 rounded-xl p-4"
            >
              <Ionicons name="log-out" size={20} color="white" />
              <Text className="text-white font-semibold ml-2">Sair da Conta</Text>
            </TouchableOpacity>
          </View>
        </ScrollView>
      </ScreenTransition>
    </SafeAreaView>
  );
}