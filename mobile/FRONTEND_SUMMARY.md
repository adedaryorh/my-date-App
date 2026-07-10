# Celebut Mobile Frontend (React Native)

## Overview
This React Native application consumes the Celebut AI backend services to provide:
- User recommendations (Discover tab)
- Semantic celebration search (Explore tab)  
- Location-based nearby content (Nearby tab)
- Content creation with moderation checks (Create Post)

## Features Implemented

### 1. Discover Screen (`src/screens/DiscoverScreen.js`)
- Fetches recommended users to follow via `GET /ai/users/recommendations`
- Displays user avatars, names, bios
- Follow/unfollow functionality that logs interactions via `/ai/interactions`
- Loading and error states

### 2. Explore Screen (`src/screens/ExploreScreen.js`)
- Search celebrations using natural language queries
- Calls `POST /ai/celebrations/search` with query and limit parameters
- Displays search results with titles and descriptions
- Submit-on-enter functionality

### 3. Nearby Screen (`src/screens/NearbyScreen.js`)
- Requests location permissions
- Gets current position using Expo Location API
- Searches for celebrations near user's coordinates
- Displays distance-based results (placeholder implementation)

### 4. Create Post Screen (`src/screens/CreatePostScreen.js`)
- Form for title, description, and optional image URL
- **Content Moderation**: Before allowing submission, sends content to `POST /ai/celebrations/moderate` to check for toxicity
- If content is approved, proceeds to create post (placeholder for actual creation API)
- Loading states and error handling

## Service Layer (`src/services/api.js`)
- Axios instance configured with base URL from config
- Automatic JWT token injection from AsyncStorage
- Response/error interceptors for logging
- Wrapper functions for all AI endpoints:
  - `getUserRecommendations(userId, limit)`
  - `moderateCelebration(text)`
  - `indexCelebration(celebrationId, text)` 
  - `searchCelebrations(query, limit)`
  - `logInteraction(userId, targetId, action, metadata)`
  - `healthCheck()`

## Configuration (`src/config.js`)
- Centralized API configuration
- Base URL with fallback to localhost:8080
- Endpoint constants for all AI services

## Navigation (`src/navigation/AppNavigator.js`)
- Stack navigator with screens:
  - Discover (initial)
  - Explore
  - Nearby
  - CreatePost
- Header configuration inherited from stack defaults

## Prerequisites for Running
1. **Backend Services Running**:
   - Go backend server on `:8080` 
   - AI service (Python/FastAPI) accessible via backend
2. **Environment**:
   - Node.js with npm/yarn
   - Expo CLI (for development)
   - iOS/Android simulator or physical device

## Getting Started (when backend is ready)
```bash
# Install dependencies (if needed)
npm install

# Start development server
npm start  # or expo start

# Then run on iOS/Android simulator or device
```

## API Endpoints Used
All API calls are made to the Go backend which proxies to or directly implements the AI service endpoints:
- `GET /ai/users/recommendations` - Get suggested users to follow
- `POST /ai/celebrations/moderate` - Check content for toxicity
- `POST /ai/celebrations/index` - Index content for search
- `POST /ai/celebrations/search` - Search celebrations semantically
- `POST /ai/interactions` - Log user interactions for ML feedback
- `GET /ai/health` - Check AI service health

## Notes
- Authentication token is assumed to be stored in AsyncStorage as 'userToken'
- Actual implementation would require proper auth flow (login/signup screens)
- Image upload functionality is not implemented (would need expo-image-picker or similar)
- Location permissions handling follows Expo best practices
- Error boundaries and retry mechanisms could be enhanced for production

## Files Created
- `App.js` - Entry point
- `src/config.js` - API configuration
- `src/services/api.js` - Service layer with axios instance
- `src/navigation/AppNavigator.js` - Screen navigation
- `src/screens/` - All four screen implementations
- `FRONTEND_SUMMARY.md` - This document

## Backend Integration Points
This frontend expects the following backend endpoints to be available:
1. Auth endpoints (login/signup) to obtain JWT token
2. User profile endpoints to get current user ID
3. The AI endpoints listed above, proxied through the Go backend
