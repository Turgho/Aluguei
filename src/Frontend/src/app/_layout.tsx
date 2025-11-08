import React, { useEffect, useState } from 'react';
import { Stack } from 'expo-router';
import { AuthProvider } from '../contexts/AuthContext';
import { ThemeProvider } from '../contexts/ThemeContext';
import { SplashScreen } from '../components';

export default function RootLayout() {
  const [isReady, setIsReady] = useState(false);

  useEffect(() => {
    const timer = setTimeout(() => {
      setIsReady(true);
    }, 2500); // 2.5 segundos de splash

    return () => clearTimeout(timer);
  }, []);

  if (!isReady) {
    return <SplashScreen />;
  }

  return (
    <ThemeProvider>
      <AuthProvider>
        <Stack screenOptions={{ headerShown: false }}>
          <Stack.Screen name="welcome" />
          <Stack.Screen name="login" />
          <Stack.Screen name="register" />
          <Stack.Screen name="dashboard" />
          <Stack.Screen name="add-property" />
        </Stack>
      </AuthProvider>
    </ThemeProvider>
  );
}