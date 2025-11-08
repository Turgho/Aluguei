import React from 'react';
import { View, Text, TouchableOpacity, Dimensions } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { router } from 'expo-router';
import WelcomeSvg from '../../assets/images/welcome.svg';
import ScreenTransition from '../components/shared/ScreenTransition';
import { useTheme } from '../contexts/ThemeContext';

export default function WelcomeScreen() {
  const { isDark } = useTheme();
  const { width } = Dimensions.get('window');
  const imageWidth = width - 64; // Largura da tela menos padding (px-8 = 32px cada lado)
  
  return (
    <SafeAreaView className={`flex-1 ${isDark ? 'bg-dark-bg' : 'bg-white'}`}>
      <ScreenTransition showLoading={false}>

      {/* Main Container */}
      <View className="flex-1 justify-center px-8">
        <View className="w-full max-w-md mx-auto">
          {/* Welcome Illustration */}
          <View className="items-center mb-8">
            <WelcomeSvg width={imageWidth} height={200} />
          </View>
          
          {/* Welcome Text */}
          <View className="items-center mb-12">
            <Text className={`${isDark ? 'text-dark-text' : 'text-gray-800'} text-2xl font-bold text-center mb-4`}>
              Pronto para começar?
            </Text>
            <Text className={`${isDark ? 'text-dark-muted' : 'text-gray-600'} text-center text-base`}>
              Faça login ou crie sua conta para gerenciar seus imóveis
            </Text>
          </View>

          {/* Action Buttons */}
          <View className="flex-row gap-4">
            <TouchableOpacity
              className="flex-1 bg-orange-500 rounded-2xl py-4 active:bg-orange-600"
              onPress={() => router.push('./login')}
            >
              <Text className="text-white text-center font-bold text-lg">
                Entrar
              </Text>
            </TouchableOpacity>

            <TouchableOpacity
              className={`flex-1 border-2 border-orange-500 rounded-2xl py-4 ${isDark ? 'active:bg-orange-900/20' : 'active:bg-orange-50'}`}
              onPress={() => router.push('./register')}
            >
              <Text className="text-orange-500 text-center font-bold text-lg">
                Criar Conta
              </Text>
            </TouchableOpacity>
          </View>


        </View>
      </View>
      </ScreenTransition>
    </SafeAreaView>
  );
}