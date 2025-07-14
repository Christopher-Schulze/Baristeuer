'use client';

import Provider from '@/lib/trpc/Provider';
import AuthProvider from '@/lib/auth/Provider';
import { HeroUIProvider } from '@heroui/react';
import React from 'react';

interface ClientProvidersProps {
  children: React.ReactNode;
}

export default function ClientProviders({ children }: ClientProvidersProps) {
  return (
    <HeroUIProvider>
      <AuthProvider>
        <Provider>
          {children}
        </Provider>
      </AuthProvider>
    </HeroUIProvider>
  );
}