import React, { useEffect, useRef, useState } from 'react';
import { Animated } from 'react-native';
import LoadingScreen from '../ui/LoadingScreen';

interface ScreenTransitionProps {
  children: React.ReactNode;
  duration?: number;
  showLoading?: boolean;
}

export default function ScreenTransition({ children, duration = 300, showLoading = true }: ScreenTransitionProps) {
  const fadeAnim = useRef(new Animated.Value(0)).current;
  const slideAnim = useRef(new Animated.Value(30)).current;
  const [isReady, setIsReady] = useState(!showLoading);

  useEffect(() => {
    const timer = setTimeout(() => {
      setIsReady(true);
      
      Animated.parallel([
        Animated.timing(fadeAnim, {
          toValue: 1,
          duration,
          useNativeDriver: true,
        }),
        Animated.timing(slideAnim, {
          toValue: 0,
          duration,
          useNativeDriver: true,
        }),
      ]).start();
    }, showLoading ? 1500 : 0);

    return () => clearTimeout(timer);
  }, [fadeAnim, slideAnim, duration, showLoading]);

  if (!isReady) {
    return <LoadingScreen />;
  }

  return (
    <Animated.View
      style={{
        flex: 1,
        opacity: fadeAnim,
        transform: [{ translateY: slideAnim }],
      }}
    >
      {children}
    </Animated.View>
  );
}