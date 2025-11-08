import React from 'react';
import { View, Text, TouchableOpacity, Image, Dimensions } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { router } from 'expo-router';
import WelcomeHeader from '../components/welcome/WelcomeHeader';

export default function WelcomeScreen() {
  const { width } = Dimensions.get('window');
  const imageWidth = width - 64; // Largura da tela menos padding (px-8 = 32px cada lado)
  
  return (
    <SafeAreaView className="flex-1 bg-white">

      {/* Main Container */}
      <View className="flex-1 justify-center px-8">
        <View className="w-full max-w-md mx-auto">
          <WelcomeHeader />
          
          {/* Welcome Image */}
          <View className="items-center mb-8">
            <Image 
              source={require('../../assets/images/welcome.svg')}
              style={{ width: imageWidth, height: 200 }}
              resizeMode="contain"
            />
          </View>
          
          {/* Welcome Text */}
          <View className="items-center mb-12">
            <Text className="text-orange-800 text-xl font-semibold text-center mb-4">
              Gestão imobiliária simplificada
            </Text>
            <Text className="text-orange-600 text-center text-base">
              Conecte proprietários e inquilinos
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
              className="flex-1 border-2 border-orange-500 rounded-2xl py-4 active:bg-orange-50"
              onPress={() => router.push('./register')}
            >
              <Text className="text-orange-500 text-center font-bold text-lg">
                Criar Conta
              </Text>
            </TouchableOpacity>
          </View>


        </View>
      </View>
    </SafeAreaView>
  );
}