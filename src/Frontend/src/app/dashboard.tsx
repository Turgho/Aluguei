import React, { useState, useEffect } from 'react';
import { View, Text, ScrollView, TouchableOpacity } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { Ionicons } from '@expo/vector-icons';
import { router } from 'expo-router';
import { useAuth } from '../contexts/AuthContext';
import { useTheme } from '../contexts/ThemeContext';
import BottomNavBar from '../components/shared/BottomNavBar';
import { apiService } from '../services/api';
import type { DashboardResponse } from '../types/api';

export default function DashboardScreen() {
  const { owner, logout } = useAuth();
  const { theme, toggleTheme, isDark } = useTheme();
  const [dashboardData, setDashboardData] = useState<DashboardResponse | null>(null);
  const [loading, setLoading] = useState(true);

  // Carrega dados do dashboard
  useEffect(() => {
    const loadDashboard = async () => {
      if (!owner?.id) return;
      
      try {
        const data = await apiService.getDashboard(owner.id);
        setDashboardData(data);
      } catch (error) {
        console.error('Erro ao carregar dashboard:', error);
      } finally {
        setLoading(false);
      }
    };

    loadDashboard();
  }, [owner?.id]);

  const handleLogout = async () => {
    await logout();
    router.replace('./welcome');
  };

  if (loading) {
    return (
      <SafeAreaView className={`flex-1 ${isDark ? 'bg-dark-bg' : 'bg-white'} justify-center items-center`}>
        <Text className={isDark ? 'text-dark-text' : 'text-gray-600'}>Carregando...</Text>
      </SafeAreaView>
    );
  }

  if (!dashboardData) {
    return (
      <SafeAreaView className={`flex-1 ${isDark ? 'bg-dark-bg' : 'bg-white'} justify-center items-center`}>
        <Text className="text-red-600">Erro ao carregar dados</Text>
      </SafeAreaView>
    );
  }

  return (
    <SafeAreaView className={`flex-1 ${isDark ? 'bg-dark-bg' : 'bg-white'}`}>
      {/* Header */}
      <View className={`px-6 py-4 border-b ${isDark ? 'border-dark-border' : 'border-gray-100'}`}>
        <View className="flex-row justify-between items-center">
          <View>
            <Text className={`${isDark ? 'text-orange-400' : 'text-orange-800'} text-xl font-semibold`}>
              Olá, {owner?.name?.split(' ')[0]}!
            </Text>
            <Text className={`${isDark ? 'text-dark-muted' : 'text-gray-600'} text-sm`}>Bem-vindo ao seu painel</Text>
          </View>
          <View className="flex-row items-center">
            <TouchableOpacity onPress={toggleTheme} className="p-2 mr-2">
              <Ionicons 
                name={isDark ? 'sunny' : 'moon'} 
                size={24} 
                color={isDark ? '#f59e0b' : '#6b7280'} 
              />
            </TouchableOpacity>
            <TouchableOpacity onPress={handleLogout} className="p-2">
              <Ionicons name="log-out-outline" size={24} color="#ea580c" />
            </TouchableOpacity>
          </View>
        </View>
      </View>

      <ScrollView className="flex-1 px-6" showsVerticalScrollIndicator={false} contentContainerStyle={{ paddingBottom: 100 }}>
        {/* Resumo Geral */}
        <View className="mt-6">
          <Text className={`${isDark ? 'text-dark-text' : 'text-gray-800'} text-lg font-semibold mb-4`}>Resumo Geral</Text>
          
          <View className="flex-row flex-wrap justify-between">
            {/* Total de Propriedades */}
            <View className={`${isDark ? 'bg-dark-surface' : 'bg-orange-50'} p-4 rounded-xl w-[48%] mb-4`}>
              <View className="flex-row items-center justify-between">
                <View>
                  <Text className={`${isDark ? 'text-orange-400' : 'text-orange-800'} text-2xl font-bold`}>{dashboardData.total_properties}</Text>
                  <Text className={`${isDark ? 'text-orange-300' : 'text-orange-600'} text-sm`}>Propriedades</Text>
                </View>
                <Ionicons name="home" size={24} color="#ea580c" />
              </View>
            </View>

            {/* Propriedades Alugadas */}
            <View className={`${isDark ? 'bg-dark-surface' : 'bg-green-50'} p-4 rounded-xl w-[48%] mb-4`}>
              <View className="flex-row items-center justify-between">
                <View>
                  <Text className={`${isDark ? 'text-green-400' : 'text-green-800'} text-2xl font-bold`}>{dashboardData.rented_properties}</Text>
                  <Text className={`${isDark ? 'text-green-300' : 'text-green-600'} text-sm`}>Alugadas</Text>
                </View>
                <Ionicons name="checkmark-circle" size={24} color="#16a34a" />
              </View>
            </View>

            {/* Receita Mensal */}
            <View className={`${isDark ? 'bg-dark-surface' : 'bg-blue-50'} p-4 rounded-xl w-[48%] mb-4`}>
              <View className="flex-row items-center justify-between">
                <View>
                  <Text className={`${isDark ? 'text-blue-400' : 'text-blue-800'} text-2xl font-bold`}>
                    R$ {dashboardData.monthly_revenue.toLocaleString()}
                  </Text>
                  <Text className={`${isDark ? 'text-blue-300' : 'text-blue-600'} text-sm`}>Receita Mensal</Text>
                </View>
                <Ionicons name="trending-up" size={24} color="#2563eb" />
              </View>
            </View>

            {/* Pagamentos Pendentes */}
            <View className={`${isDark ? 'bg-dark-surface' : 'bg-yellow-50'} p-4 rounded-xl w-[48%] mb-4`}>
              <View className="flex-row items-center justify-between">
                <View>
                  <Text className={`${isDark ? 'text-yellow-400' : 'text-yellow-800'} text-2xl font-bold`}>{dashboardData.pending_payments}</Text>
                  <Text className={`${isDark ? 'text-yellow-300' : 'text-yellow-600'} text-sm`}>Pendentes</Text>
                </View>
                <Ionicons name="time" size={24} color="#d97706" />
              </View>
            </View>
          </View>
        </View>

        {/* Alertas */}
        {dashboardData.overdue_payments > 0 && (
          <View className="mt-6">
            <Text className={`${isDark ? 'text-dark-text' : 'text-gray-800'} text-lg font-semibold mb-4`}>Alertas</Text>
            <View className={`${isDark ? 'bg-dark-surface border-red-800' : 'bg-red-50 border-red-200'} border p-4 rounded-xl`}>
              <View className="flex-row items-center">
                <Ionicons name="warning" size={20} color="#dc2626" />
                <Text className={`${isDark ? 'text-red-400' : 'text-red-800'} font-medium ml-2`}>
                  {dashboardData.overdue_payments} pagamento(s) em atraso
                </Text>
              </View>
              <Text className={`${isDark ? 'text-red-300' : 'text-red-600'} text-sm mt-1`}>
                Verifique os pagamentos vencidos e entre em contato com os inquilinos.
              </Text>
            </View>
          </View>
        )}

        {/* Gráficos Simples */}
        <View className="mt-6">
          <Text className={`${isDark ? 'text-dark-text' : 'text-gray-800'} text-lg font-semibold mb-4`}>Receita dos Últimos Meses</Text>
          <View className={`${isDark ? 'bg-dark-surface' : 'bg-gray-50'} p-4 rounded-xl mb-6`}>
            <View className="flex-row justify-between items-end h-32">
              {dashboardData.monthly_revenues.map((item, index) => (
                <View key={index} className="items-center flex-1">
                  <View 
                    className="bg-orange-500 w-8 rounded-t"
                    style={{ height: (item.revenue / 20000) * 100 }}
                  />
                  <Text className={`text-xs ${isDark ? 'text-dark-text' : 'text-gray-600'} mt-2`}>{item.month}</Text>
                  <Text className={`text-xs ${isDark ? 'text-dark-muted' : 'text-gray-500'}`}>R$ {(item.revenue / 1000).toFixed(0)}k</Text>
                </View>
              ))}
            </View>
          </View>

          <Text className={`${isDark ? 'text-dark-text' : 'text-gray-800'} text-lg font-semibold mb-4`}>Status das Propriedades</Text>
          <View className={`${isDark ? 'bg-dark-surface' : 'bg-gray-50'} p-4 rounded-xl mb-6`}>
            <View className="flex-row justify-between">
              <View className="items-center flex-1">
                <View className="w-16 h-16 bg-green-500 rounded-full items-center justify-center mb-2">
                  <Text className="text-white font-bold text-lg">{dashboardData.rented_properties}</Text>
                </View>
                <Text className={`text-sm ${isDark ? 'text-dark-text' : 'text-gray-600'}`}>Alugadas</Text>
              </View>
              <View className="items-center flex-1">
                <View className="w-16 h-16 bg-yellow-500 rounded-full items-center justify-center mb-2">
                  <Text className="text-white font-bold text-lg">{dashboardData.available_properties}</Text>
                </View>
                <Text className={`text-sm ${isDark ? 'text-dark-text' : 'text-gray-600'}`}>Disponíveis</Text>
              </View>
            </View>
          </View>
        </View>

        {/* Pagamentos Recentes */}
        <View className="mt-6">
          <Text className={`${isDark ? 'text-dark-text' : 'text-gray-800'} text-lg font-semibold mb-4`}>Pagamentos Recentes</Text>
          {dashboardData.recent_payments.length > 0 ? (
            <View className={`${isDark ? 'bg-dark-surface' : 'bg-gray-50'} rounded-xl`}>
              {dashboardData.recent_payments.map((payment, index) => (
                <View key={payment.id} className={`p-4 ${index < dashboardData.recent_payments.length - 1 ? `border-b ${isDark ? 'border-dark-border' : 'border-gray-200'}` : ''}`}>
                  <View className="flex-row justify-between items-center">
                    <View className="flex-1">
                      <Text className={`${isDark ? 'text-dark-text' : 'text-gray-800'} font-medium`}>{payment.tenant}</Text>
                      <Text className={`${isDark ? 'text-dark-text' : 'text-gray-600'} text-sm`}>{payment.property}</Text>
                      <Text className={`${isDark ? 'text-dark-muted' : 'text-gray-500'} text-xs`}>{payment.date}</Text>
                    </View>
                    <Text className="text-green-600 font-semibold">
                      R$ {payment.amount.toLocaleString()}
                    </Text>
                  </View>
                </View>
              ))}
            </View>
          ) : (
            <View className={`${isDark ? 'bg-dark-surface' : 'bg-gray-50'} p-6 rounded-xl items-center`}>
              <Text className={`${isDark ? 'text-dark-muted' : 'text-gray-500'}`}>Nenhum pagamento recente</Text>
            </View>
          )}
        </View>
      </ScrollView>

      <BottomNavBar />
    </SafeAreaView>
  );
}