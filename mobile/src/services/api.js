import axios from 'axios';
import { API_BASE_URL, API_ENDPOINTS } from '../config';
import AsyncStorage from '@react-native-async-storage/async-storage';

const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Request interceptor to add auth token
api.interceptors.request.use(
  async (config) => {
    try {
      const token = await AsyncStorage.getItem('userToken');
      if (token) {
        config.headers.Authorization = `Bearer ${token}`;
      }
    } catch (error) {
      console.warn('Failed to get token from storage', error);
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// Response interceptor for error handling
api.interceptors.response.use(
  (response) => response,
  (error) => {
    console.error('API Error:', error.response?.data || error.message);
    return Promise.reject(error);
  }
);

// AI Service Functions
export const aiService = {
  // Get user recommendations
  getUserRecommendations: async (userId, limit = 10) => {
    try {
      const response = await api.post(API_ENDPOINTS.RECOMMENDATIONS, {
        userId,
        limit,
      });
      return response.data;
    } catch (error) {
      throw error;
    }
  },

  // Moderate celebration content
  moderateCelebration: async (text) => {
    try {
      const response = await api.post(API_ENDPOINTS.MODERATE, {
        text,
      });
      return response.data;
    } catch (error) {
      throw error;
    }
  },

  // Index celebration (for search)
  indexCelebration: async (celebrationId, text) => {
    try {
      const response = await api.post(API_ENDPOINTS.INDEX, {
        celebrationId,
        text,
      });
      return response.data;
    } catch (error) {
      throw error;
    }
  },

  // Search celebrations
  searchCelebrations: async (query, limit = 10) => {
    try {
      const response = await api.post(API_ENDPOINTS.SEARCH, {
        query,
        limit,
      });
      return response.data;
    } catch (error) {
      throw error;
    }
  },

  // Log interaction (for feedback loop)
  logInteraction: async (userId, targetId, action, metadata = {}) => {
    try {
      const response = await api.post(API_ENDPOINTS.INTERACTIONS, {
        userId,
        targetId,
        action,
        metadata,
      });
      return response.data;
    } catch (error) {
      throw error;
    }
  },

  // Health check
  healthCheck: async () => {
    try {
      const response = await api.get(API_ENDPOINTS.HEALTH);
      return response.data;
    } catch (error) {
      throw error;
    }
  },
};

export default api;
