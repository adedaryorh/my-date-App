import { API_BASE_URL, API_ENDPOINTS } from '../config';

async function request(path, { method = 'GET', body, token } = {}) {
  const response = await fetch(`${API_BASE_URL}${path}`, {
    method,
    headers: {
      Accept: 'application/json',
      'Content-Type': 'application/json',
      ...(token ? { Authorization: `Bearer ${token}` } : {}),
    },
    ...(body ? { body: JSON.stringify(body) } : {}),
  });

  const payload = await response.json().catch(() => ({}));
  if (!response.ok) {
    throw new Error(payload.message || payload.detail || 'Request failed');
  }
  return payload.data ?? payload;
}

export const aiService = {
  getUserRecommendations: (limit = 10, token) =>
    request(API_ENDPOINTS.RECOMMENDATIONS, {
      method: 'POST',
      token,
      body: { limit },
    }),
  searchCelebrations: (query, limit = 10, token) =>
    request(API_ENDPOINTS.SEARCH, {
      method: 'POST',
      token,
      body: { query, limit },
    }),
  logInteraction: (targetId, action, metadata = {}, token) =>
    request(API_ENDPOINTS.INTERACTIONS, {
      method: 'POST',
      token,
      body: { target_id: targetId, action, metadata },
    }),
  healthCheck: () => request(API_ENDPOINTS.HEALTH),
};

export default request;
