import React, { useEffect, useRef } from 'react';
import { View, Text, Animated } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { Ionicons } from '@expo/vector-icons';
import { useTheme } from '../../contexts/ThemeContext';

interface LoadingScreenProps {
  message?: string;
}

export default function LoadingScreen({ message = 'Carregando...' }: LoadingScreenProps) {
  const { isDark } = useTheme();
  const spinValue = useRef(new Animated.Value(0)).current;
  const pulseValue = useRef(new Animated.Value(1)).current;

  useEffect(() => {
    const spinAnimation = Animated.loop(
      Animated.timing(spinValue, {
        toValue: 1,
        duration: 1000,
        useNativeDriver: true,
      })
    );

    const pulseAnimation = Animated.loop(
      Animated.sequence([
        Animated.timing(pulseValue, {
          toValue: 1.2,
          duration: 800,
          useNativeDriver: true,
        }),
        Animated.timing(pulseValue, {
          toValue: 1,
          duration: 800,
          useNativeDriver: true,
        }),
      ])
    );

    spinAnimation.start();
    pulseAnimation.start();

    return () => {
      spinAnimation.stop();
      pulseAnimation.stop();
    };
  }, []);

  const spin = spinValue.interpolate({
    inputRange: [0, 1],
    outputRange: ['0deg', '360deg'],
  });

  return (
    <SafeAreaView className={`flex-1 ${isDark ? 'bg-dark-bg' : 'bg-white'} justify-center items-center`}>
      <View className="items-center">
        <Animated.View
          style={{
            transform: [{ rotate: spin }, { scale: pulseValue }],
          }}
          className="mb-4"
        >
          <Ionicons 
            name="home" 
            size={48} 
            color={isDark ? '#fb923c' : '#ea580c'} 
          />
        </Animated.View>
        
        <Text className={`${isDark ? 'text-dark-text' : 'text-gray-800'} text-lg font-medium mb-2`}>
          {message}
        </Text>
        
        <View className="flex-row space-x-1">
          {[0, 1, 2].map((index) => (
            <Animated.View
              key={index}
              className={`w-2 h-2 rounded-full ${isDark ? 'bg-orange-400' : 'bg-orange-500'}`}
              style={{
                opacity: pulseValue,
                transform: [
                  {
                    scale: pulseValue.interpolate({
                      inputRange: [1, 1.2],
                      outputRange: [0.8, 1.2],
                    }),
                  },
                ],
              }}
            />
          ))}
        </View>
      </View>
    </SafeAreaView>
  );
}