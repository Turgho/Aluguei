import React, { useEffect, useRef } from 'react';
import { View, Text, Animated, Dimensions } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { Ionicons } from '@expo/vector-icons';

export default function SplashScreen() {
  const fadeAnim = useRef(new Animated.Value(0)).current;
  const scaleAnim = useRef(new Animated.Value(0.5)).current;
  const slideAnim = useRef(new Animated.Value(50)).current;
  const { width } = Dimensions.get('window');

  useEffect(() => {
    Animated.sequence([
      // Animação do ícone
      Animated.parallel([
        Animated.timing(fadeAnim, {
          toValue: 1,
          duration: 800,
          useNativeDriver: true,
        }),
        Animated.spring(scaleAnim, {
          toValue: 1,
          tension: 50,
          friction: 7,
          useNativeDriver: true,
        }),
      ]),
      // Animação do texto
      Animated.timing(slideAnim, {
        toValue: 0,
        duration: 600,
        useNativeDriver: true,
      }),
    ]).start();
  }, []);

  return (
    <SafeAreaView className="flex-1 bg-dark-bg justify-center items-center">
      <View className="items-center">
        {/* Ícone animado */}
        <Animated.View
          style={{
            opacity: fadeAnim,
            transform: [{ scale: scaleAnim }],
          }}
          className="mb-8"
        >
          <View className="w-24 h-24 bg-dark-surface rounded-3xl items-center justify-center shadow-lg">
            <Ionicons name="home" size={48} color="#fb923c" />
          </View>
        </Animated.View>

        {/* Nome do app */}
        <Animated.View
          style={{
            opacity: fadeAnim,
            transform: [{ translateY: slideAnim }],
          }}
        >
          <Text className="text-orange-400 text-4xl text-center font-black italic mb-2">ALUGUEI!</Text>
          <Text className="text-dark-muted text-lg text-center">
            Gestão imobiliária simplificada
          </Text>
        </Animated.View>

        {/* Indicador de carregamento */}
        <Animated.View
          style={{ opacity: fadeAnim }}
          className="mt-12"
        >
          <View className="flex-row space-x-2">
            {[0, 1, 2].map((index) => (
              <Animated.View
                key={index}
                className="w-2 h-2 bg-orange-400 rounded-full"
                style={{
                  opacity: fadeAnim,
                  transform: [
                    {
                      scale: fadeAnim.interpolate({
                        inputRange: [0, 1],
                        outputRange: [0.5, 1],
                      }),
                    },
                  ],
                }}
              />
            ))}
          </View>
        </Animated.View>
      </View>
    </SafeAreaView>
  );
}