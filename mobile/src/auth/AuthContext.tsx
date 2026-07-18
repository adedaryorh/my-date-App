import * as SecureStore from 'expo-secure-store';
import React, { createContext, useContext, useEffect, useMemo, useState } from 'react';
import { Platform } from 'react-native';

import { authService } from '@/services/api';

const SESSION_KEY = 'celebut.session';

type User = { id: string; email: string; first_name?: string; username?: string; role?: string };
type Session = { accessToken: string; refreshToken?: string; user: User };
export type VerificationChannel = 'email' | 'sms';
type AuthValue = {
  loading: boolean;
  token: string | null;
  user: User | null;
  login: (email: string, password: string) => Promise<void>;
  register: (values: Registration) => Promise<VerificationChannel>;
  confirmPhone: (phoneNumber: string, code: string) => Promise<void>;
  confirmEmail: (email: string, code: string) => Promise<void>;
  logout: () => Promise<void>;
};

export type Registration = { full_name: string; username: string; country_code: string; phone_number: string; email: string; date_of_birth: string; password: string };

const AuthContext = createContext<AuthValue | null>(null);

async function readSession() {
  if (Platform.OS === 'web') return typeof window === 'undefined' ? null : window.localStorage.getItem(SESSION_KEY);
  return SecureStore.getItemAsync(SESSION_KEY);
}

async function writeSession(value: string | null) {
  if (Platform.OS === 'web') {
    if (typeof window !== 'undefined') {
      if (value) window.localStorage.setItem(SESSION_KEY, value);
      else window.localStorage.removeItem(SESSION_KEY);
    }
    return;
  }
  if (value) await SecureStore.setItemAsync(SESSION_KEY, value);
  else await SecureStore.deleteItemAsync(SESSION_KEY);
}

export function AuthProvider({ children }: React.PropsWithChildren) {
  const [session, setSession] = useState<Session | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    readSession()
      .then((saved) => saved && setSession(JSON.parse(saved)))
      .catch(() => writeSession(null))
      .finally(() => setLoading(false));
  }, []);

  const value = useMemo<AuthValue>(() => ({
    loading,
    token: session?.accessToken ?? null,
    user: session?.user ?? null,
    login: async (email, password) => {
      const next = await authService.login(email.trim().toLowerCase(), password);
      setSession(next);
      await writeSession(JSON.stringify(next));
    },
    register: async (values) => {
	  const result = await authService.register({ ...values, email: values.email.trim().toLowerCase(), username: values.username.trim().toLowerCase() });
	  return result?.verification_channel === 'email' ? 'email' : 'sms';
    },
    confirmPhone: async (phoneNumber, code) => {
      await authService.confirmPhone(phoneNumber.trim(), code.trim());
    },
    confirmEmail: async (email, code) => {
      await authService.confirmEmail(email.trim().toLowerCase(), code.trim());
    },
    logout: async () => {
      setSession(null);
      await writeSession(null);
    },
  }), [loading, session]);

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
}

export function useAuth() {
  const value = useContext(AuthContext);
  if (!value) throw new Error('useAuth must be used inside AuthProvider');
  return value;
}
