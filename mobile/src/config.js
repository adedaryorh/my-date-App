// config.js
export const API_BASE_URL = process.env.EXPO_PUBLIC_API_URL || 'http://localhost:7070/v1';
export const API_ENDPOINTS = {
  RECOMMENDATIONS: '/ai/users/recommendations',
  MODERATE: '/ai/celebrations/moderate',
  INDEX: '/ai/celebrations/index',
  SEARCH: '/ai/celebrations/search',
  INTERACTIONS: '/ai/interactions',
  HEALTH: '/ai/health',
};
