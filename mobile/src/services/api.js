import { API_BASE_URL, API_ENDPOINTS } from '../config';

async function request(path, { method = 'GET', body, token } = {}) {
  const isFormData = typeof FormData !== 'undefined' && body instanceof FormData;
  const controller = new AbortController(); const timeout = setTimeout(() => controller.abort(), 12000);
  let response;
  try {
    response = await fetch(`${API_BASE_URL}${path}`, { method, signal: controller.signal, headers: { Accept: 'application/json', ...(!isFormData && body ? { 'Content-Type': 'application/json' } : {}), ...(token ? { Authorization: `Bearer ${token}` } : {}) }, ...(body ? { body: isFormData ? body : JSON.stringify(body) } : {}) });
  } catch (error) {
    if (error?.name === 'AbortError') throw new Error('The request took too long');
    throw new Error('You appear to be offline');
  } finally { clearTimeout(timeout); }

  const payload = await response.json().catch(() => ({}));
  if (!response.ok) {
    throw new Error(payload.message || payload.detail || 'Request failed');
  }
  return payload.data ?? payload;
}

export const authService = {
  login: (email, password) => request('/auth/login', { method: 'POST', body: { email, password } }),
  register: (values) => request('/auth/sign-up/user', { method: 'POST', body: values }),
  confirmPhone: (phoneNumber, token) => request('/auth/confirm-phone', { method: 'PATCH', body: { phone_number: phoneNumber, token } }),
  confirmEmail: (email, token) => request('/auth/confirm-email', { method: 'PATCH', body: { email, token } }),
};

export const celebrationService = {
  nearby: (latitude, longitude, radiusKm = 50, token) =>
    request(`/celebrations/nearby?latitude=${encodeURIComponent(latitude)}&longitude=${encodeURIComponent(longitude)}&radius_km=${encodeURIComponent(radiusKm)}`, { token }),
  create: (values, token) => {
    const form = new FormData();
    Object.entries(values).forEach(([key, value]) => {
      if (value !== undefined && value !== null && value !== '') form.append(key, String(value));
    });
    return request('/celebrations', { method: 'POST', body: form, token });
  },
};

export const profileService = {
  follow: (userId, token) => request(`/profile/${userId}/follow`, { method: 'PATCH', token }),
  block: (userId, token) => request(`/profile/${userId}/block`, { method: 'PATCH', token }),
};

export const moderationService = {
  queue: (token) => request('/admin/moderation-queue?status=pending', { token }),
  resolve: (id, decision, token) =>
    request(`/admin/moderation-queue/${id}/resolve`, { method: 'POST', body: { decision }, token }),
};

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
