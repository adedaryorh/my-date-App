// config.js
export const API_BASE_URL = process.env.EXPO_PUBLIC_API_URL || 'http://localhost:8080';
// For Expo, we can use expo-constants or similar, but for simplicity we'll use a const.
// In production, this should be set via env vars.
export const API_ENDPOINTS = {
  RECOMMENDATIONS: '/ai/users/recommendations',
  MODERATE: '/ai/celebrations/moderate',
  INDEX: '/ai/celebrations/index',
  SEARCH: '/ai/celebrations/search',
  INTERACTIONS: '/ai/interactions',
  HEALTH: '/ai/health',
};
